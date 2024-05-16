package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"http-server/config/elasticsearch"
	"http-server/config/redis"
	"io/ioutil"
)

type ServerConfig struct {
	Redis         *redis.Config
	Elasticsearch *elasticsearch.Config
}

func LoadServerConfigFromFile(path string) (*ServerConfig, error) {
	blobs, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	config := ServerConfig{}
	err = json.Unmarshal(blobs, &config)
	if err != nil {
		log.Fatal("failed to load config ", err)
		return nil, err
	}
	return &config, nil
}
