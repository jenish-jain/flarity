package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	LogLevel   string `mapstructure:"LOG_LEVEL"`
	AssetsPath string `mapstructure:"ASSETS_PATH"`
}

func InitConfig(configName string) *Config {
	viper.AutomaticEnv()
	viper.SetConfigName(configName)
	viper.SetConfigType("env")
	viper.AddConfigPath("config")
	viper.AddConfigPath("../config/")
	viper.AddConfigPath("../../config/")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using Config file:", viper.ConfigFileUsed())
	} else {
		fmt.Printf("Cannot read Config file %s.env, relying on env configs", configName)
	}

	err := viper.UnmarshalExact(&AppConfig)
	if err != nil {
		panic(fmt.Errorf("fatal error unable to Unmarshal Config file: %s", err))
	}
	fmt.Printf("Config file loaded successfully %+v\n", &AppConfig)
	return &AppConfig
}

func (c *Config) GetLogLevel() string {
	return c.LogLevel
}

func (c *Config) GetServerPort() string {
	return c.ServerPort
}

func (c *Config) GetAssetsPath() string {
	return c.AssetsPath
}

func GetConfig() *Config {
	return &AppConfig
}
