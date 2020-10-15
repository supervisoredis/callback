package main

import (
	"callback/httpv1"
	"callback/module"
	"net/http"
)

//主函数
func main() {
	module.TimeToCreatelogDir()
	http.HandleFunc("/hello", httpv1.PostAlarmInfo)
	http.HandleFunc("/telephonestatus", httpv1.Telephonestatus)
	//监听本机的IP和端口信息
	err := http.ListenAndServe("11.8.75.19:12345", nil)
	if err != nil {
		module.WriteLog("ERROR.log", err.Error())
	}
}
