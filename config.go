package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config is the base of our configuration data structure
type Config struct {
	Jobs     []Job          `yaml:"jobs"`
	Selenium SeleniumConfig `yaml:"selenium"`
	Storage  StorageConfig  `yaml:"storage,omitempty"`
}

// SeleniumConfig holds the configuration for our Selenium service
type SeleniumConfig struct {
	URL              string `yaml:"url"`
	JobStaggerOffset int32  `yaml:"job-stagger-offset"`
}

// StorageConfig holds the configuration for various storage backends.
// More than one storage backend can be used simultaneously
type StorageConfig struct {
	Graphite  GraphiteConfig  `yaml:"graphite,omitempty"`
	Dogstatsd DogstatsdConfig `yaml:"dogstatsd,omitempty"`
}

// NewConfig creates an new config object from the given filename.
func NewConfig(filename *string) (*Config, error) {
	cfgFile, err := ioutil.ReadFile(*filename)
	if err != nil {
		return &Config{}, err
	}
	c := Config{}
	err = yaml.Unmarshal(cfgFile, &c)
	if err != nil {
		return &Config{}, err
	}

	if len(c.Jobs) == 0 {
		log.Fatalln("No jobs were configured!")
	}

	return &c, nil
}