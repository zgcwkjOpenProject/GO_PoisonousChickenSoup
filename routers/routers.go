package routers

import (
	"net/http"

	"github.com/zgcwkj/poisonousChickenSoup/controller"
)

func init() {
	//配置路由
	http.HandleFunc("/", controller.Index)

}
