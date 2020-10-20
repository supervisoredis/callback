package main

import (
	"callback/httpv1"
	"callback/module"
	"net/http"
)

//主函数
func main() {
	http.HandleFunc("/hello", httpv1.PostAlarmInfo)
	http.HandleFunc("/telephonestatus", httpv1.Telephonestatus)
	//监听本机的IP和端口信息
	conf := module.C.GetConf()
	err := http.ListenAndServe(conf.LocalCallbackAddress+":"+conf.LocalPort, nil)
	if err != nil {
		module.WriteLog("ERROR.log", err.Error())
	}
}
