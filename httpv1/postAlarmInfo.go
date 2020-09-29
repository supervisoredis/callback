package httpv1

import (
	"callback/data"
	"callback/module"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var a data.Alarm

//提供一个获取用户数据和告警信息的接口，通过比较和处理，拼接成为新的JSON格式数据使用POST发送给回调接口
func PostAlarmInfo(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	alarm := string(body)
	json.Unmarshal([]byte(alarm), &a)
	data.GetAllUserInfo()
	fmt.Println(data.U)
	//从告警信息中读取用户名数据，对比获取到的用户数据，得到电话字段，与告警信息做拼接形成新的告警模式
	for _, i := range a.Users {
		for _, j := range data.U.Dat.List {
			if i == j.Username {
				urlValuse := url.Values{
					"endpoint":   {a.Endpoint},
					"sname":      {a.Sname},
					"event_type": {a.Event_type},
					"phone":      {j.Phone},
				}
				fmt.Println(urlValuse)
				reqBody := urlValuse.Encode()
				conf := module.C.GetConf()
				resp, err := http.Post(conf.CallbackAddress, "application/json;charset=UTF-8", strings.NewReader(reqBody))
				l := "[INFO]" + time.Now().Format("2006-01-02 15:04:05") + ":  endpoint: " + a.Endpoint + ",  sname: " + a.Sname + ", event_type: " + a.Event_type + ", phone: " + j.Phone
				module.WriteLog("alarm_log.log", l)
				if err != nil {
					module.WriteLog("ERROR.log", err.Error())
				}
				//fmt.Println(resp)
				defer resp.Body.Close()
				//fmt.Println(string(body))
			}
		}
	}
}
