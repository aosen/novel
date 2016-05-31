/*
获取小说文本内容

2016-2-15

@aosen
*/

package controllers

import (
	"net/http"
	"novel/models"
	"novel/utils"
	"strconv"
)

type NovelContentHandler struct {
	BaseHandler
}

func (self *NovelContentHandler) Get(w http.ResponseWriter, r *http.Request, web *utils.Web) {
	var (
		chapterid int
		err       error
	)
	if cid := self.GetParam("chapterid"); cid != "" {
		if chapterid, err = strconv.Atoi(cid); err != nil {
			self.JsonResponse(w, "", 401)
			return
		}
	} else {
		self.JsonResponse(w, "", 401)
		return
	}
	//根据chapterid获取content
	nc := models.NewNovelContentModel()
	ret, e := nc.GetContent(chapterid)
	if e != nil {
		self.JsonResponse(w, "", 500)
	} else {
		self.JsonResponse(w, ret, 200)
	}
}
