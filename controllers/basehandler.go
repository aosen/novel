package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"novel/utils"

	"github.com/aosen/goutils"
)

/*返回的JSON数据*/
type Response struct {
	Values interface{} `json:"result"`
	Code   int         `json:"code"`
	Desc   string      `json:"desc"`
}

var CodeDic = map[int]string{
	200: "Success",
	401: "Invalid Argument",
	402: "Authentication Failed",
	500: "Unknown Error",
}

type BaseHandler struct {
	utils.WebHandler
	kwargs url.Values
}

func (self *BaseHandler) JsonResponse(w http.ResponseWriter, v interface{}, code int) {
	resp, _ := json.Marshal(&Response{
		Values: v,
		Code:   code,
		Desc:   CodeDic[code],
	})
	//w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(resp))
}

func (self *BaseHandler) checkSign(form url.Values, appid, appsecret string) bool {
	for _, key := range []string{"appid", "sign_method", "sign", "timestamp"} {
		//如果不存在以上的key，则验证失败
		if ok, err := goutils.Contains(key, form); !ok && err == nil {
			return false
		}
	}
	dict := map[string]string{}
	for k, v := range form {
		if len(v) != 1 {
			return false
		}
		if k == "appid" {
			if v[0] != appid {
				return false
			}
		}
		if k != "sign" {
			dict[k] = v[0]
		}
	}
	checkstr := appsecret + goutils.MapDictSortToStr(dict) + appsecret
	//生成sign 验证是否与发过来的sign是否一致
	sign := form["sign"][0]
	var mysign string
	switch dict["sign_method"] {
	case "md5":
		mysign = goutils.Md5(checkstr)
	}
	if sign != mysign {
		return false
	}
	return true
}

func (self *BaseHandler) GetParam(key string) (value string) {
	if len(self.kwargs[key]) > 0 {
		value = self.kwargs[key][0]
	}
	return
}

func (self *BaseHandler) Prepare(w http.ResponseWriter, r *http.Request, web *utils.Web) {
	r.ParseForm()
	self.kwargs = r.Form
	//生成参数字典
	appid, _ := web.Settings["APPID"]
	appsecret, _ := web.Settings["APPSECRET"]
	debug, _ := web.Settings["DEBUG"]
	if debug == "False" {
		if !self.checkSign(r.Form, appid, appsecret) {
			self.JsonResponse(w, "", 402)
			self.StopRun()
		}
	}
}
