package DB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//数据库配置
const (
	username = "root"
	password = "123qwe"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "train"
)

// DB数据库连接池
var db *sql.DB

func InitDB() {
	//构建连接："用户名：密码@tcp(IP:端口)/数据库？charset=utf8"
	path := strings.Join([]string{username, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库，前者是驱动名，所以要导入：	_ "github.com/go-sql-driver/mysql"
	db, _ = sql.Open("mysql", path)
	//设置数据库最大限制连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	//验证连接
	if err := db.Ping(); err != nil {
		fmt.Println("link database fail")
		checkErr(err)
		return
	}
	fmt.Println("connect success")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
