package config

import (
	"log"
	"os"
	"path"

	"github.com/spf13/viper"
)

const defaultConfig = "ethcli.yml"

var userConfig = func() string {
	userDir, err := os.UserConfigDir()
	if err != nil {
		return ""
	}
	return path.Join(userDir, defaultConfig)
}

type Config struct {
	URL string `yaml:"url"`
}

var cfg Config

func Init() {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := readConfigs(); err != nil {
		if !os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	viper.AutomaticEnv()
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal(err)
	}
}

func readConfigs() error {
	var err error
	for _, file := range []string{defaultConfig, userConfig(), viper.GetString("config")} {
		if err = readConfig(file); err == nil {
			break
		}
	}
	return err
}

func readConfig(file string) error {
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// Get returns the config.
func Get() *Config {
	return &cfg
}
