package main

import (
	"./server"
	"os"
	"github.com/spf13/pflag"
)


func main(){
	port := os.Getenv("PORT")

	//设置默认监听端口
	if len(port) == 0{
		port = "4869"
	}

	//从命令行中获取指定的监听端口
	curPort := pflag.StringP("port", "p", "4869", "Port for listening")
	pflag.Parse()
	if len(*curPort) != 0{
		port = *curPort
	}

	//启动服务器
	server.Begin(port)
}