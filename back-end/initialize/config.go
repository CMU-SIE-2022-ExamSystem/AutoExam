package initialize

import (
	"io/ioutil"
	"path/filepath"
	"strings"

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

	// check autolab information
	autolabInfoCheck()

	// read basic grader from "source/autograders" folder
	global.Settings.Basic_Grader = basicGraderCheck()
}

func autolabInfoCheck() {
	auth := global.Settings.Autolabinfo
	if auth.Protocol == "" {
		panic("protocol is not found in .yaml file, please check")
	} else if auth.Ip == "" {
		panic("ip is not found in .yaml file, please check")
	} else if auth.Client_id == "" {
		panic("client_id is not found in .yaml file, please check")
	} else if auth.Client_secret == "" {
		panic("client_secret is not found in .yaml file, please check")
	} else if auth.Redirect_uri == "" {
		panic("redirect_uri is not found in .yaml file, please check")
	} else if auth.Scope == "" {
		panic("scope is not found in .yaml file, please check")
	}
}

func basicGraderCheck() []string {
	files, err := ioutil.ReadDir("source/autograders/")
	if err != nil {
		panic("Basic Grader loading fail")
	}
	var basic_grader []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".py" {
			name := strings.Split(file.Name(), ".")
			basic_grader = append(basic_grader, name[0])
		}
	}
	return basic_grader
}
