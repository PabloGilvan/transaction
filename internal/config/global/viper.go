package global

import (
	"bytes"
	"errors"
	"github.com/PabloGilvan/transaction"
	"github.com/spf13/viper"
	"io"
	"log"
	"strings"
)

func ViperConfig() {
	appContent, err := transaction.Resources.ReadFile("resources/application.yml")
	if err != nil {
		log.Println("[VIPER] Error parsing application file, ", err)
	}

	conf, err := LoadViperConfig(bytes.NewBuffer(appContent))
	if err != nil {
		log.Println("[VIPER] Error parsing application file, ", err)
	}

	Viper = conf
}

func LoadViperConfig(source io.Reader) (*viper.Viper, error) {
	if source == nil {
		return nil, errors.New("nil source reader")
	}

	log.Println("[VIPER] Loading context")
	configName := "application"
	viperSetup := viper.GetViper()
	viperSetup.SetConfigType("yml")
	viperSetup.SetConfigName(configName)
	viperSetup.AllowEmptyEnv(true)
	viperSetup.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viperSetup.AutomaticEnv()
	if err := viper.ReadConfig(source); err != nil {
		return nil, errors.New("unable to load configuration")
	}

	return viperSetup, nil
}
