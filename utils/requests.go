package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"jenkins/structures"
	"log"
	"net/http"
	"strings"
)

var jenkinsEndpoint string = "https://jenkins-prd.prd.betfair/"
var jenkinsUser string
var jenkinsToken string

//Function that does a GET request to jenkins endpoint. Return 200 OK or 400 error
func jenkins_request(jenkinsEndpoint string) *http.Response {
	var hash = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", jenkinsUser, jenkinsToken)))
	client := &http.Client{}
	req, err := http.NewRequest("GET", jenkinsEndpoint, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", hash))
	resp, _ := client.Do(req)
	return resp
}

/*Function that recibes as argument the TLA and Package number and returns the chef job name
and chef build number for the specified TLA and package number*/
func Get_latest_build_chef(tla string, buildNumber int) (string, string) {
	resp := jenkins_request(fmt.Sprintf("%s/job/%s_package/%d/api/json", jenkinsEndpoint, tla, buildNumber))
	var chef_number, chef_job string
	builds, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var Obj structures.LastBuild
	json.Unmarshal(builds, &Obj)
	for i := range Obj.Actions {
		for j := range Obj.Actions[i].Parameters {
			if strings.Contains(Obj.Actions[i].Parameters[j].JobName, "chef") {
				chef_number = Obj.Actions[i].Parameters[j].Number
				chef_job = Obj.Actions[i].Parameters[j].JobName
			}
		}
	}
	return chef_number, chef_job
}

/*Function that recibes as argument the TLA and Package number and returns the i2 job name
and i2 build number for the specified TLA and package number*/
func Get_latest_build_i2(tla string, buildNumber int) (string, string) {
	resp := jenkins_request(fmt.Sprintf("%s/job/%s_package/%d/api/json", jenkinsEndpoint, tla, buildNumber))
	var i2_number, i2_job_name string
	builds, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var Obj structures.LastBuild
	json.Unmarshal(builds, &Obj)
	for i := range Obj.Actions {
		for j := range Obj.Actions[i].Parameters {
			if strings.Contains(Obj.Actions[i].Parameters[j].JobName, "i2") {
				i2_number = Obj.Actions[i].Parameters[j].Number
				i2_job_name = Obj.Actions[i].Parameters[j].JobName
			}
		}
	}
	return i2_number, i2_job_name
}

/*Function that GETs the chef package commit messages*/
func Get_messages_chef(chef_job string, chef_number string) string {
	//var chef_number, chef_job = Get_latest_build_chef("cds", 316)
	var msg string
	resp := jenkins_request(fmt.Sprintf("%s/job/%s/%s/api/json", jenkinsEndpoint, chef_job, chef_number))
	//teste resposta http
	//fmt.Println(resp)
	messages, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var ObjMessages structures.Messages
	json.Unmarshal(messages, &ObjMessages)

	for j := range ObjMessages.ChangeSet.Items {
		msg = ObjMessages.ChangeSet.Items[j].Comment
	}
	return msg
}

/*Function that GETs the i2 package commit messages*/
func Get_messages_i2(i2_job string, i2_number string) string {
	var msg string
	resp := jenkins_request(fmt.Sprintf("%s/job/%s/%s/api/json", jenkinsEndpoint, i2_job, i2_number))
	//teste resposta http
	//fmt.Println(resp)
	messages, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var ObjMessages structures.Messages
	json.Unmarshal(messages, &ObjMessages)

	for j := range ObjMessages.ChangeSet.Items {
		msg = ObjMessages.ChangeSet.Items[j].Comment
		//msg = strings.TrimSuffix(msg, "\n")
	}
	return msg
}

/*Function that gets the last promoted package*/
func GetLastCompletedBuild(tla string) int {
	resp := jenkins_request(fmt.Sprintf("%s/job/%s_package/promotion/process/ie1-prd-promoted/api/json", jenkinsEndpoint, tla))
	builds, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var ObjBuilds structures.PromotedPackage
	json.Unmarshal(builds, &ObjBuilds)
	return ObjBuilds.LastCompletedBuild.Number
}

/*Function that returns the PROD package number*/
func GetProdPackage(tla string) int {
	lastCompletedBuild := GetLastCompletedBuild(tla)
	resp := jenkins_request(fmt.Sprintf("%s/job/%s_package/promotion/process/ie1-prd-promoted/%d/api/json", jenkinsEndpoint, tla, lastCompletedBuild))
	packages, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var ObjPackages structures.ProdPackage
	json.Unmarshal(packages, &ObjPackages)
	return ObjPackages.Target.Number
}
