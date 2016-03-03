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
	"github.com/aosen/search"
	"github.com/aosen/search/indexer"
	"github.com/aosen/search/ranker"
	"github.com/aosen/search/segmenter"
	"github.com/gorilla/mux"

	_ "novel/models"
)

func loadconf(path string) (settings map[string]string) {
	//生成配置文件对象,加载配置文件
	config := goutils.NewConfig().Load(path)
	return config.GlobalContent()
}

//初始化搜索引擎
func initsearchengine(settings map[string]string) *search.Engine {
	dict, _ := settings["DICT"]
	stop, _ := settings["STOP"]
	//初始化分词器
	seg := segmenter.InitChinaCut(dict)
	//生成一个搜索引擎
	searcher := search.NewSearchEngine()
	searcher.Init(search.EngineInitOptions{
		//分词器采用引擎自带的分词器
		Segmenter:            seg,
		StopTokenFile:        stop,
		UsePersistentStorage: false,
		IndexerInitOptions: &search.IndexerInitOptions{
			IndexType: search.LocationsIndex,
			BM25Parameters: &search.BM25Parameters{
				K1: 2.0,
				B:  0.75,
			},
		},
		//索引器接口实现，采用自带的wukong索引器
		CreateIndexer: func() search.SearchIndexer {
			return indexer.NewWuKongIndexer()
		},
		//排序器生成方法
		CreateRanker: func() search.SearchRanker {
			return ranker.NewWuKongRanker()
		},
	})
	return searcher
}

func main() {
	//配置文件信息
	settings := loadconf("conf/app.conf")
	if key, ok := utils.CheckSettings(settings); !ok {
		log.Fatal(fmt.Sprintf("not found %s in config file", key))
	}
	//启动系统任务
	go tasks.SysTask(settings)
	//初始化搜索引擎
	searcher := initsearchengine(settings)
	//生成一个索引任务对象
	//加载索引到内存
	go tasks.NewIndexTask(searcher).Index()
	port, _ := settings["PORT"]
	host, _ := settings["HOST"]
	web := utils.NewWeb(settings, searcher)
	log.Printf("server run on %s:%s", host, port)
	// Routes consist of a path and a handler function.
	r := mux.NewRouter()
	routers.Register(web, r)
	static, _ := settings["STATIC"]
	http.Handle("/"+static+"/", http.StripPrefix("/"+static+"/", http.FileServer(http.Dir("./"+static))))
	http.Handle("/", r)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
