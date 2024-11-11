package routers

import (
	"net/http"

	"poisonousChickenSoup/controller"
)

// init 初始化路由
func init() {
	http.HandleFunc("/", controller.Index)                                                       //默认页面
	http.HandleFunc("/help", controller.Help)                                                    //帮助页面
	http.HandleFunc("/api", controller.Api)                                                      //接口
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))) //静态文件
}
