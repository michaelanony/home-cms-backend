package utils

import (
	"github.com/BurntSushi/toml"
	"home-cms/model"
	"log"
)

func InitEnv() (Config *model.Config, err error) {
	Config = &model.Config{}
	if _, err := toml.DecodeFile("./conf/config.toml", &Config); err != nil {
		log.Fatal(err)
	}
	return
}
