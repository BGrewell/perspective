package helpers

type ConnectionAttempt struct {
	SrcIP         string `json:"src_ip" yaml:"src_ip" xml:"src_ip"`
	SrcPort       int    `json:"src_port" yaml:"src_port" xml:"src_port"`
	DstIP         string `json:"dst_ip" yaml:"dst_ip" xml:"dst_ip"`
	DstPort       int    `json:"dst_port" yaml:"dst_port" xml:"dst_port"`
	CollectorData string `json:"collector_data" yaml:"collector_data" xml:"collector_data"`
}
