package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type HTTPServer struct {
	Addr string `yaml:"address"`
}

type Config struct {
	Env          string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	DATABASE_URI string `env:"DATABASE_URI" env-required:"true"`
	HTTPServer   `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	if err := godotenv.Load(); err != nil {
		log.Printf(".env not loaded: %v", err)
	}

	configPath = os.Getenv("CONFIG_PATH")


	if configPath == "" {
		flags := flag.String("config", "", "path to the config file")
		flag.Parse()

		configPath = *flags
		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config path does not exist %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("Config path does not exist %s", err.Error())
	}

	return &cfg

}
