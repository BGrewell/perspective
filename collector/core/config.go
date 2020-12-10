package core

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type SensorConfig struct {
	Name         string   `json:"name" yaml:"name" xml:"name"`
	Host         string   `json:"host" yaml:"host" xml:"host"`
	Port         int      `json:"port" yaml:"port" xml:"port"`
	PollInterval int      `json:"poll_interval" yaml:"poll_interval" xml:"poll_interval"`
	Latitude     float64  `json:"latitude" yaml:"latitude" xml:"latitude"`
	Longitude    float64  `json:"longitude" yaml:"longitude" xml:"longitude"`
	Tags         []string `json:"tags" yaml:"tags" xml:"tags"`
}

type Configuration struct {
	ServerUrl   string          `json:"server_url" yaml:"server_url" xml:"server_url"`
	ServerPort  int             `json:"server_port" yaml:"server_port" xml:"server_port"`
	SSLCertFile string          `json:"ssl_cert_file" yaml:"ssl_cert_file" xml:"ssl_cert_file"`
	SSLKeyFile  string          `json:"ssl_key_file" yaml:"ssl_key_file" xml:"ssl_key_file"`
	Sensors     []*SensorConfig `json:"sensors" yaml:"sensors" xml:"sensors"`
}

func LoadConfig(filename string) (config *Configuration, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	c := &Configuration{}
	if err = yaml.Unmarshal(data, c); err != nil {
		return nil, err
	}
	return c, nil
}
