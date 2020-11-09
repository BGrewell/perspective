package network

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type TcpSensor struct {
	listenAddr string
	listenPort int
	connChan chan <-*ConnectionAttempt
	errChan chan <-error
	running bool
	listener *net.TCPListener
}

func (s *TcpSensor) Start(addr string, port int, ctx context.Context, conns chan<- *ConnectionAttempt, errs chan<-error) (err error) {
	s.connChan = conns
	s.errChan = errs
	s.listenAddr = addr
	s.listenPort = port
	a := &net.TCPAddr{
		IP: net.ParseIP(addr),
		Port: port,
	}
	err = s.listenTcp("tcp", a)
	if err != nil {
		return err
	}
	s.running = true
	go s.handleConnections(ctx)
	return nil
}

func (s *TcpSensor) Stop() (err error) {
	s.running = false
	time.Sleep(100 * time.Millisecond)
	if s.listener != nil {
		s.listener.Close()
	}
	return nil
}

func (s *TcpSensor) handleConnections(ctx context.Context) {
	for s.running {
		select {
		case <- ctx.Done():
			return
			default:
				conn, err := s.acceptTcp()
				if err != nil {
					s.errChan <- err
				}
				addrParts := strings.Split(conn.LocalAddr().String(), ":")
				port, _ := strconv.Atoi(addrParts[1])
				c := &ConnectionAttempt{
					IP:       addrParts[0], // because we are TPROXY'ing the connection the local address is actually the originator of the connection
					Port:     port,
					Location: "ip location not implemented",
					Lat:      0,
					Lon:      0,
				}
				s.connChan <- c
		}
	}
}

func (s *TcpSensor) listenTcp(network string, localAddr *net.TCPAddr) (err error) {
	s.listener, err = net.ListenTCP(network, localAddr)
	if err != nil {
		return err
	}

	var f *os.File
	f, err = s.listener.File()
	if err != nil {
		return fmt.Errorf("failed to get listener socket descriptor: %v", err)
	}
	defer f.Close()

	if err = syscall.SetsockoptInt(int(f.Fd()), syscall.SOL_IP, syscall.IP_TRANSPARENT, 1); err != nil {
		return fmt.Errorf("failed to set socket options IP_TRANSPARENT: %v", err)
	}

	return nil
}

func (s *TcpSensor) acceptTcp() (conn *net.TCPConn, err error) {
	return s.listener.AcceptTCP()
}
