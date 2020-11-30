package geoip

import (
	"github.com/oschwald/geoip2-golang"
	log "github.com/sirupsen/logrus"
	"net"
)

var (
	db *geoip2.Reader
)

func init() {
	var err error
	db, err = geoip2.Open("geolite2.mmdb")
	if err != nil {
		log.Fatal("failed to open ip database: %v", err)
	}
}

func Close() {
	db.Close()
}

func Lookup(address string) (record *geoip2.City, err error) {
	ip := net.ParseIP(address)
	record, err = db.City(ip)
	return record, err
}