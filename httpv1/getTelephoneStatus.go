package httpv1

import (
	"callback/data"
	"callback/module"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//提供一个获取请求的接口，处理和记录返回的电话拨通的状态和信息
func Telephonestatus(_ http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	status := string(body)
	var t data.Telephonestatus
	_ = json.Unmarshal([]byte(status), &t)
	//保存到日志中
	go reSend(t)
}
func reSend(t data.Telephonestatus) {
	if t.Status == "1" && t.Alarm.Times != "5" {
		times, _ := strconv.Atoi(t.Alarm.Times)
		times = times + 1
		//l := "[TELEPHONE_STATUS]" + time.Now().Format("2006-01-02 15:04:05") + ":  status: " + t.Status + ",  message: " + t.Msg
		reqBody := "endpoint=" + t.Alarm.Endpoint + "&\nsname=" + t.Alarm.Sname + "&\nevent_type=" + t.Alarm.Event_type + "&\nphone=" + t.Alarm.Phone + "&\ntimes=" + strconv.Itoa(times)
		module.WriteLog("telephone_log.log", reqBody)
		fmt.Println(reqBody)
		conf := module.C.GetConf()
		time.Sleep(300 * time.Second)
		resp, err := http.Post(conf.CallbackAddress, "text/plain", strings.NewReader(reqBody))
		if err != nil {
			module.WriteLog("ERROR.log", err.Error())
		}
		//ll := "[INFO resend]" + time.Now().Format("2006-01-02 15:04:05") + ":  endpoint: " + t.Alarm.Endpoint + ",  sname: " + t.Alarm.Sname + ", event_type: " + t.Alarm.Event_type + ", phone: " + t.Alarm.Phone
		module.WriteLog("alarm_log.log", reqBody)
		defer resp.Body.Close()
	} else {
		fmt.Println(t)
		l := "[TELEPHONE_STATUS]" + time.Now().Format("2006-01-02 15:04:05") + ":  status: " + t.Status + ",  message: " + t.Msg
		module.WriteLog("telephone_log.log", l)
	}
}
