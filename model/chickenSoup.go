package model

import (
	"log"
	"time"
)

// chickenSoup 映射 chickenSoup 表的结构数据
type chickenSoup struct {
	ID         int       `json:"id"`          //主键
	Title      string    `json:"title"`       //标题
	Content    string    `json:"content"`     //内容
	Awesome    int       `json:"awesome"`     //点赞次数
	Hits       int       `json:"hits"`        //点击次数
	Creaator   string    `json:"creaator"`    //创建者信息
	IsDelete   bool      `json:"is_delete"`   //删除否
	CreateTime time.Time `json:"create_time"` //创建时间
}

// GetRow 获取一行数据
// 随机获取 chickensoup 表的一行数据
func GetRow() chickenSoup {
	// rows, err := Db.Query("SELECT * FROM chickensoup ORDER BY RAND() LIMIT 1")
	// if err != nil {
	// 	log.Fatalln(rows, err)
	// }
	// log.Println(rows)
	// defer rows.Close()
	row := Db.QueryRow("SELECT * FROM chickensoup WHERE is_delete != 1 ORDER BY RAND() LIMIT 1") //随机查询一条数据
	chickensoup := chickenSoup{}
	err := row.Scan(&chickensoup.ID,
		&chickensoup.Title,
		&chickensoup.Content,
		&chickensoup.Awesome,
		&chickensoup.Hits,
		&chickensoup.Creaator,
		&chickensoup.IsDelete,
		&chickensoup.CreateTime)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(chickensoup.Content)

	chickensoup.Hits++
	Db.Exec("UPDATE chickensoup SET hits = ? WHERE id = ? AND is_delete != 1", chickensoup.Hits, chickensoup.ID) //增加点击次数

	return chickensoup
}
