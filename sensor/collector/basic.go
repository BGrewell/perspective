package collector

import (
	"net"
)

type BasicData struct {
	Type string `json:"collector_type"`
	Payload []byte `json:"payload"`
}

type BasicCollector struct {
}

func (bc *BasicCollector) Handle(conn net.Conn) (collectorData interface{}, err error) {
	payload := &BasicData{}
	payload.Type = "basic"
	buffer := make([]byte, 1500)
	read, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}
	payload.Payload = buffer[:read]
	//jsonByte, err := json.Marshal(payload)
	return payload, nil
}
