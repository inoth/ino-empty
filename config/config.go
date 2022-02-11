package config

import (
	"os"

	"github.com/spf13/viper"
)

var Cfg *viper.Viper

type ViperConfig struct{}

func (m *ViperConfig) Init() error {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName(selectConfigName(nil))
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	Cfg = v
	return nil
}

func selectConfigName(path []string) string {
	if len(path) > 0 {
		return path[0]
	} else {
		e := os.Getenv("GORUNEVN")
		if len(e) > 0 {
			return "config." + e
		}
		return "config"
	}
}
