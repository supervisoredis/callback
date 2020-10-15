package main

import (
	"fmt"
	"github.com/go-ping/ping"
)

func ServerPing(target string) bool {
	pinger, err := ping.NewPinger(target)
	if err != nil {
		panic(err)
	}

	pinger.Count = 3
	pinger.Run() // blocks until finished
	stats := pinger.Statistics()

	fmt.Println(stats)
	// 有回包，就是说明IP是可用的
	if stats.PacketsRecv >= 1 {
		return true
	}
	return false
}
func main() {
	if ServerPing("10.10.10.10") {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}

}
