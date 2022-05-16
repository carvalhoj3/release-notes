package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestJenkinsRequest(t *testing.T) {
	if jenkins_request(jenkinsEndpoint).StatusCode != 200 {
		fmt.Errorf(jenkins_request(jenkinsEndpoint).Status)
	}
}
func TestTlaParameter(t *testing.T) {
	tla = "kdp"
	var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	var IsValidTla = regexp.MustCompile("^[a-zA-Z]{3}$|^[a-zA-Z]{5}$").MatchString
	fmt.Println(IsValidTla(tla))
	if !IsLetter(tla) {
		t.Error(tla, "is not a valid TLA, can't contain any integer value.")
	} else if !IsValidTla(tla) {
		t.Error(tla, "should've exacly 3 or 5 characters.")
	} else if tla == "" {
		t.Errorf("Not a valid TLA.")
	}
}
func TestPackagesValues(t *testing.T) {
	atual_package = 3
	package_released = 4

	if reflect.TypeOf(atual_package).Kind() != reflect.Int || atual_package <= 0 {
		t.Error(atual_package, "isn't an integer value or is <= 0.")
	}
	if reflect.TypeOf(package_released).Kind() != reflect.Int || package_released <= 0 {
		t.Error(atual_package, "isn't an integer value or is <= 0.")
	}
	if atual_package > package_released || atual_package == package_released {
		t.Error("Package in production: ", atual_package, "Package to be released: ", package_released, "Package in production can't be equal or greather than Package to be Released.")
	}
}

func TestGetLastPackage(t *testing.T) {
	jenkins_request(jenkinsEndpoint)
	var res int = GetLastPackage("kdp")
	if res != 353 {
		t.Errorf("Not the latest package")
	}
}

func TestXxx(t *testing.T) {
	jenkins_request(jenkinsEndpoint)
	var res int = GetProdPackage("kdp")
	if res != 351 {
		t.Error(res, "is not the package in production.")
	}
}
