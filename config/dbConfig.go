package config

import (
	"log"

	"github.com/spf13/viper"
)

type dbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

func GetDbConfig() *dbConfig {
	dbConfig := readDbConfig()
	return dbConfig
}

func readDbConfig() *dbConfig {
	var cfg dbConfig
	viper.AddConfigPath("./config")
	viper.SetConfigName("dbConfig")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
	}

	return &cfg
}
