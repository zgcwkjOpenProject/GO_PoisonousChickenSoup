package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	_ "github.com/zgcwkj/poisonousChickenSoup/routers"
)

func main() {
	if runtime.GOOS == "windows" { // Windwos
		//调用浏览器打开网站
		exec.Command(`cmd`, `/c`, `start`, `http://localhost`).Start()
	} else if runtime.GOOS == "darwin" { // MAC
		exec.Command(`open`, `http://localhost`).Start()
	} else if runtime.GOOS == "linux" { // Linux
		exec.Command(`xdg-open`, `http://localhost`).Start()
	}
	//输出浏览地址
	fmt.Println("访问 http://localhost 打开预览")
	//监听端口
	http.ListenAndServe(":80", nil)
}
