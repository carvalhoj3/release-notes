package main

import (
	"jenkins/utils"
	"testing"
)

func TestGetLastPackage(t *testing.T) {
	var res int = utils.GetLastPackage("kdp")
	if res == 354 {
		t.Errorf("Not the latest package")
	}
}

// func TestGetLastPackageError(t *testing.T) {
// 	var res int = utils.GetLastPackage("kdp")

// 	if res != 1 {
// 		fmt.Errorf("Invalid result")
// 	}
// }

// func TestGetLastPackageInvalid(t *testing.T) {
// 	var res int = utils.GetLastPackage(123)

// 	if res != 353 {
// 		fmt.Println("Got %s")
// 	}
// }
