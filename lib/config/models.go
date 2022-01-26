package config

import (
	env "github.com/Netflix/go-env"
)

// The model that the environment configuration is parsed to
type Config struct {
	Extras env.EnvSet

	Mongo struct {
		Uri      string `env:"MONGO_CONNECTION_STRING,required=true"`
		Database string `env:"MONGO_DATABASE,default=devices"`
	}

	Slack struct {
		Token   string `env:"SLACK_TOKEN,required=true"`
		Channel string `env:"SLACK_CHANNEL,default=devices"`
	}

	LowBatteryThreshold int `env:"BATTERY_THRESHOLD,default=20"`

	Token string `env:"API_KEY,required=true"`
}
