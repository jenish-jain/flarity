package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	AppName         string `mapstructure:"APP_NAME"`
	ServerPort      string `mapstructure:"SERVER_PORT"`
	LogLevel        string `mapstructure:"LOG_LEVEL"`
	AssetsPath      string `mapstructure:"ASSETS_PATH"`
	MongoUsername   string `mapstructure:"MONGO_USERNAME"`
	MongoPass       string `mapstructure:"MONGO_PASSWORD"`
	MongoReplicaSet string `mapstructure:"MONGO_REPLICA_SET"`
	MongoHost       string `mapstructure:"MONGO_HOST"`
	MongoDbName     string `mapstructure:"MONGO_DB_NAME"`
	MongoAuthDb     string `mapstructure:"MONGO_AUTH_DB"`
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

func (c *Config) GetMongoURI() string {
	return fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority&appName=%s",
		c.MongoUsername,
		c.MongoPass,
		c.MongoHost,
		c.MongoDbName,
		c.AppName,
	)
}

func (c *Config) GetMongoDbName() string {
	return c.MongoDbName
}

func GetConfig() *Config {
	return &AppConfig
}
