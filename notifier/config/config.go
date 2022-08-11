package config

import "github.com/kelseyhightower/envconfig"

type (
	Config struct {
		App       App
		Secrets   Secrets
		Endpoints Endpoints
		Database  Database
	}
	App struct {
		SenderEmail string `envconfig:"MULTIVERSE_NOTIFIER_SENDER_EMAIL"`
	}

	Endpoints struct {
	}

	Database struct {
	}

	Secrets struct {
		SendGridToken string `envconfig:"MULTIVERSE_NOTIFIER_SENDGRID_TOKEN"`
	}
)

var Configs Config

func Load() error {
	err := envconfig.Process("", &Configs)
	return err
}
