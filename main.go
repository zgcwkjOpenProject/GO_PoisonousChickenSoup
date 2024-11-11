package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	_ "poisonousChickenSoup/routers"
)

func main() {
	//调用浏览器打开网站
	if runtime.GOOS == "windows" { // Windwos
		exec.Command(`cmd`, `/c`, `start`, `http://localhost`).Start()
	} else if runtime.GOOS == "darwin" { // Mac
		exec.Command(`open`, `http://localhost`).Start()
	} else if runtime.GOOS == "linux" { // Linux
		exec.Command(`xdg-open`, `http://localhost`).Start()
	}
	//输出浏览地址
	fmt.Println("访问 http://localhost 打开预览")
	//监听端口
	http.ListenAndServe(":80", nil)
}
