package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)
import "log"

func getNumericFileInfos() []os.FileInfo {

	//fmt.Print("Estoy a punto de imprimir algo");

	files, err := ioutil.ReadDir("/proc")

	if err != nil {
		log.Fatal(err)
	}

	var numberProcs []os.FileInfo

	for _, f := range files {
		_, err := strconv.Atoi(f.Name())
		if err == nil {
			numberProcs = append(numberProcs, f)
		}
	}

	return numberProcs;

}
type LinuxProcessInfo struct {

	pid string
	user string // lo tengo en numeros
	name string

	// falta ram

	state string
	//isZombie bool
	//isSleeping bool
	//isRunning bool
	//isStoped bool
}

func getLinuxProcesses() []map[string]interface{} {
	var linuxProcesses []map[string]interface{}

	numericFileInfos := getNumericFileInfos()

	for _, f := range numericFileInfos {

		file, err := os.Open("/proc/" + f.Name() + "/status")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		fileContentBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		fileContent := fmt.Sprintf("%s", fileContentBytes)

		processInfo := extractLinuxProcessInfo(f.Name(), fileContent)

		linuxProcesses = append(linuxProcesses, processInfo)
		//fmt.Println("asdf")
		//fmt.Printf("%s", fileContent)

		//break
	}

	return linuxProcesses
}

func extractLinuxProcessInfo(pid string, content string) map[string]interface{} {
	//processInfo := new(LinuxProcessInfo)


	processInfo := make(map[string]interface{})

	processInfo["pid"] = pid;

	var line = 0
	var saveChars = false
	var value = ""

	for _, c := range content {
		if saveChars && c != '\n' {
			value += string(c)
		}
		if c == ':' {
			saveChars = true
		}
		if c == '\n' {

			switch line {
			case 0:
				processInfo["name"] = strings.TrimSpace(value)
				break
			case 2:
				processInfo["state"] = strings.TrimSpace(value)
				break
			case 8:
				processInfo["uid"] = strings.TrimSpace(value)
				break
			}

			line += 1
			saveChars = false
			value = ""
		}
	}


	return processInfo;
}