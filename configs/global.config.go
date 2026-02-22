package configs

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

type GlobalConf struct {
	Env     string `envconfig:"ENV" default:"development"`
	AppPort string `envconfig:"APP_PORT" validate:"required"`
	AppHost string `envconfig:"APP_HOST" validate:"required"`
}

var GlobalConfig GlobalConf

func ReadGlobalCfg() error {
	if err := envconfig.Process("", &GlobalConfig); err != nil {
		return fmt.Errorf("read global cfg err: %v", err)
	}
	return nil
}

func (g *GlobalConf) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(g); err != nil {
		return fmt.Errorf("validate global cfg err: %v", err)
	}
	return nil
}
