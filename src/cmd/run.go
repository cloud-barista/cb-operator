package cmd

import (
	"fmt"
	"strings"

	"github.com/cloud-barista/cb-operator/src/common"
	"github.com/spf13/cobra"
)

var mode string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Setup and Run Cloud-Barista System",
	Long:  `Setup and Run Cloud-Barista System`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n[Setup and Run Cloud-Barista]")
		fmt.Println()

		if common.FileStr == "" {
			fmt.Println("file is required")
		} else {
			common.FileStr = common.GenConfigPath(common.FileStr, common.CBOperatorMode)

			var cmdStr string
			switch common.CBOperatorMode {
			case common.ModeDockerCompose:
				cmdStr = fmt.Sprintf("COMPOSE_PROJECT_NAME=%s docker-compose -f %s up", common.CBComposeProjectName, common.FileStr)
				//fmt.Println(cmdStr)
				common.SysCall(cmdStr)
			case common.ModeKubernetes:
				// For Kubernetes 1.19 and above
				cmdStr = fmt.Sprintf("kubectl create ns %s --dry-run=client -o yaml | kubectl apply -f -", common.CBK8sNamespace)
				// For Kubernetes 1.18 and below
				//cmdStr = fmt.Sprintf("kubectl create ns %s --dry-run -o yaml | kubectl apply -f -", common.CBK8sNamespace)
				common.SysCall(cmdStr)

				cmdStr = fmt.Sprintf("helm install --namespace %s %s -f %s ../helm-chart --debug", common.CBK8sNamespace, common.CBHelmReleaseName, common.FileStr)
				if strings.ToLower(mode) == "gcp" || strings.ToLower(mode) == "gke" {
					cmdStr += " --set metricServer.enabled=false"
				}
				//fmt.Println(cmdStr)
				common.SysCall(cmdStr)
			default:

			}

		}

	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	pf := runCmd.PersistentFlags()
	pf.StringVarP(&common.FileStr, "file", "f", common.NotDefined, "User-defined configuration file")
	pf.StringVarP(&mode, "mode", "", common.NotDefined, "Cloud Service Provider / Kind of Managed K8s services")

	/*
		switch common.CBOperatorMode {
		case common.ModeDockerCompose:
			pf.StringVarP(&common.FileStr, "file", "f", "../docker-compose-mode-files/docker-compose.yaml", "Path to Cloud-Barista Docker Compose YAML file")
		case common.ModeKubernetes:
			pf.StringVarP(&common.FileStr, "file", "f", "../helm-chart/values.yaml", "Path to Cloud-Barista Helm chart file")
		default:

		}
	*/

	//	cobra.MarkFlagRequired(pf, "file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
