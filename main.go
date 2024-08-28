package main

import (
	"bytes"
	"forward/gocron"
	"github.com/go-vgo/robotgo"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var StrFirst string

func getMousePoint() string {
	a, b := robotgo.GetMousePos()
	return strconv.Itoa(a) + strconv.Itoa(b)
}
func run() {
	str := getMousePoint()
	strTemp := ""
	strPowerPlan := getPowerPlan()
	//log.Println(StrFirst + "??" + str)
	if StrFirst == str {
		//log.Println(StrFirst)
		if strPowerPlan != "381b4222-f694-41f0-9685-ff5bb260df2e" {
			strTemp = `D:\AServer\service\setBalance.bat`
			log.Println("setBalance")
		}
	} else {
		if strPowerPlan != "e9a42b02-d5df-448d-aa00-03f14749eb61" {
			strTemp = `D:\AServer\service\setPerformance.bat`
			log.Println("setPerformance")
		}
		//log.Println("changed")
	}
	StrFirst = str
	if strTemp != "" {
		var out bytes.Buffer
		cmd := exec.Command("cmd.exe", "/C", strTemp, "qqq")
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Println(err)
			panic(err)
			return
		}
		//log.Println(out.String())
	} else {
		log.Println("no need change")
	}
}
func main() {
	f, err := os.OpenFile("setPower.log", os.O_CREATE|os.O_APPEND|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return
	}
	defer func() {
		f.Close()
	}()

	// 组合一下即可，os.Stdout代表标准输出|
	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	StrFirst = getMousePoint()
	gocron.Every(3).Minutes().Do(run)
	<-gocron.Start()
}
func getPowerPlan() string {
	var out bytes.Buffer
	cmd := exec.Command("cmd.exe", "/C", `D:\\AServer\\service\\getPowerPlan.bat`, "qqq")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		panic(err)
		return ""
	} else {
		r := strings.Split(out.String(), " ")
		return strings.TrimSpace(r[4])
	}
}
func ExistsFile(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
