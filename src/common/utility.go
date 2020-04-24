package common

import (
	"fmt"
	"os/exec"
	"bufio"
)


func SysCall(cmdStr string) {
	//cmdStr := "sudo docker-compose up"
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