/*
Author: Aosen
Data: 2016-01-11
QQ: 316052486
Desc: 小说服务器主入口, 为了防止web api被恶意使用，在配置文件中增加appid， appsecret
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"novel/routers"
	"novel/tasks"
	"novel/utils"

	"github.com/aosen/goutils"
	"github.com/gorilla/mux"

	_ "novel/models"
)

func loadconf(path string) (settings map[string]string) {
	//生成配置文件对象,加载配置文件
	config := goutils.NewConfig().Load(path)
	return config.GlobalContent()
}

func main() {
	//配置文件信息
	settings := loadconf("conf/app.conf")
	if key, ok := utils.CheckSettings(settings); !ok {
		log.Fatal(fmt.Sprintf("not found %s in config file", key))
	}
	//启动系统任务
	go tasks.SysTask(settings)
	port, _ := settings["PORT"]
	host, _ := settings["HOST"]
	web := goutils.NewWeb(settings)
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
