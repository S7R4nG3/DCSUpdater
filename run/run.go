package run

import (
	"os/exec"
)

func Cmd(c string, a []string) string {
	var out []byte
	cmd := exec.Command(c, a...)
	stdout, _ := cmd.StdoutPipe()
	err := cmd.Run()
	if err != nil {
		return string(err.Error())
	}
	stdout.Read(out)
	return string(out)
}
