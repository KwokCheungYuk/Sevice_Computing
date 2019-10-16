package entity

import (
	"fmt"  
    "regexp"
	"os"
	"encoding/json"
	"bufio"
	"io"
) 

var users []User


func UserJsonDecode(js []byte) User{
	var temp User
	err := json.Unmarshal(js, &temp)
	if err != nil {
		fmt.Println("Deconde error")
	}
	return temp
}



func UserJsonEncode(user User) []byte {
	js, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Encode error")
		os.Exit(1)
	}
	return js
}

func WriteUserToFile(user User) {
	file, err := os.OpenFile("data/User.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
    if err != nil {
        fmt.Println("open file failed.", err.Error())
        os.Exit(1)
    }

    file.WriteString(string(UserJsonEncode(user)[:]))
    file.WriteString("\n")

	file.Close()
}

func ReadUserFromFile() []User{
	var tmp []User
	f, err := os.OpenFile("data/User.txt", os.O_CREATE|os.O_RDONLY|os.O_APPEND, 0600)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    rd := bufio.NewReader(f)
    for {
        line, err := rd.ReadString('\n') 
        if err != nil || io.EOF == err {
            break
        }
        tmp = append(tmp, UserJsonDecode([]byte(line)))
    }
    return tmp
}

func ReadHost() []string{
	var tmp []string
	f, err := os.OpenFile("data/Host.txt", os.O_CREATE|os.O_RDONLY|os.O_APPEND, 0600)
    if err != nil {
        panic(err)
    }
    rd := bufio.NewReader(f)
    for {
        line, err := rd.ReadString('\n') 
        if err != nil || io.EOF == err {
            break
        }
        tmp = append(tmp, line)
	}
	f.Close()
    return tmp
}

func WriteHost(name string) {
	file, err := os.Create("data/Host.txt")
    if err != nil {
        fmt.Println("open file failed.", err.Error())
        os.Exit(1)
    }
    file.WriteString(name)
	file.WriteString("\n")
	file.Close()
}

func Init() {
	tmp_u := ReadUserFromFile()
	for i := 0; i < len(tmp_u); i++ {
		users = append(users, tmp_u[i])
	}
}

func CheckEmail(str string) bool {  
    var check bool  
    check, _ = regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", str)  
    return check  
}

func CheckPhone(str string) bool {  
    var check bool  
    check, _ = regexp.MatchString("^1[0-9]{10}$", str)  
    return check  
}  

func CheckUserExist(name string) bool{
	for i := 0; i < len(users); i++ {
		if users[i].Name == name {
			return true
		}
	}
	return false
}

func FindUser(name string) User{
	for i := 0; i < len(users); i++ {
		if users[i].Name == name {
			return users[i]
		}
	}
	return User{"null", "null", "null", "null"}
}

func RegisterUser(name string, password string, email string, phone string) bool {
	var user User
	err := false
	if (CheckEmail(email) == false) {
		fmt.Println("Email is invaild!")
		err = true
	}
	if (CheckPhone(phone) == false) {
		fmt.Println("Phone is invaild!")
		err = true
	}

	if (CheckUserExist(name)) {
		fmt.Println("This username has exited. Please use another username!")
		err = true
	}
	if (err) {
		return false
	}
	user.Name = name
	user.Password = password
	user.Email = email
	user.Phone = phone
	users = append(users,user)
	WriteUserToFile(user)
	fmt.Println("Register successfully!")
	return true
}

func LogOut(){
	os.Truncate("data/Host.txt", 0)
}

func Delete(){
	os.Truncate("data/Host.txt", 0)
	os.Truncate("data/User.txt", 0)
	users = users[0:0]
}