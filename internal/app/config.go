package app

import (
	"fmt"
	"log"
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
		log.Println("Read config file error:", err)
		return err
	}

	if err := bindingConfig(v, CFG); err != nil {
		log.Println("Binding config error:", err)
		return err
	}
	return nil
}

func bindingConfig(vp *viper.Viper, cfg *Configs) error {
	if err := vp.Unmarshal(&cfg); err != nil {
		log.Println("Unmarshal config error:", err)
		return err
	}
	return nil
}
