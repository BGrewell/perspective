package collector

import "net"

type Collector interface {
	Handle(conn net.Conn) (collectorData interface{}, err error)
}
