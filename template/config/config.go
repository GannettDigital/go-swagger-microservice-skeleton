package config

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	viper.AutomaticEnv()
	viper.SetDefault("taxo.host", "aimepublish.gmti.gbahn.net")
	viper.SetDefault("demo", "false")
}
