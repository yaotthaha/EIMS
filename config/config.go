package config

import (
	"encoding/json"
	"fmt"
	"net"
)

type Config struct {
	Log           Log      `json:"log,omitempty"`
	ListenAddress string   `json:"listen_address,omitempty"`
	DataBase      DataBase `json:"database,omitempty"`
	Redis         Redis    `json:"redis,omitempty"`
}

type Log struct {
	Debug bool   `json:"debug,omitempty"`
	File  string `json:"file,omitempty"`
}

type DataBase struct {
	DriverName string `json:"driver_name,omitempty"`
	Url        string `json:"url,omitempty"`
	Init       bool   `json:"init,omitempty"`
}

type Redis struct {
	Address  string `json:"address,omitempty"`
	Port     uint16 `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type config Config

func (c *Config) UnmarshalJSON(content []byte) error {
	var _c config
	if err := json.Unmarshal(content, &_c); err != nil {
		return err
	}
	host, port, err := net.SplitHostPort(_c.ListenAddress)
	if err != nil {
		return err
	} else {
		_c.ListenAddress = net.JoinHostPort(host, port)
	}
	switch _c.DataBase.DriverName {
	case "mysql", "":
		_c.DataBase.DriverName = "mysql"
	default:
		return fmt.Errorf("invalid database driver name: %s", _c.DataBase.DriverName)
	}
	if _c.DataBase.DriverName != "" && _c.DataBase.Url == "" {
		return fmt.Errorf("database url is required")
	}
	if _c.Redis.Address != "" {
		if _c.Redis.Port >= 65535 {
			return fmt.Errorf("invalid redis port: %d", _c.Redis.Port)
		}
		if _c.Redis.Port == 0 {
			_c.Redis.Port = 6379
		}
	}
	*c = Config(_c)
	return nil
}
