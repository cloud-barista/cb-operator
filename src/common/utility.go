package common

import (
	"bufio"
	"fmt"
	"os/exec"
)

var FileStr string
var CommandStr string
var TargetStr string
var CB_OPERATOR_MODE string

const (
	Mode_DockerCompose   string = "DockerCompose"
	Mode_Kubernetes      string = "Kubernetes"
	CB_K8s_Namespace     string = "cloud-barista"
	CB_Helm_Release_Name string = "cloud-barista"
)

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
