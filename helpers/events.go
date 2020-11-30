package helpers

import (
	"github.com/oschwald/geoip2-golang"
	"time"
)

type SensorEvent struct {
	EventTime       string            `json:"event_time" yaml:"event_time" xml:"event_time"`
	SourceIP        string            `json:"source_ip" yaml:"source_ip" xml:"source_ip"`
	DestinationIP   string            `json:"destination_ip" yaml:"destination_ip" xml:"destination_ip"`
	SourcePort      int               `json:"source_port" yaml:"source_port" xml:"source_port"`
	DestinationPort int               `json:"destination_port" yaml:"destination_port" xml:"destination_port"`
	Location        EventLocationData `json:"location" yaml:"location" xml:"location"`
	CollectorData   interface{}            `json:"collector_data" yaml:"collector_data" xml:"collector_data"`
}

func NewSensorEvent(conn *ConnectionAttempt, record *geoip2.City) (event *SensorEvent, err error) {
	se := &SensorEvent{
		EventTime:       time.Now().Format(time.RFC3339Nano),
		SourceIP:        conn.SrcIP,
		DestinationIP:   conn.DstIP,
		SourcePort:      conn.SrcPort,
		DestinationPort: conn.DstPort,
		Location:        EventLocationData{},
		CollectorData:   conn.CollectorData,
	}
	err = se.Location.Parse(record)
	if err != nil {
		return nil, err
	}
	return se, nil
}

type EventLocationData struct {
	Country        string   `json:"country" yaml:"country" xml:"country"`
	Subdivisions   []string `json:"subdivisions" yaml:"subdivisions" xml:"subdivisions"`
	City           string   `json:"city" yaml:"city" xml:"city"`
	Metro          uint     `json:"metro" yaml:"metro" xml:"metro"`
	Latitude       float64  `json:"latitude" yaml:"latitude" xml:"latitude"`
	Longitude      float64  `json:"longitude" yaml:"longitude" xml:"longitude"`
	AccuracyRadius uint16   `json:"accuracy_radius" yaml:"accuracy_radius" xml:"accuracy_radius"`
	TimeZone       string   `json:"time_zone" yaml:"time_zone" xml:"time_zone"`
}

func (eld *EventLocationData) Parse(record *geoip2.City) (err error) {
	eld.Country = record.Country.Names["en"]
	if len(record.Subdivisions) > 0 {
		eld.Subdivisions = make([]string, 0)
		for _, subdiv := range record.Subdivisions {
			eld.Subdivisions = append(eld.Subdivisions, subdiv.Names["en"])
		}
	}
	eld.City = record.City.Names["en"]
	eld.Metro = record.Location.MetroCode
	eld.TimeZone = record.Location.TimeZone
	eld.Latitude = record.Location.Latitude
	eld.Longitude = record.Location.Longitude
	eld.AccuracyRadius = record.Location.AccuracyRadius
	return nil
}
