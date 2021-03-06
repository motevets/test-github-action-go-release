package main

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type configYaml struct {
	Federates []string
	Port      uint
}

type Config struct {
	yaml configYaml
}

func ConfigFromFile(path string) (config Config, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		err = errors.Wrapf(err, "Could not read file %s", path)
		return
	}
	var rawConfig configYaml
	err = yaml.Unmarshal(data, &rawConfig)
	config.yaml = rawConfig
	if err != nil {
		err = errors.Wrap(err, "Could not unmarshal config yaml")
	}
	return
}

func (config Config) Federates() []string {
	return config.yaml.Federates
}

func (config Config) Port() uint {
	envPort, err := strconv.ParseUint(os.Getenv("PORT"), 10, 16)
	if err != nil && envPort != 0 {
		return uint(envPort)
	} else if config.yaml.Port != 0 {
		return config.yaml.Port
	} else {
		return 8000
	}
}
