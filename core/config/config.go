package config

import "github.com/kelseyhightower/envconfig"

type (
	Config struct {
		App       App
		Secrets   Secrets
		Endpoints Endpoints
		Database  Database
	}
	//todo: change APP_NAME in config and env file
	App struct {
		Port           string `envconfig:"APP_NAME_APP_PORT" default:"3000"`
		Environment    string `envconfig:"APP_NAME_APP_ENV"`
		AllowedOrigins string `envconfig:"APP_NAME_APP_ALLOWED_ORIGINS" default:"*"`
	}

	Endpoints struct {
		Core          string `envconfig:"APP_NAME_CORE_API_URL"`
		SecretManager string `envconfig:"APP_NAME_SECRET_MGR_URL"`
	}

	Database struct {
		Host     string `envconfig:"APP_NAME_DATABASE_HOST"`
		Port     int    `envconfig:"APP_NAME_DATABASE_PORT"`
		User     string `envconfig:"APP_NAME_DATABASE_USER"`
		Password string `envconfig:"APP_NAME_DATABASE_PASSWORD"`
		DbName   string `envconfig:"APP_NAME_DATABASE_DBNAME"`
		Extras   string `envconfig:"APP_NAME_DATABASE_EXTRAS"`
		Driver   string `envconfig:"APP_NAME_DATABASE_DRIVER" default:"postgres"`
	}

	Secrets struct {
		AuthServerJwtSecret    string `envconfig:"APP_NAME_AUTH_SERVER_JWT_SECRET"`
		AppName                string `envconfig:"APP_NAME_APP_NAME"`
		AppSecret              string `envconfig:"APP_NAME_APP_SECRET"`
		SimpleSecretPassPhrase string `envconfig:"APP_NAME_SIMPLE_SECRET_PP"`
	}
)

var Configs Config

func Load() error {
	err := envconfig.Process("", &Configs)
	return err
}
