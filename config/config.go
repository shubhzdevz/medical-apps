package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Org1Name    string   `yaml:"org1_name"`
	Org1MSPID   string   `yaml:"org1_mspid"`
	Org1Domain  string   `yaml:"org1_domain"`
	Org1Peers   []string `yaml:"org1_peers"`
	Org1Ca      string   `yaml:"org1_ca"`
	Org1User    string   `yaml:"org1_user"`
	Org1Channel string   `yaml:"org1_channel"`
}

var (
	cfg *Config
)

func Init() {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "local"
	}
	configFile := fmt.Sprintf("config/%s.yaml", env)
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Sprintf("Error reading configuration file: %s", err.Error()))
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling configuration: %s", err.Error()))
	}
}

func Get() *Config {
	if cfg == nil {
		Init()
	}
	return cfg
}
