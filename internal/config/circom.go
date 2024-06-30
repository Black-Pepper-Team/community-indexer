package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CircomConfig struct {
	PathToCircuits string `fig:"path_to_circuits,required"` //nolint
}

func (c *config) Circom() *CircomConfig {
	return c.circom.Do(func() interface{} {
		config := CircomConfig{}
		err := figure.
			Out(&config).
			From(kv.MustGetStringMap(c.getter, "circom")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out ethereum config"))
		}

		return &config
	}).(*CircomConfig)
}
