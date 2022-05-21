package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Server struct {
	Host            string
	Port            int
	Debug           bool
	LogLevel        string
	ShutdownTimeout int
	Pid             string
}

type DataBase struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

type Config struct {
	Server   Server
	DataBase DataBase
	Fetcher  Fetcher
}

type Fetcher struct {
	Qps int64 //限流
}

var Conf Config

func Init(path string) error {
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		return err
	}
	fmt.Printf("**** %v", Conf)
	return nil
}
