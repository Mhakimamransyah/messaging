package config

import (
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

//AppConfig Application configuration
type AppConfig struct {
	AppPort           int    `mapstructure:"app_port"`
	DbDriver          string `mapstructure:"db_mysql_driver"`
	DbHost            string `mapstructure:"db_mysql_host"`
	DbPort            int    `mapstructure:"db_mysql_port"`
	DbUsername        string `mapstructure:"db_mysql_username"`
	DbPassword        string `mapstructure:"db_mysql_password"`
	DbName            string `mapstructure:"db_mysql_name"`
	MongoDriver       string `mapstructure:"db_mongodb_driver"`
	MongoUsername     string `mapstructure:"db_mongodb_username"`
	MongoPassword     string `mapstructure:"db_mongodb_password"`
	MongoPort         int    `mapstructure:"db_mongodb_port"`
	MongoHost         string `mapstructure:"db_mongodb_host"`
	MongoCollection   string `mapstructure:"db_mongodb_collection"`
	CronDate          string `mapstructure:"cron_start"`
	MsgBrokerPort     int    `mapstructure:"message_broker_port"`
	MsgBrokerHost     string `mapstructure:"message_broker_host"`
	MsgBrokerUsername string `mapstructure:"message_broker_username"`
	MsgBrokerPassword string `mapstructure:"message_broker_password"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

//GetConfig Initiatilize config in singleton way
func GetConfig() *AppConfig {
	if appConfig != nil {
		return appConfig
	}

	lock.Lock()
	defer lock.Unlock()

	//re-check after locking
	if appConfig != nil {
		return appConfig
	}

	appConfig = initConfig()

	return appConfig
}

func initConfig() *AppConfig {
	godotenv.Load("config/.env")

	var defaultConfig AppConfig
	var finalConfig AppConfig

	defaultConfig.AppPort = 8000
	defaultConfig.DbDriver = "mysql"
	defaultConfig.DbHost = "127.0.0.1"
	defaultConfig.DbPort = 3306
	defaultConfig.DbUsername = "root"
	defaultConfig.DbPassword = ""
	defaultConfig.DbName = "your DB name"

	//use this if .env file (dont forget to run "source PATH_TO/.env" example "source config/.env")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("serotonin")
	viper.BindEnv("app_port")
	viper.BindEnv("db_mysql_driver")
	viper.BindEnv("db_mysql_host")
	viper.BindEnv("db_mysql_port")
	viper.BindEnv("db_mysql_username")
	viper.BindEnv("db_mysql_password")
	viper.BindEnv("db_mysql_name")
	viper.BindEnv("key")
	// viper.ReadInConfig()

	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed to extract config, will use default value")
		return &defaultConfig
	}
	return &finalConfig
}
