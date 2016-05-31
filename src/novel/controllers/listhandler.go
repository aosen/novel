/*
Author: Aosen
Date: 2016-02-02
Desc:
获取某分类下的小说列表
get
/novellist/
first second page limit
title novelid author picture intorduction
*/

package controllers

import (
	"net/http"
	"novel/models"
	"novel/utils"
	"strconv"
)

//根据一级分类id 二级分类id获取小说列表
type NovelListHandler struct {
	BaseHandler
}

func (self *NovelListHandler) Get(w http.ResponseWriter, r *http.Request, web *utils.Web) {
	var (
		firstid  int
		secondid int
		page     int
		limit    int
		err      error
	)
	if fid := self.GetParam("first"); fid != "" {
		if firstid, err = strconv.Atoi(fid); err != nil {
			self.JsonResponse(w, "", 401)
			return
		}
	} else {
		self.JsonResponse(w, "", 401)
		return
	}
	if sid := self.GetParam("second"); sid != "" {
		if secondid, err = strconv.Atoi(sid); err != nil {
			self.JsonResponse(w, "", 401)
			return
		}
	} else {
		self.JsonResponse(w, "", 401)
		return
	}
	if p := self.GetParam("page"); p != "" {
		if page, err = strconv.Atoi(p); err != nil {
			self.JsonResponse(w, "", 401)
			return
		}
	} else {
		self.JsonResponse(w, "", 401)
		return
	}
	if l := self.GetParam("limit"); l != "" {
		if limit, err = strconv.Atoi(l); err != nil {
			self.JsonResponse(w, "", 401)
			return
		}
	} else {
		self.JsonResponse(w, "", 401)
		return
	}
	novelpic, _ := web.Settings["NOVELPIC"]
	novelobj := models.NewNovelListModel()
	if novellist, err := novelobj.GetNovelList(
		firstid,
		secondid,
		page,
		limit,
		novelpic,
	); err != nil {
		self.JsonResponse(w, "", 500)
	} else {
		self.JsonResponse(w, novellist, 200)
	}
}
