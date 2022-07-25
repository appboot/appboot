package configs

import (
	"os"
	"path"

	"github.com/go-ecosystem/utils/log"

	"github.com/spf13/viper"
)

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	// Find home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		log.E(err.Error())
		os.Exit(1)
	}

	// Search config in home directory with name ".appboot" (without extension).
	configPath := path.Join(home, ".appboot")
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.W("ReadInConfig error: %v", err)
		return
	}
	log.H("Using config file: %v", viper.ConfigFileUsed())
}

// GetConfig get config with key
func GetConfig(key string) (string, error) {
	err := viper.ReadInConfig()
	if err == nil {
		return viper.GetString(key), nil
	}
	return "", err
}
