package utils

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func InitViper(appName, configFile string) error {
	_ = godotenv.Load()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetDefault("app", appName)

	if err := LoadConfigInLocal(configFile); err != nil {
		return err
	}

	if Debug() {
		viper.Debug()
	}

	return nil
}

func LoadConfigInLocal(filename string) error {

	if filename == "" {
		viper.SetConfigFile("config/app.yml")
	} else {
		viper.SetConfigFile(filename)
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func ReleaseTag() string {
	env := os.Getenv("RELEASE")
	if env == "" {
		env = "local"
	}
	return env
}

func Env() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	return env
}
