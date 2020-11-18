package httpv1

import (
	"callback/data"
	"callback/module"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var a data.Alarm

//提供一个获取用户数据和告警信息的接口，通过比较和处理，拼接成为新的JSON格式数据使用POST发送给回调接口
func PostAlarmInfo(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	alarm := string(body)
	//fmt.Println(alarm)
	json.Unmarshal([]byte(alarm), &a)
	data.GetAllUserInfo()
	if a.Sname == "主机失联" {
		if module.ServerPing(a.Endpoint) {
			//fmt.Println("ping success")
			conf := module.C.GetConf()
			reqBody := "endpoint=" + a.Endpoint + "服务器&\nsname=agent断开连接&\nevent_type=" + a.Event_type + "&\nphone=" + conf.AdminPhone

			resp, err := http.Post(conf.CallbackAddress, "text/plain", strings.NewReader(reqBody))
			l := "[INFO]" + time.Now().Format("2006-01-02 15:04:05") + ":  endpoint: " + a.Endpoint + ",  sname: agent断开连接, event_type:" + a.Event_type + "&\nphone:" + conf.AdminPhone + "\ntimes:1"
			module.WriteLog("alarm_log.log", l)
			if err != nil {
				module.WriteLog("ERROR.log", err.Error())
			}
			defer resp.Body.Close()
		} else {
			//fmt.Println("ping fail")
			conf := module.C.GetConf()
			reqBody := "endpoint=" + a.Endpoint + "服务器&\nsname=主机断开连接，请及时关注&\nevent_type=\nphone=13262821472"

			resp, err := http.Post(conf.CallbackAddress, "text/plain", strings.NewReader(reqBody))
			l := "[INFO]" + time.Now().Format("2006-01-02 15:04:05") + ":  endpoint: " + a.Endpoint + ",  sname: 主机断开连接, event_type:alert , phone:15601740807 "
			module.WriteLog("alarm_log.log", l)
			if err != nil {
				module.WriteLog("ERROR.log", err.Error())
			}
			defer resp.Body.Close()
		}
	} else {
		for _, i := range a.Users {
			for _, j := range data.U.Dat.List {
				if i == j.Username {
					reqBody := "endpoint=" + a.Endpoint + "服务器&\nsname=" + a.Sname + "&\nevent_type=" + a.Event_type + "\nphone=" + j.Phone + "\ntimes:1"
					fmt.Println(reqBody)
					conf := module.C.GetConf()
					resp, err := http.Post(conf.CallbackAddress, "text/plain", strings.NewReader(reqBody))
					l := "[INFO]" + time.Now().Format("2006-01-02 15:04:05") + ":  endpoint: " + a.Endpoint + ",  sname: " + a.Sname + ", event_type: " + a.Event_type + ", phone: " + j.Phone
					module.WriteLog("alarm_log.log", l)
					if err != nil {
						module.WriteLog("ERROR.log", err.Error())
					}
					defer resp.Body.Close()
				}
			}
		}
	}
	//从告警信息中读取用户名数据，对比获取到的用户数据，得到电话字段，与告警信息做拼接形成新的告警模式
}
