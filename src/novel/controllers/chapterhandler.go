/*
获取小说的章节列表
get
/novelchapter/
参数 novelid
返回字段
title
subtitle
chapterid
novelid
chapter

2016-02-04
@aosen
*/

package controllers

import (
	"net/http"
	"novel/models"
	"novel/utils"
	"strconv"
)

type NovelChapterHandler struct {
	BaseHandler
}

func (self *NovelChapterHandler) Get(w http.ResponseWriter, r *http.Request, web *utils.Web) {
	var novelid int
	var err error
	if nid := self.GetParam("novelid"); nid != "" {
		if novelid, err = strconv.Atoi(nid); err != nil {
			self.JsonResponse(w, "", 401)
			return
		}
	} else {
		self.JsonResponse(w, "", 401)
		return
	}
	chapterobj := models.NewNovelChapterModel()
	if chapters, err := chapterobj.GetChapterList(novelid); err != nil {
		self.JsonResponse(w, "", 500)
		return
	} else {
		//根据chapter排序
		self.JsonResponse(w, chapters, 200)
		return
	}
}
