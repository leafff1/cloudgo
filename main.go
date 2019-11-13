package main

import (
	"os"
	"./server"
	"github.com/spf13/pflag"
)

func main() {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8888"
	}//设置默认监听端口为8888

	curPort := pflag.StringP("port", "p", "8888", "Port for listening")
	pflag.Parse()
	if len(*curPort) != 0 {
		port = *curPort
	}//从命令行中获取指定的监听端口

	server.Begin(port)//启动服务器
}
