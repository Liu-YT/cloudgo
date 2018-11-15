package main

import (
	"os"

	_ "github.com/Liu-YT/cloudgo/service" // 引入自定义设置的路由设置等
	"github.com/astaxie/beego"
	flag "github.com/spf13/pflag"
)

const (
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	// 允许自定义服务器启动的端口，默认端口为8080
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}
	beego.Run(":" + port) // 运行
}
