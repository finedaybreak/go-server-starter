package config

import (
	"flag"
	"fmt"
	"go-server-starter/internal/enum"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	JWT        JWTConfig        `mapstructure:"jwt"`
	Database   DatabaseConfig   `mapstructure:"database"`
	Redis      RedisConfig      `mapstructure:"redis"`
	AsynQ      AsynQConfig      `mapstructure:"asynQ"`
	Logger     LoggerConfig     `mapstructure:"logger"`
	GormLogger GormLoggerConfig `mapstructure:"gormLogger"`
	Mode       enum.ServerMode  `mapstructure:"-"`
}

type ViperConfig struct {
	v      *viper.Viper
	config *Config
	env    *enum.ServerMode
}

func NewViperConfig(env *enum.ServerMode) (*ViperConfig, error) {
	config := &Config{Mode: *env}
	v := viper.New()
	var vc = &ViperConfig{v, config, env}
	setDefaultConfig(v)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./configs")

	vc.LoadConfig()

	return vc, nil
}

func (vc *ViperConfig) GetConfig() *Config {
	return vc.config
}

func (vc *ViperConfig) Viper() *viper.Viper {
	return vc.v
}

func (vc *ViperConfig) LoadConfig() error {
	tryFiles := []string{
		"config",
		"config." + vc.env.String(),
	}
	if err := vc.LoadMultiFilesConfig(tryFiles); err != nil {
		return err
	}
	// vc.v.SetEnvPrefix("GO_SERVER_STARTER")
	vc.v.AutomaticEnv()
	if err := vc.v.Unmarshal(vc.config); err != nil {
		return fmt.Errorf("parse config failed: %w", err)
	}
	return nil
}

func (vc *ViperConfig) LoadMultiFilesConfig(tryFiles []string) error {
	for i, configName := range tryFiles {
		vc.v.SetConfigName(configName)
		var err error
		if i == 0 {
			err = vc.v.ReadInConfig()
		} else {
			err = vc.v.MergeInConfig()
		}
		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Printf("config file %s.yml not found, skip\n", configName)
				continue
			}
			return fmt.Errorf("load config file %s failed: %w", configName, err)
		}
		fmt.Printf("✓ loaded: %s.yml\n", configName)
	}
	return nil
}

func setDefaultConfig(v *viper.Viper) {
	setDefaultsFromStruct(v, "server", DefaultConfig.Server)
	setDefaultsFromStruct(v, "jwt", DefaultConfig.JWT)
	setDefaultsFromStruct(v, "database", DefaultConfig.Database)
	setDefaultsFromStruct(v, "redis", DefaultConfig.Redis)
	setDefaultsFromStruct(v, "asynQ", DefaultConfig.AsynQ)
	setDefaultsFromStruct(v, "logger", DefaultConfig.Logger)
	setDefaultsFromStruct(v, "gormLogger", DefaultConfig.GormLogger)
}

func setDefaultsFromStruct(v *viper.Viper, prefix string, structValue interface{}) {
	_type := reflect.TypeOf(structValue)
	_value := reflect.ValueOf(structValue)
	for i := 0; i < _value.NumField(); i++ {
		field := _value.Field(i)
		fieldType := _type.Field(i)
		tag := fieldType.Tag.Get("mapstructure")
		if tag == "" {
			tag = strings.ToLower(fieldType.Name)
		}
		configKey := fmt.Sprintf("%s.%s", prefix, tag)
		if field.IsValid() && field.CanInterface() {
			v.SetDefault(configKey, field.Interface())
		}
	}
}

func ParseMode() (*enum.ServerMode, error) {
	var mode = "dev"
	if envMode := os.Getenv("APP_MODE"); envMode != "" {
		mode = envMode
	}

	flagMode := flag.String("mode", mode, "set app mode, valid modes: dev, prod, test")

	flag.Parse()

	serverMode, err := enum.ParseServerMode(*flagMode)
	if err != nil {
		return nil, err
	}
	return &serverMode, nil
}
