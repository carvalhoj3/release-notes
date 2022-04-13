package utils

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

//tla name
var tla string

//package number
var pckg string

//Cobra is built on a structure of commands, arguments & flags.
var cmdRoot = &cobra.Command{
	Use:     "release-notes",
	Short:   "Generate release changes to be deployed in a release",
	Example: "go run main.go --tla" + " cds" + " --package 300",
	Version: "1.0",
	Run: func(cmd *cobra.Command, args []string) {
		jenkins_request(jenkinsEndpoint)
		p, err := strconv.Atoi(pckg)
		if err != nil {
			panic(err)
		}
		chef_number, chef_job := Get_latest_build_chef(tla, p)
		i2_number, i2_job := Get_latest_build_i2(tla, p)

		output := Get_messages_chef(chef_job, chef_number) + "\n" + Get_messages_i2(i2_job, i2_number)

		// //generates test file and uses geti2 and getchef messages funcions to write to the file.
		f, err := os.Create("releases.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		_, err2 := f.Write([]byte(output))

		if err2 != nil {
			log.Fatal(err2)
		}
	},
}

//Required flags to pass as argument on CLI
func init() {
	cmdRoot.Flags().StringVar(&tla, "tla", "", "TLA name")
	cmdRoot.MarkFlagRequired("tla")
	cmdRoot.Flags().StringVar(&pckg, "package", "", "Package number")
	cmdRoot.MarkFlagRequired("package")
}

//Funcion to execute or cobra funcions
func Execute() error {
	return cmdRoot.Execute()
}
