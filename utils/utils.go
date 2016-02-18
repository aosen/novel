/*
Author: Aosen
Data: 2016-2-2
Desc: novel server 工具箱
*/

package utils

import (
	"net/http"
	"sort"

	"github.com/aosen/goutils"
)

//检测settings文件
func CheckSettings(settings map[string]string) (string, bool) {
	for _, key := range []string{"DEBUG", "DBINFO", "HOST", "PORT", "DICT", "STOP", "INDEXSTOREPATH", "INDEXSTORENUM", "COLLECTIONPREFIX", "APPID", "APPSECRET", "STATIC", "NOVELPIC"} {
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

//根据novel结构体进行排序
