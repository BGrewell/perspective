package main

import (
    "github.com/BGrewell/perspective/collector/core"
    log "github.com/sirupsen/logrus"
)

/*
Collector - This is the component of perspective that polls the REST endpoints and
aggregates all of the event records into a central store. It also provides the endpoints
that the dashboard connects to in order to poll the event information
 */



func main() {

    // Parse the configuration file
    config, err := core.LoadConfig("/etc/collectord/config.yaml")
    if err != nil {
        log.Fatalf("failed to parses config: %v", err)
    }

    // Start a go-routine to poll each sensor
    for idx, sensor := range config.Sensors {
        log.Printf("sensor %d: %s", idx, sensor.Name)
    }

    // Start web server
}
