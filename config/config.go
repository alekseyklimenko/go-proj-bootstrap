package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"os"
	"runtime"
)

type Config struct {
	Database struct {
		Dsn string `yaml:"dsn" envconfig:"DATABASE_URL"`
	} `yaml:"database"`

	Web struct {
		Port int `yaml:"port" envconfig:"PORT"`
	} `yaml:"web"`
}

func New() *Config {
	config := &Config{}
	err := readFile(config)
	if err == nil {
		err = readEnv(config)
	}
	if err != nil {
		fmt.Println("Error parsing config")
		fmt.Println(err)
		os.Exit(2)
	}
	config.ConfigRuntime()
	return config
}

func (conf *Config) ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func readFile(cfg *Config) error {
	f, err := os.Open("config/config.yml")
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	return decoder.Decode(cfg)
}

func readEnv(cfg *Config) error {
	return envconfig.Process("", cfg)
}
