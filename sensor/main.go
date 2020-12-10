package main

import (
	"context"
	"fmt"
	"github.com/BGrewell/perspective/helpers"
	"github.com/BGrewell/perspective/sensor/geoip"
	"github.com/BGrewell/perspective/sensor/iptables"
	"github.com/BGrewell/perspective/sensor/network"
	"github.com/BGrewell/perspective/sensor/routes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	eventBacklogLimit = 10
)

var (
	eventChan chan *helpers.SensorEvent
	discardedEvents int
)

func EventsHandler(c *gin.Context) {
	eventCount := len(eventChan)
	events := make([]*helpers.SensorEvent, eventCount)
	eventsMissed := discardedEvents
	discardedEvents = 0
	for idx := 0; idx < eventCount; idx++ {
		events[idx] = <-eventChan
	}
	er := &helpers.EventsResponse{
		Date:         time.Now().Format(time.RFC3339Nano),
		EventCount:   eventCount,
		MissedEvents: eventsMissed,
		Events:       events,
	}
	c.JSON(http.StatusOK, er)
}

func main() {

	tcpPort := 9901
	//udpPort := 9902

	// create cancelable context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// setup channel for error and channel for connection attempts
	eventChan = make(chan *helpers.SensorEvent, eventBacklogLimit) //todo: for testing make small. we need to clear out events if the channel is full and not being serviced
	connChan := make(chan *helpers.ConnectionAttempt, 10000)
	errChan := make(chan error, 10000)

	// place route rule
	routes.AddTProxyRoute()

	// place iptables rule
	iptables.AddTProxyRule("tcp", tcpPort, 1)

	// setup tcp sensor
	tcp := network.TcpSensor{}
	err := tcp.Start("0.0.0.0", 9901, ctx, connChan, errChan)
	if err != nil {
		log.Fatal("failed to start tcp sensor: %v", err)
	}

	// setup REST handler //todo: look into using graphQL instead of REST
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/events", EventsHandler)
	go r.Run("0.0.0.0:65534")

	for {
		conn := <- connChan
		fmt.Printf("connection from: %s:%d -> %s:%d\n", conn.SrcIP, conn.SrcPort, conn.DstIP, conn.DstPort)
		record, err := geoip.Lookup(conn.SrcIP)
		event, err := helpers.NewSensorEvent(conn, record)
		if err != nil {
			fmt.Errorf("failed to generate new event: %v", err)
		} else {
			if len(eventChan) >= eventBacklogLimit {
				<- eventChan // discard oldest event
				discardedEvents++
			}
			eventChan <- event
		}
	}
}
