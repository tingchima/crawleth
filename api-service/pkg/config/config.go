// Package config provides
package config

import "github.com/spf13/viper"

// Config .
type Config struct {
	viper *viper.Viper
}

// NewConfig .
func NewConfig(configName, configPath string) (*Config, error) {
	v := viper.New()
	v.SetConfigName(configName)
	v.AddConfigPath(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Config{v}, nil
}

// ReadConfig .
func (c *Config) ReadConfig(k string, v interface{}) error {
	return c.viper.UnmarshalKey(k, v)
}
