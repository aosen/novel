/*
Author: Aosen
Data: 2016-2-2
Desc: novel server 工具箱
*/

package utils

import (
	"errors"
	"log"
	"net/http"
	"sort"

	"github.com/aosen/goutils"
	"github.com/aosen/search"
)

//检测settings文件
func CheckSettings(settings map[string]string) (string, bool) {
	for _, key := range []string{"DEBUG", "DBINFO", "HOST", "PORT", "DICT", "STOP", "INDEXSTORENUM", "APPID", "APPSECRET", "STATIC", "NOVELPIC"} {
		if ok, _ := goutils.Contains(key, settings); !ok {
			return key, false
		}
	}
	return "", true
}

//根据appid和appsecret检测请求的合法性
func CheckSign(r *http.Request) bool {
	r.ParseForm()
	return true
}

//根据字典中的某一个key来进行排序
type KVL []map[string]interface{}

func (self KVL) Len() int {
	return len(self)
}

func (self KVL) Less(i, j int) bool {
	return self[i]["chapter"].(int) < self[j]["chapter"].(int)
}

func (self KVL) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self KVL) Append(target KVL, kv map[string]interface{}) KVL {
	target = append(target, kv)
	return target
}

func MapDicSortToMap(dict KVL) {
	sort.Sort(dict)
}

func PutError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

type HandleInterface interface {
	Prepare(w http.ResponseWriter, r *http.Request, web *Web)
	Get(w http.ResponseWriter, r *http.Request, web *Web)
	Put(w http.ResponseWriter, r *http.Request, web *Web)
	Post(w http.ResponseWriter, r *http.Request, web *Web)
	Options(w http.ResponseWriter, r *http.Request, web *Web)
	Head(w http.ResponseWriter, r *http.Request, web *Web)
	Delete(w http.ResponseWriter, r *http.Request, web *Web)
	Connect(w http.ResponseWriter, r *http.Request, web *Web)
	Finish(w http.ResponseWriter, r *http.Request, web *Web)
	Closed() bool
}

/*全局控制对象*/
type Web struct {
	//配置信息
	Settings map[string]string
	Searcher *search.Engine
}

func NewWeb(setting map[string]string, searcher *search.Engine) *Web {
	return &Web{
		Settings: setting,
		Searcher: searcher,
	}
}

func (self *Web) Go(handler HandleInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//为了保证程序不会异常退出，增加recover
		debug, ok := GetSetting(self.Settings, "DEBUG")
		if !ok {
			PutError(errors.New("not found setting for debug"))
		}
		if debug != "True" {
			defer func() {
				if x := recover(); x != nil {
					log.Printf("[%v] caught panic: %v", r.RemoteAddr, x)
				}
			}()
		}
		log.Println(r.Method + " " + r.URL.String())
		//无论什么方法 都预先调用prepare方法
		handler.Prepare(w, r, self)
		//相应http方法关联处理
		if !handler.Closed() {
			switch r.Method {
			case "GET":
				handler.Get(w, r, self)
			case "PUT":
				handler.Put(w, r, self)
			case "POST":
				handler.Post(w, r, self)
			case "OPTIONS":
				handler.Options(w, r, self)
			case "HEAD":
				handler.Head(w, r, self)
			case "DELETE":
				handler.Delete(w, r, self)
			case "CONNECT":
				handler.Connect(w, r, self)
			}
		}
		//无论什么方法 结束后都调用finish方法
		if !handler.Closed() {
			handler.Finish(w, r, self)
		}
	}
}

//所有http处理类都继承此类
type WebHandler struct {
	closed bool
}

func (self *WebHandler) Prepare(w http.ResponseWriter, r *http.Request, web *Web) {
}

func (self *WebHandler) Get(w http.ResponseWriter, r *http.Request, web *Web) {
	w.WriteHeader(404)
	w.Write([]byte("404 not found"))
}

func (self *WebHandler) Put(w http.ResponseWriter, r *http.Request, web *Web) {
	w.WriteHeader(404)
	w.Write([]byte("404 not found"))
}

func (self *WebHandler) Post(w http.ResponseWriter, r *http.Request, web *Web) {
	w.WriteHeader(404)
	w.Write([]byte("404 not found"))
}

func (self *WebHandler) Options(w http.ResponseWriter, r *http.Request, web *Web) {
	w.WriteHeader(404)
	w.Write([]byte("404 not found"))
}

func (self *WebHandler) Head(w http.ResponseWriter, r *http.Request, web *Web) {
	w.WriteHeader(404)
	w.Write([]byte("404 not found"))
}

func (self *WebHandler) Delete(w http.ResponseWriter, r *http.Request, web *Web) {
	w.WriteHeader(404)
	w.Write([]byte("404 not found"))
}

func (self *WebHandler) Connect(w http.ResponseWriter, r *http.Request, web *Web) {
	w.WriteHeader(404)
	w.Write([]byte("404 not found"))
}

func (self *WebHandler) Finish(w http.ResponseWriter, r *http.Request, web *Web) {
}

func (self *WebHandler) StopRun() {
	self.closed = true
}

func (self *WebHandler) Closed() bool {
	return self.closed
}

//获取配置文件信息
func GetSetting(settings map[string]string, key string) (string, bool) {
	value, ok := settings[key]
	return value, ok
}
