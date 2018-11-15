package main

import (
	"os"

	_ "github.com/Liu-YT/cloudgo/service"
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

	// 允许自定义启动的端口，默认端口为8080
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}
	beego.Run(":" + port) // 运行
}
