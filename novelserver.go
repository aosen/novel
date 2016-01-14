/*
Author: Aosen
Data: 2016-01-11
QQ: 316052486
Desc: 小说服务器主入口
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"novel/routers"

	"github.com/aosen/utils"
	"github.com/gorilla/mux"

	_ "novel/models"
)

func loadconf(path string) (settings map[string]string) {
	//生成配置文件对象,加载配置文件
	config := utils.NewConfig().Load(path)
	return config.GlobalContent()
}

func main() {
	//配置文件信息
	settings := loadconf("conf/app.conf")
	port, ok := utils.GetSetting(settings, "PORT")
	if !ok {
		log.Fatal("not found PORT in config file")
	}
	host, ok := utils.GetSetting(settings, "HOST")
	if !ok {
		log.Fatal("not found HOST in config file")
	}
	web := utils.NewWeb(settings)
	log.Printf("server run on %s:%s", host, port)
	// Routes consist of a path and a handler function.
	r := mux.NewRouter()
	routers.Register(web, r)
	http.Handle("/", r)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
