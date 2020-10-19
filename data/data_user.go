package data

import (
	"callback/module"
	"encoding/json"
	"os/exec"
)

//User对象 通过调用官方提供的API接口获取所有用户数据
type User struct {
	Dat Data   //JSON格式的数据内容写在一个新的结构体之中
	Err string //报错信息
}

//JSON格式的用户数据内容
type Data struct {
	List  []UserInfo //一个新的结构体，记录每一组单独的用户数据
	Total int        //用户数据总数
}

//具体的每一组的用户数据的结构体
type UserInfo struct {
	Id       int    //用户ID标识
	Username string //用户名
	Dispname string //用户别名
	Phone    string //用户登记手机号
	Email    string //用户登记邮箱
	Im       string //用户登记IM帐号（无用）
	Is_root  int    //是否是管理员的标识
}

var U User

func GetAllUserInfo() {
	//通过系统指令curl调用官方API获取所有用户信息（使用root帐号）
	conf := module.C.GetConf()
	command := "curl -u " + conf.UserName + ":" + conf.UserPasswd + " http://" + conf.LocalCallbackAddress + "/api/portal/user"
	cmd := exec.Command("/bin/bash", "-c", command)
	bytes, _ := cmd.Output()
	resp := string(bytes)
	json.Unmarshal([]byte(resp), &U)
}
