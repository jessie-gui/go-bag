package core

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/jessie-gui/go-bag/constant"
	"github.com/jessie-gui/go-bag/model"
	"github.com/spf13/viper"
)

// NewConfig 初始化配置。
func NewConfig() *model.Config {
	v := viper.New()
	c := &model.Config{}

	v.SetConfigFile(constant.ConfigFile)

	if err := v.ReadInConfig(); err != nil {
		log.Fatal("read config/config.yaml file failed: ", err)
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config file is updated: %v", e.Name)

		if err := viper.Unmarshal(c); err != nil {
			log.Printf("config file OnConfigChange Unmarshal failed: %v", err)
		}
	})

	if err := v.Unmarshal(c); err != nil {
		log.Fatal("config file Unmarshal failed: ", err)
	}

	log.Println("配置初始化完成！")

	return c
}
