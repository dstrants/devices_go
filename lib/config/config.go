package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

// Loads the environment configuration to a usable struct.
func LoadConfig() Config {
	var conf Config

	es, err := env.UnmarshalFromEnviron(&conf)
	if err != nil {
		//TODO: Move this to logrus
		log.Fatal(err)
	}

	conf.Extras = es

	return conf
}
