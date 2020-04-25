package config

import (
	"log"
	"gopkg.in/ini.v1"
	"os"
)

type ConfigList struct{
	Port int
	DbName string
	SQLDriver string
	ApiKey string
	ApiSecret string
	LogFile string
}

var Config ConfigList

func init(){
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustInt(),
		//example.sqlは値がないときのデフォルト値
		DbName: cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		ApiKey:    cfg.Section("configtest").Key("api_key").String(),
		ApiSecret: cfg.Section("configtest").Key("api_secret").String(),
		LogFile: cfg.Section("logtest").Key("log_file").String(),
	}
}