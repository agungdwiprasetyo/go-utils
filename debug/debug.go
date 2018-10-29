package debug

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"time"
)

func getCaller(i int) string {
	_, file, line, _ := runtime.Caller(i)
	paths := strings.Split(file, "/")
	n := len(paths) - 2
	if n < 0 {
		n = 0
	}
	path := strings.Join(paths[n:], "/")
	return fmt.Sprintf("%s:%d", path, line)
}

// Println for show data with stack trace
func Println(data ...interface{}) {
	var str []string
	for _, val := range data {
		valStr := fmt.Sprint(val)
		if valStr == "" {
			valStr = stringYellow(1, "<empty_string>")
		} else {
			valStr = stringRed(6, valStr)
		}
		str = append(str, valStr)
	}
	messages := strings.Join(str, stringYellow(1, " | "))

	// fmt.Println(stringYellow(5, "==================================================="))
	fmt.Println(stringGreen(1, fmt.Sprintf("Trace\t : %s", getCaller(2))))
	fmt.Println(stringGreen(1, fmt.Sprintf("Time\t : %v", time.Now())))
	fmt.Println(stringGreen(1, "Data\t :"))
	fmt.Println(stringRed(6, messages))
	fmt.Println(stringYellow(5, "---------------------------------------------------"))
}

// PrintJSON for show data in pretty JSON with stack trace
func PrintJSON(data interface{}) {
	buff, _ := json.Marshal(data)
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, buff, "", "     ")
	fmt.Println(stringYellow(5, "---------------------------------------------------"))
	fmt.Println(stringRed(6, fmt.Sprintf("Trace\t : %s", getCaller(2))))
	fmt.Println(stringRed(6, fmt.Sprintf("Time\t : %v", time.Now())))
	fmt.Println(stringRed(6, "Data\t :"))
	fmt.Println(stringGreen(1, string(prettyJSON.Bytes())))
	fmt.Println(stringYellow(5, "---------------------------------------------------"))
}

func stringRed(fontType int, data interface{}) string {
	return fmt.Sprintf("\x1b[31;%dm%+v\x1b[0m", fontType, data)
}

func stringYellow(fontType int, data interface{}) string {
	return fmt.Sprintf("\x1b[33;%dm%+v\x1b[0m", fontType, data)
}

func stringGreen(fontType int, data interface{}) string {
	return fmt.Sprintf("\x1b[32;%dm%+v\x1b[0m", fontType, data)
}
