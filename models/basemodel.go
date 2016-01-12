/*
Author: Aosen
Data: 2016-01-11
QQ: 316052486
Desc: models的基类，所有其他model都继承此类, 使用beego提供的orm
*/
package models

import (
	"log"

	"github.com/aosen/utils"
	"github.com/astaxie/beego/orm"
)

func init() {
	//读取配置文件信息
	settings := func(path string) map[string]string {
		return utils.NewConfig().Load(path).GlobalContent()
	}("conf/app.conf")
	//获取settings中的信息
	dbinfo, ok := utils.GetSetting(settings, "DBINFO")
	if !ok {
		log.Fatal("not found DBINFO in config file")
	}
	//ORM 必须注册一个别名为 default 的数据库，作为默认使用。
	// 参数1        数据库的别名，用来在ORM中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	orm.RegisterDataBase("default", "mysql", DBINFO)
	//根据数据库的别名，设置数据库的最大空闲连接
	orm.SetMaxIdleConns("default", 30)
	//根据数据库的别名，设置数据库的最大数据库连接 (go >= 1.2)
	orm.SetMaxOpenConns("default", 30)
}

type BaseModel struct {
}
