package config

import "time"

type Config struct {
	Server Server `yaml:"server"`
	Db     Db     `yaml:"db"`
	Client Client `yaml:"client"`
}

type Server struct {
	Host    string  `yaml:"host"`
	Port    Port    `yaml:"port"`
	Timeout Timeout `yaml:"timeout"`
}

type Db struct {
	Driver   string `yaml:"driver"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	HostDb   string `yaml:"hostdb"`
	PortDb   string `yaml:"portdb"`
	Dbname   string `yaml:"dbname"`
	Mode     string `yaml:"mode"`
}

type Client struct {
	APILayerAPIKey   string `yaml:"rest_api_api_key"`
	APILayerBasePath string `yaml:"rest_api_base_path"`
}

type Port struct {
	Http string `yaml:"http"`
}

type Timeout struct {
	ServerTimeout time.Duration `yaml:"server_timeout"`
}
