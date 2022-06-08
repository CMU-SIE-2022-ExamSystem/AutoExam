package initialize

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/config"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/spf13/viper"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("./settings-dev.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	global.Settings = serverConfig
	// color.Blue("11111111", global.Settings.LogsAddress)
}
