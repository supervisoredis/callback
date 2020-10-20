package main

import (
	"callback/module"
	"fmt"
	"os/exec"
)

func main() {
	conf := module.C.GetConf()
	command := "curl -u " + conf.UserName + ":" + conf.UserPasswd + " http://" + conf.LocalCallbackAddress + "/api/portal/user"
	cmd := exec.Command("/bin/bash", "-c", command)
	bytes, _ := cmd.Output()
	resp := string(bytes)
	fmt.Println(resp)
	fmt.Println(conf.LocalCallbackAddress + ":" + conf.LocalPort)
}
