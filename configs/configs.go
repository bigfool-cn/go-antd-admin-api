package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type configs struct {
	App struct {
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		Debug       bool   `yaml:"debug"`
		SecretKey   string `yaml:"secret_key"`
	}
	Db struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
}

var Conf configs

func init() {

	rootPath,_ := os.Getwd()

	configPath := rootPath + "/configs/configs.yaml"

	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("open configs.yaml error: %v\n", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("read configs.yaml error: %v\n", err)
	}

	err = yaml.Unmarshal([]byte(data), &Conf)
	if err != nil {
		log.Fatalf("configs yaml unmarshal error: %v\n", err)
	}
}
