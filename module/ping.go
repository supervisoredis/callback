package module

import (
	"os/exec"
	"strings"
)

func ServerPing(target string) bool {
	//pinger, err := ping.NewPinger(target)
	//if err != nil {
	//	panic(err)
	//}
	//pinger.Count = 3
	//pinger.Interval = 50 * time.Millisecond
	//pinger.Timeout = time.Second * 2
	//pinger.Run() // blocks until finished
	//stats := pinger.Statistics()
	cmd := exec.Command("ping", "-c 1", target)
	out, _ := cmd.Output()
	context := string(out)
	//fmt.Println(context)
	if !strings.Contains(context, "Request") {
		return true
	}
	return false
}
