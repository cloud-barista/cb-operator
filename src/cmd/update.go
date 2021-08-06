package cmd

import (
	"fmt"

	"github.com/cloud-barista/cb-operator/src/common"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"apply"},
	Short:   "Update Cloud-Barista System",
	Long:    `Update Cloud-Barista System`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n[Update Cloud-Barista]")
		fmt.Println()

		if common.FileStr == "" {
			fmt.Println("file is required")
		} else {
			common.FileStr = common.GenConfigPath(common.FileStr, common.CBOperatorMode)

			var cmdStr string
			switch common.CBOperatorMode {
			case common.ModeDockerCompose:
				fmt.Println("cb-operator Docker Compose mode does not support 'update/apply' subcommand.")

			case common.ModeKubernetes:
				cmdStr = fmt.Sprintf("helm upgrade --namespace %s --install %s -f %s ../helm-chart", common.CBK8sNamespace, common.CBHelmReleaseName, common.FileStr)
				//fmt.Println(cmdStr)
				common.SysCall(cmdStr)
			default:

			}

		}

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	pf := updateCmd.PersistentFlags()
	pf.StringVarP(&common.FileStr, "file", "f", common.NotDefined, "User-defined configuration file")

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
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
