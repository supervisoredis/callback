package module

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//记录回调地址的机构提，从配置文件中获取地址
type CallbackAddr struct {
	CallbackAddress      string `yaml:"callbackAddress"`
	LocalCallbackAddress string `yaml:"localCallbackAddress"`
	AdminPhone           string `yaml:"adminPhone"` //有一些重要的告警信息直接通知管理员
}

var C CallbackAddr

//从配置文件中获取回调地址信息，并作为返回值进行返回
func (c *CallbackAddr) GetConf() *CallbackAddr {
	yamlFile, err := ioutil.ReadFile("./etc/config.yml")
	if err != nil {
		WriteLog("ERROR.log", err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		WriteLog("ERROR.log", err.Error())
	}
	return c
}
