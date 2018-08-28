package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id          int
	Name        string
}


func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:@/test?charset=utf8")

	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))

	orm.RunSyncdb("default", false, true)
}