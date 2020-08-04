package configs

import (
	"log"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Search config in home directory with name ".appboot" (without extension).
	configPath := path.Join(home, ".appboot")
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")

	log.Printf("configPath: %v", configPath)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("ReadInConfig error: %v", err)
	}
	log.Println("Using config file:", viper.ConfigFileUsed())
}

// GetConfig get config with key
func GetConfig(key string) (string, error) {
	err := viper.ReadInConfig()
	if err == nil {
		return viper.GetString(key), nil
	}
	return "", err
}
