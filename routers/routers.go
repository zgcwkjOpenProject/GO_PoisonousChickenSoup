package routers

import (
	"net/http"

	"github.com/zgcwkj/poisonousChickenSoup/controller"
)

// init 初始化路由
func init() {
	http.HandleFunc("/", controller.Index)    //默认页面
	http.HandleFunc("/help", controller.Help) //帮助页面
	http.HandleFunc("/api", controller.Api)   //接口
}
