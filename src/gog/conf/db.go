package conf

import "fmt"

/**
数据库配置
*/

var (
	DBDialect string //数据库方言
	DBConnStr string //链接字符串
	host      string
	port      string
	user      string
	pwd       string
	schema    string //要链接的库
)

func init() {
	DBDialect = "mysql"
	host = ""
	port = "3306"
	user = ""
	pwd = ""
	schema = "alm"
	DBConnStr = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, port, schema)
}
