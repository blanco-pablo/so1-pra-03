package main

import (
	"fmt"
	"testing"
)



func TestAlgo(t *testing.T){

	//fileInfos := getNumericFileInfos()
	//
	//for _, fileInfo := range fileInfos {
	//	println(fileInfo.Name())
	//}

	processes := getLinuxProcesses()

	for _, p := range processes {
		//fmt.Printf("pid: %s - name: %s - state: %s\n", p.pid, p.name, p.state)
	}

	//t.Error("fallo algo :(")
}


// t.Error
// t.Fail
// t.Errorf
// t.log

//go test
