package config

import (
	"io/ioutil"
	"os"
	"gopkg.in/yaml.v2"
	"github.com/golang/glog"
)

type Conf struct {
	NodeName       string   `yaml:"node"`
	Database       string   `yaml:"database"`
	Etcd   		   []string `yaml:"etcd"`
}

func GetConfFromFile() *Conf {
	var config Conf
	confPath := os.Args[2]
	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		glog.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		glog.Fatalf("error: %v", err)
	}
	return &config
}
