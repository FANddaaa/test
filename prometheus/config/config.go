package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Prometheus Prometheus
}

type Prometheus struct {
	Port     int
	Interval int
}

func (cfg *Config) Init(configPath string) error {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("解析配置失败: %w", err)
	}

	return nil
}
