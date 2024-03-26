package ioc

type Config struct {
	ServiceName    string         `mapstructure:"service_name"`
	ServiceVersion string         `mapstructure:"service_version"`
	ServicePort    string         `mapstructure:"service_port"`
	PrometheusPort string         `mapstructure:"prometheus_port"`
	Database       DatabaseConfig `mapstructure:"database"`
	Redis          RedisConfig    `mapstructure:"redis"`
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

//func InitConfig(path string) {
//	viper.SetConfigFile(path)
//	err := viper.ReadInConfig()
//	if err != nil {
//		log.Fatalf("Failed to read config file: %v", err)
//	}
//
//	err = viper.Unmarshal(&AppConfig)
//	if err != nil {
//		log.Fatalf("Failed to unmarshal config file: %v", err)
//	}
//}
