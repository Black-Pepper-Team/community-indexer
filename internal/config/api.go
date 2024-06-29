package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type APIConfig struct {
	MockAPI bool `fig:"mock_api"` //nolint
}

func (c *config) API() *APIConfig {
	return c.api.Do(func() interface{} {
		var result APIConfig
		err := figure.
			Out(&result).
			From(kv.MustGetStringMap(c.getter, "api")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out ethereum config"))
		}

		return &result
	}).(*APIConfig)
}
