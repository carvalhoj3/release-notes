package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"text/template"

	"github.com/spf13/cobra"
)

//tla name
var tla string
var date string
var time string
var destination string

//package number
var atual_package int
var package_released int

//lists
var chef_messages []string
var i2_messages []string
var totalMessages []string
var list []string

type Change string

type Release struct {
	Tla      string
	USNumber int
	Date     string
	Time     string
	Changes  []Change
}

func (c Change) String() string {

	rg := regexp.MustCompile("\\[\\w+\\]\\[TP\\-(\\d+)\\]\\s\\-\\s.+")
	str := rg.FindStringSubmatch(string(c))

	if len(str) < 1 {
		return string(c)
	}

	url := "https://ppb.tpondemand.com/entity/" + str[1]

	return "<a href=\"" + url + "\">" + string(c) + "</a>"
}
func GetChanges(list []string) []Change {
	changes := make([]Change, len(list))
	for i := range list {
		changes[i] = Change(list[i])
	}
	return changes

}
func Template() {
	i, _ := strconv.Atoi(tla)
	if date == "" || time == "" {
		fmt.Printf("In order to send the email, you must inform the date and time of the release. \n")
	}
	std1 := Release{
		tla,
		i,
		date,
		time,
		GetChanges(list),
	}
	x, _ := os.Open("utils/template.html")
	b, _ := ioutil.ReadAll(x)
	tmp1, err := template.New("template").Parse(string(b))
	if err != nil {
		fmt.Println(err)
	}
	// // standard output to print merged data
	outFile, err := os.Create("utils/template1.html")
	if err != nil {
		fmt.Println("error")
	}
	// tmp1.Execute(outFile, std1)
	tmp1.Execute(outFile, std1)
}

var config = Config{}

func init() {
	config.Read()
}
func SendEmail() {
	subject := "Test email release notes"

	r := NewRequest([]string{destination}, subject)
	r.Send("utils/template1.html", map[string]string{"username": "kolibri"})
}

//Cobra is built on a structure of commands, arguments & flags.
var cmdRoot = &cobra.Command{
	Use:     "release-notes",
	Short:   "Generate release changes to be deployed in a release",
	Example: "go run main.go --tla" + " cds" + " --package 300" + " --packageR 310",
	Version: "1.0",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// We need to support as string, sice we can pass an empty string as parameter
		// if at_package != "" {
		// 	atual_package, _ = strconv.Atoi(at_package)
		// }
		if reflect.TypeOf(atual_package).Kind() != reflect.Int {
			fmt.Println(atual_package, "atual package should be an integer")
			os.Exit(1)
		}
		if atual_package == -1 {
			atual_package = GetProdPackage(tla)
		}
		// if rl_package != "" {
		// 	package_released, _ = strconv.Atoi(rl_package)
		// }
		if reflect.TypeOf(package_released).Kind() != reflect.Int {
			fmt.Println(atual_package, "atual package should be an integer")
			os.Exit(1)
		}
		if package_released == -1 {
			package_released = GetLastPackage(tla)
		}
		if len(tla) > 5 || len(tla) <= 0 {
			//fmt.Println("Invalid TLA")
			fmt.Println(tla, "is an invalid TLA")
			os.Exit(1)

		} else if atual_package > package_released {
			fmt.Printf("Atual package %d can't be greatter than the package to be released %d \n", atual_package, package_released)
			os.Exit(1)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		jenkins_request(jenkinsEndpoint)

		if atual_package == package_released {
			fmt.Printf("Both packages are equal, nothing to be released.")
			os.Exit(1)
		}
		/* Get chef messages */
		chef_initial := Get_latest_build_chef(tla, atual_package)
		chef_final := Get_latest_build_chef(tla, package_released)

		for j := chef_initial; j <= chef_final; j++ {
			chef_messages = append(chef_messages, Get_messages_chef(j)...)
		}
		for i := atual_package; i <= package_released; i++ {
			i2_number := Get_latest_build_i2(tla, i)
			i2_messages = append(i2_messages, Get_messages_i2(i2_number)...)
		}
		totalMessages = append(chef_messages, i2_messages...)
		verifiedMessages := make(map[string]bool)
		for _, m := range totalMessages {
			if _, value := verifiedMessages[m]; !value {
				verifiedMessages[m] = true
				list = append(list, m)
				sort.Strings(list)
			}
		}

		Template()
		if destination != "" && date != "" && time != "" {
			SendEmail()
		}

		//fmt.Println(list)
		fmt.Println("Click here to check the changes \n https://jenkins-prd.prd.betfair/job/release-notes-generator/ws/release-notes/releases.txt")
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
	cmdRoot.Flags().IntVar(&atual_package, "package", -1, "Prod Package number")
	cmdRoot.Flags().IntVar(&package_released, "packageR", -1, "Package to be released number")
	cmdRoot.Flags().StringVar(&jenkinsUser, "jenkinsUser", "", "Jenkins User")
	cmdRoot.MarkFlagRequired("jenkinsUser")
	cmdRoot.Flags().StringVar(&jenkinsToken, "jenkinsToken", "", "Jenkins Token")
	cmdRoot.MarkFlagRequired("jenkinsToken")
	cmdRoot.Flags().StringVar(&destination, "destination", "", "Mail destination")
	cmdRoot.Flags().StringVar(&date, "date", "", "Date of release")
	cmdRoot.Flags().StringVar(&time, "time", "", "Time of release")

}

//Funcion to execute or cobra funcions
func Execute() error {
	return cmdRoot.Execute()
}
