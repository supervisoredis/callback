package httpv1

import (
	"callback/data"
	"callback/module"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var t data.Telephonestatus

//提供一个获取请求的接口，处理和记录返回的电话拨通的状态和信息
func Telephonestatus(_ http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	status := string(body)
	_ = json.Unmarshal([]byte(status), &t)
	//保存到日志中
	if t.Status == "0" || t.Status == "1" {
		l := "[TELEPHONE_STATUS]" + time.Now().Format("2006-01-02 15:04:05") + ":  status: " + t.Status + ",  message: " + t.Msg
		module.WriteLog("telephone_log.log", l)
		urlValuse := url.Values{
			"endpoint":   {t.Result.Endpoint},
			"sname":      {t.Result.Sname},
			"event_type": {t.Result.Event_type},
			"phone":      {t.Result.Phone},
		}
		reqBody := urlValuse.Encode()
		conf := module.C.GetConf()
		resp, err := http.Post(conf.CallbackAddress, "application/json;charset=UTF-8", strings.NewReader(reqBody))
		if err != nil {
			module.WriteLog("ERROR.log", err.Error())
		}
		ll := "[INFO resend]" + time.Now().Format("2006-01-02 15:04:05") + ":  endpoint: " + t.Result.Endpoint + ",  sname: " + t.Result.Sname + ", event_type: " + t.Result.Event_type + ", phone: " + t.Result.Phone
		module.WriteLog("alarm_log.log", ll)
		defer resp.Body.Close()
	} else {
		l := "[TELEPHONE_STATUS]" + time.Now().Format("2006-01-02 15:04:05") + ":  status: " + t.Status + ",  message: " + t.Msg
		module.WriteLog("telephone_log.log", l)
	}
}
