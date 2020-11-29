package collector

import "net"

type Collector interface {
	Handle(conn net.Conn) (json string, err error)
}
