package models

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func init() {
	mysqladmin := beego.AppConfig.String("mysqladmin")
	mysqlpwd := beego.AppConfig.String("mysqlpwd")
	mysqldb := beego.AppConfig.String("mysqldb")
	mysqlip := beego.AppConfig.String("mysqlip")
	mysqlport := beego.AppConfig.String("mysqlport")
	/*
			如果你想指定主机，你需要使用 (). 例如:

		user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local
	*/
	DB, err = gorm.Open("mysql", mysqladmin+":"+mysqlpwd+"@tcp("+mysqlip+":"+mysqlport+")/"+mysqldb+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		beego.Error(err)
	}
	DB.LogMode(true)

}
