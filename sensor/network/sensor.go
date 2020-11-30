package network

import "github.com/BGrewell/perspective/helpers"

type Sensor interface {
	Start(addr string, port int, conns chan<- *helpers.ConnectionAttempt, errs chan<-error) (err error)
	Stop() (err error)
}