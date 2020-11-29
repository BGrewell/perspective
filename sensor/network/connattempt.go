package network

type ConnectionAttempt struct {
	SrcIP string
	SrcPort int
	DstIP string
	DstPort int
	Location string
	Lat float32
	Lon float32
	CollectorPayload string
}
