/*
Author: Aosen
Date: 2016-02-01
Desc: 处理前端发过来的搜索请求
*/
package controllers

import (
	"net/http"
	"novel/utils"
)

type NovelSearchHandler struct {
	BaseHandler
}

func (self *NovelSearchHandler) Get(w http.ResponseWriter, r *http.Request, web *utils.Web) {
}
