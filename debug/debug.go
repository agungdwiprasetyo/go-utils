package debug

import (
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
	fmt.Printf("\x1b[33;1mDebugger: \x1b[0m")
	fmt.Printf("\x1b[33;1m%s:\x1b[0m\n", this.Mark)
	fmt.Printf("\x1b[32;1m%s\x1b[0m\n", string(prettyJSON.Bytes()))
	fmt.Println("\x1b[33;1m=======================================\x1b[0m")
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

func StringRed(data ...interface{}) string {
	var str []string
	for _, val := range data {
		str = append(str, fmt.Sprint(val))
	}
	return fmt.Sprintf("\x1b[31;1m%v\x1b[0m\n", strings.Join(str, " >> "))
}

func StringYellow(data ...interface{}) string {
	var str []string
	for _, val := range data {
		str = append(str, fmt.Sprint(val))
	}
	return fmt.Sprintf("\x1b[33;1m%v\x1b[0m\n", strings.Join(str, " >> "))
}

func StringGreen(data ...interface{}) string {
	var str []string
	for _, val := range data {
		str = append(str, fmt.Sprint(val))
	}
	return fmt.Sprintf("\x1b[32;1m%v\x1b[0m\n", strings.Join(str, " >> "))
}

func PrintRed(data ...interface{}) {
	str := StringRed(data...)
	fmt.Printf(str)
}

func PrintYellow(data ...interface{}) {
	fmt.Printf("\x1b[33;1m%v\x1b[0m\n", StringYellow(data...))
}

func PrintGreen(data ...interface{}) {
	fmt.Printf(StringGreen(data...))
}
