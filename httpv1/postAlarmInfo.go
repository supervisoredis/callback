package httpv1

import (
	"callback/data"
	"callback/module"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//提供一个获取用户数据和告警信息的接口，通过比较和处理，拼接成为新的JSON格式数据使用POST发送给回调接口
func PostAlarmInfo(_ http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	alarm := string(body)
	//fmt.Println("最开始接受到的信息是：" + alarm)
	var a data.Alarm
	json.Unmarshal([]byte(alarm), &a)
	fmt.Println("最开始接受到的信息是：" + a.Endpoint)
	data.GetAllUserInfo()
	for _, i := range a.Users {
		for _, j := range data.U.Dat.List {
			if i == j.Username {
				fmt.Println("循环之后进入了：" + a.Endpoint)
				reqBody := "endpoint=" + a.Endpoint + "&\nsname=" + a.Sname + "&\nevent_type=" + a.Event_type + "&\nphone=" + j.Phone + "&\ntimes=1"
				//fmt.Println(alarm)
				fmt.Println(reqBody)
				conf := module.C.GetConf()
				resp, err := http.Post(conf.CallbackAddress, "text/plain", strings.NewReader(reqBody))
				//l := "[INFO]" + time.Now().Format("2006-01-02 15:04:05") + ":  endpoint: " + a.Endpoint + ",  sname: " + a.Sname + ", event_type: " + a.Event_type + ", phone: " + j.Phone
				module.WriteLog("alarm_log.log", alarm)
				if err != nil {
					module.WriteLog("ERROR.log", err.Error())
				}
				defer resp.Body.Close()
			}
		}
		//if a.Sname == "主机失联" {
		//	if module.ServerPing(a.Endpoint) {
		//		conf := module.C.GetConf()
		//		for _, i := range a.Users {
		//			for _, j := range data.U.Dat.List {
		//				if i == j.Username {
		//					reqBody := "endpoint=" + a.Endpoint + "&\nsname=agent断开连接&\nevent_type=" + a.Event_type + "&\nphone=" + j.Phone + "&\ntimes=1"
		//					resp, err := http.Post(conf.CallbackAddress, "text/plain", strings.NewReader(reqBody))
		//					//l := "[INFO]" + time.Now().Format("2006-01-02 15:04:05") + ":  endpoint: " + a.Endpoint + ",  sname: agent断开连接, event_type:" + a.Event_type + "  phone:" + j.Phone
		//					module.WriteLog("alarm_log.log", alarm)
		//					if err != nil {
		//						module.WriteLog("ERROR.log", err.Error())
		//					}
		//					defer resp.Body.Close()
		//				}
		//			}
		//		}
		//	} else {
		//		conf := module.C.GetConf()
		//		for _, i := range a.Users {
		//			for _, j := range data.U.Dat.List {
		//				if i == j.Username {
		//					reqBody := "endpoint=" + a.Endpoint + "&\nsname=服务器断开连接&\nevent_type=" + a.Event_type + "&\nphone=" + j.Phone + "&\ntimes=1"
		//					resp, err := http.Post(conf.CallbackAddress, "text/plain", strings.NewReader(reqBody))
		//					//l := "[INFO]" + time.Now().Format("2006-01-02 15:04:05") + ":  endpoint: " + a.Endpoint + ",  sname: 服务器断开连接, event_type:" + a.Event_type + "   phone:" + j.Phone
		//					module.WriteLog("alarm_log.log", alarm)
		//					if err != nil {
		//						module.WriteLog("ERROR.log", err.Error())
		//					}
		//					defer resp.Body.Close()
		//				}
		//			}
		//		}
		//	}
		//} else {

		//	}
		//}
		//从告警信息中读取用户名数据，对比获取到的用户数据，得到电话字段，与告警信息做拼接形成新的告警模式
	}
}
