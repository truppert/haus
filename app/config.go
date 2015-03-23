package haus

import(
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Name string
	Email string
	Path string
	Pwd string
	Environments map[string]Environment
}

func ReadConfig(filename string)(*Config, error) {
	config := &Config{}
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(source, config)
	if err != nil {
		return config, err
	}
	config.Pwd,err = os.Getwd()
	if err != nil {
		return config,err
	}
	return config, nil
}