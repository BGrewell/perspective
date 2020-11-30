package collector

import (
	"encoding/json"
	"net"
)

type BasicData struct {
	Payload []byte `json:"payload"`
}

type BasicCollector struct {

}

func (bc *BasicCollector) Handle(conn net.Conn) (jsonPayload string, err error) {
	payload := &BasicData{}
	buffer := make([]byte, 1500)
	read, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}
	payload.Payload = buffer[:read]
	jsonByte, err := json.Marshal(payload)
	return string(jsonByte), err
}
