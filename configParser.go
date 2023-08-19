package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"net/url"
	"os"
)

type YAMLUrl struct {
	Url *url.URL
}

// Intermediate type needed to unmarshal urls properly
type YamlConfig struct {
	ServerUrls []*YAMLUrl `yaml:"servers"`
	Algorithm  string     `yaml:"algorithm"`
}

// Implement UnmarshalYAML interface so that we can directly parse it as a url
func (yamlUrl *YAMLUrl) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var urlString string
	err := unmarshal(&urlString)
	if err != nil {
		log.Fatal("Can't read url: ", urlString)
	}

	newParsedUrl, err := url.Parse(urlString)

	if err != nil {
		log.Fatal("Can't read url: ", urlString)
	}

	yamlUrl.Url = newParsedUrl

	return nil
}

func (yamlConfig *YamlConfig) convertToConfig() *Config {
	var urls []*url.URL

	for _, yamlUrl := range yamlConfig.ServerUrls {
		urls = append(urls, yamlUrl.Url)
	}

	if len(urls) == 0 {
		log.Fatal("Error Parsing YAML: Please enter at least one server.")
	}

	return &Config{
		ServerUrls: urls,
		Algorithm:  yamlConfig.Algorithm,
	}
}

func parseYAML(filePath string) *Config {
	config := YamlConfig{}
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}
	err = yaml.Unmarshal(file, &config)

	if err != nil {
		log.Fatal("Error parsing YAML files: ", err)
	}

	return config.convertToConfig()
}