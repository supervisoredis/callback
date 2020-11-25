package module

import (
	"os/exec"
	"strings"
)

func ServerPing(target string) bool {
	cmd := exec.Command("ping", "-c 1", target)
	out, _ := cmd.Output()
	context := string(out)
	if !strings.Contains(context, "Request") {
		return true
	}
	return false
}
