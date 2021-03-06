package config

import (
	"sync"
	"time"

	"github.com/jinzhu/configor"
)

var (
	Cfg Configuration
	mu  sync.RWMutex
)

type ServerConfig struct {
	RunMode          string        `json:"run_mode"`
	ListenAddr       string        `json:"listen_addr"`
	LimitConnection  int           `json:"limit_connection"`
	RootRouterPrefix string        `json:"root_router_prefix"`
	ReadTimeout      time.Duration `json:"read_timeout"`
	WriteTimeout     time.Duration `json:"write_timeout"`
	IdleTimeout      time.Duration `json:"idle_timeout"`
	MaxHeaderBytes   int           `json:"max_header_bytes"`
}

type WitnessConfig struct {
	Country map[string]string `json:"country"`
	Logo    string            `json:"logo"`
}

type CommitteeConfig struct {
	Country map[string]string `json:"country"`
	Logo    string            `json:"logo"`
}

type LangsConfig struct {
	Allows  []string `json:"allows"`
	Default string   `json:"default"`
}

type MongoConfig struct {
	Uri        string `json:"uri"`
	Database   string `json:"database"`
	Collection string `json:"collection"`
	BlockCollection string `json:"block_collection"`
}

type (
	MysqlConfig struct {
		Driver   string `json:"driver"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DbName   string `json:"db_name"`
	}

	Configuration struct {
		Server    ServerConfig               `json:"server"`
		Witnesses map[string]WitnessConfig   `json:"witnesses"`
		Committee map[string]CommitteeConfig `json:"committee"`
		Langs     LangsConfig                `json:"langs"`
		Mongo     MongoConfig                `json:"mongo"`
		Mysql     MysqlConfig                `json:"mysql"`
		ApiUrl    string                     `json:"api_url"`
	}
)

func Init(file *string) (Configuration, error) {
	mu.Lock()
	defer mu.Unlock()

	err := configor.Load(&Cfg, *file)
	if err != nil {
		return Configuration{}, err
	}
	return Cfg, err
}

func GetConfig() Configuration {
	mu.Lock()
	defer mu.Unlock()
	return Cfg
}
