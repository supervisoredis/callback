package httpv1

import (
	"callback/data"
	"callback/module"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

var t data.Telephonestatus

//提供一个获取请求的接口，处理和记录返回的电话拨通的状态和信息
func Telephonestatus(_ http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	status := string(body)
	_ = json.Unmarshal([]byte(status), &t)
	//保存到日志中
	l := "[TELEPHONE_STATUS]" + time.Now().Format("2006-01-02 15:04:05") + ":  status: " + t.Status + ",  message: " + t.Msg
	module.WriteLog("telephone_log.log", l)
}
