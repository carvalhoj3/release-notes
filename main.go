package main

import (
	"jenkins/utils"
)

// var jenkinsEndpoint string = "https://jenkins-prd.prd.betfair/"
// var jenkinsUser string = "carvalhoj3"
// var jenkinsToken string = "11035844f6391d96a2d329b468f17ebe7c"

// func jenkins_request(jenkinsEndpoint string) *http.Response {
// 	var hash = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", jenkinsUser, jenkinsToken)))
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", jenkinsEndpoint, nil)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", hash))
// 	resp, _ := client.Do(req)
// 	return resp
// }
// func Get_latest_build_chef(tla string, buildNumber int) (string, string) {
// 	resp := jenkins_request(fmt.Sprintf("%s/job/%s_package/%d/api/json", jenkinsEndpoint, tla, buildNumber))
// 	var chef_number, chef_job string
// 	builds, _ := io.ReadAll(resp.Body)
// 	defer resp.Body.Close()
// 	var Obj structures.LastBuild
// 	json.Unmarshal(builds, &Obj)
// 	for i := range Obj.Actions {
// 		for j := range Obj.Actions[i].Parameters {
// 			if strings.Contains(Obj.Actions[i].Parameters[j].JobName, "chef") {
// 				chef_number = Obj.Actions[i].Parameters[j].Number
// 				chef_job = Obj.Actions[i].Parameters[j].JobName
// 			}
// 		}
// 	}
// 	return chef_number, chef_job
// }

// func Get_latest_build_i2(tla string, buildNumber int) (string, string) {
// 	resp := jenkins_request(fmt.Sprintf("%s/job/%s_package/%d/api/json", jenkinsEndpoint, tla, buildNumber))
// 	var i2_number, i2_job_name string
// 	builds, _ := io.ReadAll(resp.Body)
// 	defer resp.Body.Close()
// 	var Obj structures.LastBuild
// 	json.Unmarshal(builds, &Obj)
// 	for i := range Obj.Actions {
// 		for j := range Obj.Actions[i].Parameters {
// 			if strings.Contains(Obj.Actions[i].Parameters[j].JobName, "i2") {
// 				i2_number = Obj.Actions[i].Parameters[j].Number
// 				i2_job_name = Obj.Actions[i].Parameters[j].JobName
// 			}
// 		}
// 	}
// 	return i2_number, i2_job_name
// }

// func Get_messages_chef(chef_job string, chef_number string) string {
// 	//var chef_number, chef_job = Get_latest_build_chef("cds", 316)
// 	var msg string
// 	resp := jenkins_request(fmt.Sprintf("%s/job/%s/%s/api/json", jenkinsEndpoint, chef_job, chef_number))
// 	//teste resposta http
// 	//fmt.Println(resp)
// 	messages, _ := io.ReadAll(resp.Body)
// 	defer resp.Body.Close()
// 	var ObjMessages structures.Messages
// 	json.Unmarshal(messages, &ObjMessages)

// 	for j := range ObjMessages.ChangeSet.Items {
// 		msg = ObjMessages.ChangeSet.Items[j].Comment
// 	}
// 	return msg
// }

// func Get_messages_i2(i2_job string, i2_number string) string {
// 	//var i2_number, i2_job_name = Get_latest_build_i2("cds", 316)
// 	var msg string
// 	resp := jenkins_request(fmt.Sprintf("%s/job/%s/%s/api/json", jenkinsEndpoint, i2_job, i2_number))
// 	//teste resposta http
// 	//fmt.Println(resp)
// 	messages, _ := io.ReadAll(resp.Body)
// 	defer resp.Body.Close()
// 	var ObjMessages structures.Messages
// 	json.Unmarshal(messages, &ObjMessages)

// 	for j := range ObjMessages.ChangeSet.Items {
// 		msg = ObjMessages.ChangeSet.Items[j].Comment
// 	}
// 	return msg
// }

func main() {
	utils.Execute()

}
