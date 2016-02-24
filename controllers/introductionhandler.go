/*
参数：novelid 小说id
返回字段:
title
novelid
author
picture
introduction
chapternum 章节数量

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

//获取小说简介
type NovelIntroductionHandler struct {
	BaseHandler
}

func (self *NovelIntroductionHandler) Get(w http.ResponseWriter, r *http.Request, web *utils.Web) {
	var (
		novelid int
		err     error
	)
	if nid := self.GetParam("novelid"); nid != "" {
		if novelid, err = strconv.Atoi(nid); err != nil {
			self.JsonResponse(w, "", 401)
			return
		}
	} else {
		self.JsonResponse(w, "", 401)
		return
	}

	novelpic, _ := web.Settings["NOVELPIC"]
	introobj := models.NewNovelIntroductionModel()
	if intro, err := introobj.GetNovelIntroduction(novelid, novelpic); err != nil {
		self.JsonResponse(w, "", 500)
	} else {
		self.JsonResponse(w, intro, 200)
	}
}
