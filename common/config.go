package common

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	cfg *ini.File
	err error
)

func InitConfig() {
	cfg, err = ini.Load("config.ini")
}

func GetConfigSection2Map(section string) map[string]string {
	if cfg == nil {
		cfg, err = ini.Load("config.ini")
	}
	getErr("load config", err)
	return cfg.Section(section).KeysHash()
}

func GetConfigValue(section, key string) string {
	if cfg == nil {
		cfg, err = ini.Load("config.ini")
	}
	getErr("load config", err)
	return cfg.Section(section).Key(key).Value()
}

func getErr(msg string, err error) {
	if err != nil {
		log.Printf("%v err->%v\n", msg, err)
	}
}
