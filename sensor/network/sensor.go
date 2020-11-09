package network

type Sensor interface {
	Start(addr string, port int, conns chan<- *ConnectionAttempt, errs chan<-error) (err error)
	Stop() (err error)
}