package modules

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Minio  *MinioConfig  `yaml:"minio"`
	Docker *DockerConfig `yaml:"docker"`
}

type MinioConfig struct {
	Endpoint  string `yaml:"endpoint"`
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
	Bucket    string `yaml:"bucket"`
	SSL       bool   `yaml:"ssl"`
}
type DockerConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func ReadConfig() (Config, error) {
	var c Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	//viper.AddConfigPath("D:\\go_project\\src\\operator-dev\\deploy-tool")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Config file not found")
		} else {
			// Config file was found but another error was produced
			fmt.Println("Config read err")
			return c, err
		}
	}
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Println("Config unmarshal err")
		return c, err
	}

	return c, nil
}
