package ioc

import (
	"github.com/spf13/viper"
	"log"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 17:03

var AppConfig *Config

type Config struct {
	ServiceName    string         `mapstructure:"service_name"`
	ServiceEnv     string         `mapstructure:"service_env"`
	ServiceId      int64          `mapstructure:"service_id"`
	ServiceVersion string         `mapstructure:"service_version"`
	ServicePort    string         `mapstructure:"service_port"`
	Language       string         `mapstructure:"language"`
	PrometheusPort string         `mapstructure:"prometheus_port"`
	Database       DatabaseConfig `mapstructure:"database"`
	Redis          RedisConfig    `mapstructure:"redis"`
	MongoCfg       MongoDBConfig  `mapstructure:"mongodb"`
	KafkaCfg       KafkaConfig    `mapstructure:"kakfa"`
	OTELCfg        OtelConfig     `mapstructure:"otel"`
}

type ServerConfig struct {
	Port string
	Mode string
}

type DatabaseConfig struct {
	DSN          string `mapstructure:"dsn"`
	MaxIdleConns int    `mapstructure:"max_idle_conn"`
	MaxOpenConns int    `mapstructure:"max_open_conn"`
}

type RedisConfig struct {
	DSN string `mapstructure:"dsn"`
}

type MongoDBConfig struct {
	DSN string `mapstructure:"dsn"`
}

type KafkaConfig struct {
	Brokers []string `mapstructure:"borkers"`
}

type OtelConfig struct {
	Url string `mapstructure:"url"`
}

func InitConfig(path string) *Config {
	var appCfg Config
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		panic(err)
	}

	err = viper.Unmarshal(&appCfg)
	if err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
		panic(err)
	}
	AppConfig = &appCfg
	return &appCfg
}
