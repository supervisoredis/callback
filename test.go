package main

import (
	"callback/module"
	"fmt"
)

func main() {
	if module.ServerPing("11.8.75.81") {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
}
