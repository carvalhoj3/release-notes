package utils

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

//tla name
var tla string

//package number
var atual_package int
var at_package string
var package_released int

//lists
var chef_messages []string
var i2_messages []string
var messages []string
var totalMessages []string

//Cobra is built on a structure of commands, arguments & flags.
var cmdRoot = &cobra.Command{
	Use:     "release-notes",
	Short:   "Generate release changes to be deployed in a release",
	Example: "go run main.go --tla" + " cds" + " --package 300" + " --packageR 310",
	Version: "1.0",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if at_package != "" {
			atual_package, _ = strconv.Atoi(at_package)
		}
		if atual_package == 0 {
			atual_package = GetProdPackage(tla)
		}
		if len(tla) > 5 || len(tla) <= 0 {
			return fmt.Errorf("Invalid TLA")
		} else if atual_package > package_released {
			return fmt.Errorf("Atual package can't be greatter than Package to be released.")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		jenkins_request(jenkinsEndpoint)

		for i := atual_package; i <= package_released; i++ {
			chef_number, chef_job := Get_latest_build_chef(tla, i)
			i2_number, i2_job := Get_latest_build_i2(tla, i)

			chef_messages = append(chef_messages, Get_messages_chef(chef_job, chef_number))
			i2_messages = append(i2_messages, Get_messages_i2(i2_job, i2_number))
			totalMessages = append(chef_messages, i2_messages...)
		}

		verifiedMessages := make(map[string]bool)
		list := []string{}
		for _, m := range totalMessages {
			if _, value := verifiedMessages[m]; !value {
				verifiedMessages[m] = true
				list = append(list, m)
				sort.Strings(list)
			}
		}
		fmt.Println(list)
		//generates test file and uses geti2 and getchef messages funcions to write to the file.
		// create file
		f, err := os.Create("releases.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		separator := "\n"
		for _, line := range list {
			_, err = f.WriteString(line + separator)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	},
}

//Required flags to pass as argument on CLI
func init() {
	cmdRoot.Flags().StringVar(&tla, "tla", "", "TLA name")
	cmdRoot.MarkFlagRequired("tla")
	cmdRoot.Flags().StringVar(&at_package, "package", "", "Prod Package number")
	//cmdRoot.MarkFlagRequired("package")
	cmdRoot.Flags().IntVar(&package_released, "packageR", 0, "Package to be released number")
	cmdRoot.MarkFlagRequired("packageR")
	cmdRoot.Flags().StringVar(&jenkinsUser, "jenkinsUser", "", "Jenkins User")
	cmdRoot.MarkFlagRequired("jenkinsUser")
	cmdRoot.Flags().StringVar(&jenkinsToken, "jenkinsToken", "", "Jenkins Token")
	cmdRoot.MarkFlagRequired("jenkinsToken")
}

//Funcion to execute or cobra funcions
func Execute() error {
	return cmdRoot.Execute()
}
