package app

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var CFG = &Configs{}

type Configs struct {
	MongoDB struct {
		MongoDBHost  string        `mapstructure:"hosts"`
		Timeout      time.Duration `mapstructure:"timeout"`
		Username     string        `mapstructure:"username"`
		Password     string        `mapstructure:"password"`
		AuthDatabase string        `mapstructure:"auth_database"`
	} `mapstructure:"mongo_db"`
}

func InitConfig() error {
	configName := fmt.Sprintf("config.%s", "dev")

	v := viper.New()
	v.AddConfigPath("configs")
	v.SetConfigName(configName)

	if err := v.ReadInConfig(); err != nil {
		log.Infof("Read config file error: %+v", err)
		return err
	}

	if err := bindingConfig(v, CFG); err != nil {
		log.Infof("Binding config error: %+v", err)
		return err
	}
	return nil
}

func bindingConfig(vp *viper.Viper, cfg *Configs) error {
	if err := vp.Unmarshal(&cfg); err != nil {
		log.Infof("Unmarshal config error: %+v", err)
		return err
	}
	return nil
}
