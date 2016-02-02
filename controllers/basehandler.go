package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/aosen/goutils"
)

/*返回的JSON数据*/
type Response struct {
	Values interface{} `json:"result"`
	Err    int         `json:"code"`
	Errmsg string      `json:"desc"`
}

var ERR = map[int]string{
	200: "Success",
	401: "Invalid argument",
	500: "Unknown Error",
}

type BaseHandler struct {
	goutils.WebHandler
}

func (self *BaseHandler) JsonResponse(w http.ResponseWriter, v interface{}, code int) {
	resp, _ := json.Marshal(&Response{
		Values: v,
		Err:    code,
		Errmsg: ERR[code],
	})
	//w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(resp))
}

func (self *BaseHandler) Prepare(w http.ResponseWriter, r *http.Request, web *goutils.Web) {
	r.ParseForm()
	var appid string
	if len(r.Form["appid"]) > 0 {
		appid = r.Form["appid"][0]
	}
	log.Println(appid)
}
