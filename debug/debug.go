package debug

import (
	"aimsis/backend/utils/notification"
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Debug struct {
	Mark string
}

func New(mark string) *Debug {
	return &Debug{Mark: mark}
}

func (this *Debug) PrettyJSON(data interface{}) {
	buff, _ := json.Marshal(data)
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, buff, "", "\t")
	// log.Printf("%s: %s", name, string(prettyJSON.Bytes()))
	fmt.Printf("\x1b[33;1mDebugger \x1b[0m")
	fmt.Printf("\x1b[32;1m%s:\x1b[0m\n", this.Mark)
	fmt.Printf("\x1b[32;1m%s\x1b[0m\n", string(prettyJSON.Bytes()))
	fmt.Println("\x1b[33;1m=======================================\x1b[0m")
}

func (this *Debug) PrintError(data ...interface{}) {
	go notification.BackendError(data...)
	var arr []string
	for _, val := range data {
		arr = append(arr, fmt.Sprint(val))
	}
	str := strings.Join(arr, " >>> ")
	fmt.Printf("\x1b[31;1m%v\x1b[0m\n", str)
}

func (this *Debug) PrintRed(data interface{}) {
	fmt.Printf("\x1b[31;1m%s >>> %v\x1b[0m\n", this.Mark, data)
}

func PrettyJSON(data interface{}) string {
	buff, _ := json.Marshal(data)
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, buff, "", "\t")
	// log.Printf("%s: %s", name, string(prettyJSON.Bytes()))
	return fmt.Sprintf("\x1b[32;1m%s\x1b[0m", string(prettyJSON.Bytes()))
}

func PrintError(data ...interface{}) {
	go notification.BackendError(data...)
	var arr []string
	for _, val := range data {
		arr = append(arr, fmt.Sprint(val))
	}
	str := strings.Join(arr, " >>> ")
	fmt.Printf("\x1b[31;1m%v\x1b[0m\n", str)
}

func PrintRed(data ...interface{}) {
	var str []string
	for _, val := range data {
		str = append(str, fmt.Sprint(val))
	}
	fmt.Printf("\x1b[31;1m%v\x1b[0m\n", strings.Join(str, " >> "))
}

func PrintGreen(data ...interface{}) {
	var str []string
	for _, val := range data {
		str = append(str, fmt.Sprint(val))
	}
	fmt.Printf("\x1b[32;1m%v\x1b[0m\n", strings.Join(str, " >> "))
}

func PrintYellow(data interface{}) {
	fmt.Printf("\x1b[33;1m%v\x1b[0m\n", data)
}
