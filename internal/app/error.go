package app

import (
	"github.com/spf13/viper"
)

var EM = &ErrorMessage{}

func InitErrorMessage() error {
	v := viper.New()
	v.AddConfigPath("configs")
	v.SetConfigName("error")

	if err := v.ReadInConfig(); err != nil {
		log.Infof("Read config file error: %+v", err)
		return err
	}

	if err := bindingErrorMessage(v, EM); err != nil {
		log.Infof("Binding config error: %+v", err)
		return err
	}
	return nil
}

func bindingErrorMessage(vp *viper.Viper, em *ErrorMessage) error {
	if err := vp.Unmarshal(&em); err != nil {
		log.Infof("Unmarshal config error: %+v", err)
		return err
	}
	return nil
}

type ErrorMessage struct {
	Internal struct {
		InternalServerError struct {
			Name    string `mapstructure:"name"`
			Details []ErrorDetail
		} `mapstructure:"internal_server_error"`
		BadRequest struct {
			Name    string `mapstructure:"name"`
			Details []ErrorDetail
		} `mapstructure:"bad_request"`
	} `mapstructure:"internal"`
}

type ErrorDetail struct {
	Field string `bson:"field,omitempty" json:"field,omitempty"`
	Issue string `bson:"issue,omitempty" json:"issue,omitempty"`
}
