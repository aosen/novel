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

	_ "github.com/go-sql-driver/mysql"
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
}

type BaseModel struct {
}
