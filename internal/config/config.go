package config

import (
	"flag"
	"os"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Environment string

const (
	EnvDev  Environment = "dev"
	EnvProd Environment = "prod"
)

type Config struct {
	Env           Environment `yaml:"env" env-default:"local"`
	StoragePath   string      `yaml:"storage_path" env-required:"true"`
	GRPC          GRPCConfig  `yaml:"grpc"`
	MigrationPath string
	TokenTTL      time.Duration `yaml:"token_ttl" env-default:"1h"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

var instance Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		instance = *MustLoad()
	})

	return &instance
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exists " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("config path is empty: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
