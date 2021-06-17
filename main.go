package main

import (
	"zeus/models"
	_ "zeus/routers"
	"encoding/gob"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

// https://beego.me/docs/mvc/controller/session.md
//因为 session 内部采用了 gob 来注册存储的对象，例如 struct，所以如果你采用了非 memory 的引擎，请自己在 main.go 的 init 里面注册需要保存的这些结构体，不然会引起应用重启之后出现无法解析的错误
// 错误：gob: name not registered for interface: "beegoxiaomi/models.Manager"
func init() {
	gob.Register(models.Manager{})
}
func main() {
	//注册模板函数
	beego.AddFuncMap("unixToDate", models.UnixToDate)
	beego.AddFuncMap("formatImg", models.FormatImg)
	beego.AddFuncMap("formatAttr", models.FormatAttr)
	beego.AddFuncMap("mul", models.Mul)

	//配置session保存在redis
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = beego.AppConfig.String("redisConn")

	beego.Run()
	defer models.DB.Close()
}
