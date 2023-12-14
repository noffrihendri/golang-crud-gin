package config

import "time"

type MainConfig struct {
	Server   ServerConfig `yaml:"Server"`
	Database DBConfig     `yaml:"Database"`
}

type (
	ServerConfig struct {
		Port            string        `yaml:"Port"`
		GracefulTimeout time.Duration `yaml:"GracefulTimeout"`
		ReadTimeout     time.Duration `yaml:"ReadTimeout"`
		WriteTimeout    time.Duration `yaml:"WriteTimeout"`
	}

	DBConfig struct {
		SlaveDSN        string `yaml:"SlaveDSN"`
		MasterDSN       string `yaml:"MasterDSN"`
		RetryInterval   int    `yaml:"RetryInterval"`
		MaxIdleConn     int    `yaml:"MaxIdleConn"`
		MaxConn         int    `yaml:"MaxConn"`
		ConnMaxLifetime string `yaml:"ConnMaxLifetime"`
	}
)
