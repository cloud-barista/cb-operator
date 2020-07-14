package common

import (
	"bufio"
	"fmt"
	"os/exec"
)

var FileStr string
var CommandStr string
var TargetStr string

func SysCall(cmdStr string) {
	//cmdStr := "sudo docker-compose -f " + common.FileStr + " up"
	cmd := exec.Command("/bin/sh", "-c", cmdStr)

	cmdReader, _ := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	cmd.Start()
	scanner := bufio.NewScanner(cmdReader)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	cmd.Wait()
}
