package config

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	Timezone       string `mapstructure:"DEFAULT_TIMEZONE"`
	WebPort        string `mapstructure:"WEB_PORT"`
	DBPort         int    `mapstructure:"DB_PORT"`
	DBVersion      string `mapstructure:"DB_VERSION"`
	DBDriver       string `mapstructure:"DB_DRIVER"`
	DBPublicHost   string `mapstructure:"DB_PUBLIC_HOST"`
	DBPublicName   string `mapstructure:"DB_PUBLIC_NAME"`
	DBInternalHost string `mapstructure:"DB_INTERNAL_HOST"`
	DBInternalName string `mapstructure:"DB_INTERNAL_NAME"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
}

func Load() (*Config, error) {
	var cfg Config
	v := viper.New()

	// load env vars
	v.AutomaticEnv()

	// load vars from env when in production
	if gin.Mode() == gin.ReleaseMode {
		t := reflect.TypeOf(cfg)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			envName := field.Tag.Get("mapstructure")
			v.BindEnv(envName)
		}
	}

	// load vars from config.env when developing
	if gin.Mode() == gin.DebugMode {
		v.AddConfigPath(".")
		v.SetConfigType("env")
		v.SetConfigName("config")
		if err := v.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	if err := validateConfig(&cfg); err != nil {
		return nil, fmt.Errorf("unable to validate config: %w", err)
	}

	return &cfg, nil
}

func validateConfig(cfg *Config) error {
	v := reflect.ValueOf(*cfg)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name
		mapstructureTag := t.Field(i).Tag.Get("mapstructure")

		if field.IsZero() {
			return fmt.Errorf(
				"config validation error: %s (%s) is not set",
				fieldName,
				mapstructureTag,
			)
		}
	}

	return nil
}

func (cfg *Config) DBUrl() string {
	host := cfg.DBPublicHost
	name := cfg.DBPublicName
	if gin.Mode() == gin.ReleaseMode {
		host = cfg.DBInternalHost
		name = cfg.DBInternalName
	}
	DBUrl := fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s",
		cfg.DBDriver,
		cfg.DBUser,
		cfg.DBPassword,
		host,
		cfg.DBPort,
		name,
	)
	return DBUrl
}
