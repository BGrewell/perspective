package main

import (
	"context"
	"fmt"
	"github.com/BGrewell/perspective/sensor/iptables"
	"github.com/BGrewell/perspective/sensor/network"
	"github.com/BGrewell/perspective/sensor/routes"
	log "github.com/sirupsen/logrus"
)

func main() {

	tcpPort := 9901
	//udpPort := 9902

	// create cancelable context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// setup channel for error and channel for connection attempts
	connChan := make(chan *network.ConnectionAttempt, 10000)
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

	for {
		conn := <- connChan
		fmt.Printf("connection from: %s:%d -> %s:%d\n", conn.SrcIP, conn.SrcPort, conn.DstIP, conn.DstPort)git add
		if conn.CollectorPayload != "" {
			fmt.Printf("payload: %s\n", conn.CollectorPayload)
			fmt.Printf("\n")
		} else {
			fmt.Println("payload: [none]")
			fmt.Printf("\n")
		}

	}
}
