package autolab

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/config"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/error"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
)

func Read_Autolab_Env() (autolab config.AutolabConfig) {
	autolab = global.Settings.Autolabinfo
	if autolab.Ip == "" {
		panic(error.EnvResponse("ip"))
	} else if autolab.Client_id == "" {
		panic(error.EnvResponse("client_id"))
	} else if autolab.Client_secret == "" {
		panic(error.EnvResponse("client_secret"))
	} else if autolab.Redirect_uri == "" {
		panic(error.EnvResponse("redirect_uri"))
	} else if autolab.Scope == "" {
		panic(error.EnvResponse("scope"))
	}
	return autolab
}
