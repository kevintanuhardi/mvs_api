package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var cfg Config

// Load reads in config file and ENV variables if set.
func Load() *Config {
	// Search config in config directory with name ".brook-backend" (without extension).
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./../../../config") // load config from repository/mysql
	viper.SetConfigName(".brook.yml")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // read in environment variables that match
	bindEnvs(Config{})   // binding environtment variables https://github.com/spf13/viper/issues/188

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error read config file: %w", err))
	}

	// Parse config to config struct
	if err := viper.Unmarshal(&cfg, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
	}); err != nil {
		panic(fmt.Errorf("fatal error parse config file: %w", err))
	}

	return &cfg
}

func bindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(v.Interface(), append(parts, tv)...)
		default:
			_ = viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
