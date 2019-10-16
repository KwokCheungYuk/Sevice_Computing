package src

import (
	"fmt"  
    "os"
    "strings"
	"log"
	"../entity"
)

var login bool
var logFile *os.File
var hostName, hostPassword string

func IsLogin() bool {
	return login
}

func Init(){
	entity.Init()
	temp,err := os.OpenFile("data/agenda.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	logFile = temp
	if err != nil {
        log.Fatalln("Open file error!")
	}
	host := entity.ReadHost()
	if (len(host) == 0) {
		login = false
	} else {
		login = true
		hostName = strings.Replace(host[0],"\n","",-1)
	}
}

func RegisterUser(name string, password string, email string, phone string) {
	debugLog := log.New(logFile,"[Operation]",log.LstdFlags)
	i := entity.RegisterUser(name, password, email, phone)
	if (i) {
		debugLog.Println(name, " register successfully!")
	} else {
		debugLog.Println(name, " register failed!")
	}
	defer logFile.Close()
}

func Login(name string, password string) {
	debugLog := log.New(logFile,"[Operation]",log.LstdFlags)
	if entity.CheckUserExist(name) {
		hostName = name
		hostPassword = password
		tempUser := entity.FindUser(name)
		if (tempUser.GetPassword() != password) {
			debugLog.Println(name, " log in failed! Wrong password!")
			fmt.Println("Password is wrong!")
		} else {
			debugLog.Println(name, " log in successfully!")
			fmt.Println("Log in successfully! Welcome to Agenda!")
			entity.WriteHost(name)
		}
	} else {
		debugLog.Println(name, " log in failed! The username does not exist.")
		fmt.Println("The username does not exist!")
	}
	defer logFile.Close()
}

func LogOut(){
	debugLog := log.New(logFile,"[Operation]",log.LstdFlags)
	debugLog.Println(hostName, " log out successfully!")
	entity.LogOut()
	fmt.Println("Log out successfully!")
	defer logFile.Close()
}

func Delete(){
	os.Truncate("data/agenda.log", 0)
	entity.Delete()
}