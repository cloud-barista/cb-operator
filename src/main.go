/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cloud-barista/cb-operator/src/cmd"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func scanAndWriteMode() {

	fmt.Println("")
	fmt.Println("[Options]")
	fmt.Println("1: Single-node version (Docker Compose based)")
	fmt.Println("2: K8s-cluster version (requires a K8s cluster with Helm 3 enabled)")
	fmt.Println("")
	fmt.Print("Choose 1 or 2: ")

	var userInput uint8
	fmt.Scanf("%d", &userInput)

	var tempStr string

	switch userInput {
	case 1:
		fmt.Println("[1: Single-node version (Docker Compose based)] selected.")
		tempStr = "singlenode"
	case 2:
		fmt.Println("[2: K8s-cluster version (requires a K8s cluster with Helm 3 enabled)] selected.")
		tempStr = "k8scluster"
	default:
		fmt.Println("You should choose between 1 and 2.")
		return
	}

	err := ioutil.WriteFile("./CB_OPERATOR_MODE", []byte(tempStr), os.FileMode(644))
	errCheck(err)

	fmt.Println("")
	fmt.Println("CB_OPERATOR_MODE is set to: " + tempStr)
	fmt.Println("To change CB_OPERATOR_MODE, just delete the CB_OPERATOR_MODE file and re-run the cb-operator.")
}

func readMode() string {
	if _, err := os.Stat("./CB_OPERATOR_MODE"); err == nil {
		// if file exists
		data, err := ioutil.ReadFile("./CB_OPERATOR_MODE")
		errCheck(err)

		CB_OPERATOR_MODE = string(data)
		fmt.Println("CB_OPERATOR_MODE: " + CB_OPERATOR_MODE)

		//if CB_OPERATOR_MODE == "singlenode" || CB_OPERATOR_MODE == "k8scluster" {
		return CB_OPERATOR_MODE
		//}

	} else if os.IsNotExist(err) == true {
		// path/to/whatever does *not* exist
		fmt.Println("CB_OPERATOR_MODE file does not exist.")
		scanAndWriteMode()
		result := readMode()
		return result

	} else {
		// Schrodinger: file may or may not exist. See err for details.

		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence

		errCheck(err)
		return ""
	}
	return ""
}

var CB_OPERATOR_MODE string

func main() {

	mode := readMode()

	if mode == "singlenode" {
		cmd.Execute()
	} else if mode == "k8scluster" {

	} else {
		fmt.Println("Invalid CB_OPERATOR_MODE: " + mode)
		fmt.Println("CB_OPERATOR_MODE should be one of these: singlenode, k8scluster")

		//fmt.Println("To change CB_OPERATOR_MODE, just delete the CB_OPERATOR_MODE file and re-run the cb-operator.")
		scanAndWriteMode()
		main()
	}
}
