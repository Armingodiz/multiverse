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
		Port           string `envconfig:"MULTIVERSE_CORE_APP_PORT" default:"3000"`
		Environment    string `envconfig:"MULTIVERSE_CORE_APP_ENV"`
		AllowedOrigins string `envconfig:"MULTIVERSE_CORE_APP_ALLOWED_ORIGINS" default:"*"`
	}

	Endpoints struct {
		Core          string `envconfig:"MULTIVERSE_CORE_CORE_API_URL"`
		SecretManager string `envconfig:"MULTIVERSE_CORE_SECRET_MGR_URL"`
	}

	Database struct {
		Host           string `envconfig:"MULTIVERSE_CORE_DATABASE_HOST"`
		Port           int    `envconfig:"MULTIVERSE_CORE_DATABASE_PORT"`
		User           string `envconfig:"MULTIVERSE_CORE_DATABASE_USER"`
		Url            string `envconfig:"MULTIVERSE_CORE_DATABASE_URL"`
		Password       string `envconfig:"MULTIVERSE_CORE_DATABASE_PASSWORD"`
		DbName         string `envconfig:"MULTIVERSE_CORE_DATABASE_DBNAME"`
		CollectionName string `envconfig:"MULTIVERSE_CORE_DATABASE_COLLECTION_NAME"`
		Extras         string `envconfig:"MULTIVERSE_CORE_DATABASE_EXTRAS"`
		Driver         string `envconfig:"MULTIVERSE_CORE_DATABASE_DRIVER" default:"postgres"`
	}

	Secrets struct {
		AuthServerJwtSecret    string `envconfig:"MULTIVERSE_CORE_AUTH_SERVER_JWT_SECRET"`
		AppName                string `envconfig:"MULTIVERSE_CORE_MULTIVERSE_CORE"`
		AppSecret              string `envconfig:"MULTIVERSE_CORE_APP_SECRET"`
		SimpleSecretPassPhrase string `envconfig:"MULTIVERSE_CORE_SIMPLE_SECRET_PP"`
	}
)

var Configs Config

func Load() error {
	err := envconfig.Process("", &Configs)
	return err
}
