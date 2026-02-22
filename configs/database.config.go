package configs

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

type DatabaseConf struct {
	Host     string `envconfig:"DATABASE_HOST" validate:"required"`
	Port     int    `envconfig:"DATABASE_PORT" default:"5432"`
	User     string `envconfig:"DATABASE_USER" validate:"required"`
	Passwd   string `envconfig:"DATABASE_PASSWD" validate:"required"`
	Database string `envconfig:"DATABASE_NAME" validate:"onvif"`
}

var DatabaseConfig DatabaseConf

func ReadDatabaseCfg() error {
	if err := envconfig.Process("", &DatabaseConfig); err != nil {
		return fmt.Errorf("read database cfg err: %v", err)
	}
	return nil
}

func (g *DatabaseConf) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(g); err != nil {
		return fmt.Errorf("validate database cfg err: %v", err)
	}
	return nil
}
