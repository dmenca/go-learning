package config

import (
	"database/sql"
	"encoding/json"
	redisCli "github.com/go-redis/redis"
	"github.com/olivere/elastic"
	log "github.com/sirupsen/logrus"
	"http-server/config/elasticsearch"
	"http-server/config/mysql"
	"http-server/config/redis"
	"io/ioutil"
)

type DatabaseClient struct {
	Db    *sql.DB
	Es    *elastic.Client
	Redis *redisCli.Client
}

type ServerConfig struct {
	Redis         *redis.Config
	Elasticsearch *elasticsearch.Config
	MySQL         *mysql.Config
}

func NewDatabaseClient(config *ServerConfig) (*DatabaseClient, error) {
	mysqlClient, err := config.MySQL.GetClient()
	if err != nil {
		log.Error("fail to create mysql client", err)
		return nil, err
	}
	esClient, err := config.Elasticsearch.GetClient()
	if err != nil {
		log.Error("fail to create es client", err)
		return nil, err
	}
	redisClient := config.Redis.GetClient()
	return &DatabaseClient{
		Es:    esClient,
		Db:    mysqlClient,
		Redis: redisClient,
	}, err
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
