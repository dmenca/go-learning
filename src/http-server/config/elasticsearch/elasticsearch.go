package elasticsearch

import (
	"github.com/olivere/elastic"
	eCofg "github.com/olivere/elastic/config"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	*eCofg.Config
	Scheme string
}

func (c *Config) GetClient() (*elastic.Client, error) {
	log.Info("elasticsearch url :", c.URL)
	urlFunc := elastic.SetURL(c.URL)
	sniffFunc := elastic.SetSniff(false)
	schemeFunc := elastic.SetScheme(c.Scheme)
	basicAuthFunc := elastic.SetBasicAuth(c.Username, c.Password)
	funcs := []elastic.ClientOptionFunc{
		urlFunc, sniffFunc, schemeFunc, basicAuthFunc,
	}
	client, err := elastic.NewClient(funcs...)
	if err != nil {
		return nil, err
	}
	return client, err
}
