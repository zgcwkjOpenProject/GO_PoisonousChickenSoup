package controller

import (
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"github.com/zgcwkj/poisonousChickenSoup/model"
)

// Index 默认页面
func Index(resp http.ResponseWriter, req *http.Request) {
	chickenSoup := model.GetRow()                     //查询数据库的数据
	t, _ := template.ParseFiles("./views/index.html") //解析模板文件
	t.Execute(resp, chickenSoup)                      //执行模板的merger操作
}

// Help 帮助页面
func Help(resp http.ResponseWriter, req *http.Request) {
	f, _ := os.Open("./views/help.html") //读取文件
	buf, _ := ioutil.ReadAll(f)          //将文件转换成数据
	resp.Write(buf)                      // 输出到页面
}
