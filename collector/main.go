package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/BGrewell/perspective/collector/core"
    "github.com/BGrewell/perspective/helpers"
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
    "net/http"
)

var (
    EventCollection []*CollectorEvent
)

/*
Collector - This is the component of perspective that polls the REST endpoints and
aggregates all of the event records into a central store. It also provides the endpoints
that the dashboard connects to in order to poll the event information
 */

type CollectorEvent struct {
    Sensor string `json:"sensor"`
    Host string `json:"host"`
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    Tags []string `json:"tags"`
    Event *helpers.SensorEvent `json:"event"`
}

func (ce CollectorEvent) String() string {
    return ce.Json()
}

func (ce CollectorEvent) Json() string {
    jbytes, err := json.Marshal(&ce)
    if err != nil {
        return "error marshalling to json"
    }
    return string(jbytes)
}

func EventsHandler(c *gin.Context) {
    c.JSON(http.StatusOK, EventCollection)
}

func main() {

    // Parse the configuration file
    config, err := core.LoadConfig("/etc/collectord/config.yaml")
    if err != nil {
        log.Fatalf("failed to parses config: %v", err)
    }

    EventCollection = make([]*CollectorEvent, 0)

    // Start a go-routine to poll each sensor
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    for idx, sensor := range config.Sensors {
        log.Printf("sensor %d: %s", idx, sensor.Name)
        eventChan, err := core.Poll(sensor.Host, sensor.Port, sensor.PollInterval, ctx)
        if err != nil {
            log.Fatalf("failed to start sensor polling: %v", err)
        }

        // Start event processor
        s := sensor
        go func() {
            for {
                select {
                case event := <-eventChan:
                    ce := &CollectorEvent{
                        Sensor:    s.Name,
                        Host:      s.Host,
                        Latitude:  s.Latitude,
                        Longitude: s.Longitude,
                        Tags:      s.Tags,
                        Event:     event,
                    }
                    // collect and aggregate the events
                    EventCollection = append(EventCollection, ce)
                case <- ctx.Done():
                    // finish polling
                    return
                }

            }

        }()
    }




    // Start web server
    // setup REST handler //todo: look into using graphQL instead of REST
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()
    r.GET("/data", EventsHandler)
    if err := r.Run(fmt.Sprintf("0.0.0.0:%d", config.ServerPort)); err != nil {
        log.Fatal("error running web server: %v", err)
    }

    //for {
    //    conn := <- connChan
    //    fmt.Printf("connection from: %s:%d -> %s:%d\n", conn.SrcIP, conn.SrcPort, conn.DstIP, conn.DstPort)
    //    record, err := geoip.Lookup(conn.SrcIP)
    //    event, err := helpers.NewSensorEvent(conn, record)
    //    if err != nil {
    //        fmt.Errorf("failed to generate new event: %v", err)
    //    } else {
    //        if len(eventChan) >= eventBacklogLimit {
    //            <- eventChan // discard oldest event
    //            discardedEvents++
    //        }
    //        eventChan <- event
    //    }
    //}
}
