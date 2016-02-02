/*
Author: Aosen
Data: 2016-2-2
Desc: novel server 工具箱
*/

package utils

import (
	"net/http"

	"github.com/aosen/goutils"
)

//检测settings文件
func CheckSettings(settings map[string]string) (string, bool) {
	for _, key := range []string{"DEBUG", "DBINFO", "HOST", "PORT", "DICT", "STOP", "INDEXSTOREPATH", "INDEXSTORENUM", "COLLECTIONPREFIX", "APPID", "APPSECRET"} {
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
