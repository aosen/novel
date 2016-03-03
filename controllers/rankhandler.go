/*
获取排行榜
*/

package controllers

import (
	"net/http"
	"novel/models"
	"novel/utils"
	"strconv"
)

type RankHandler struct {
	BaseHandler
}

func (self *RankHandler) Get(w http.ResponseWriter, r *http.Request, web *utils.Web) {
	page := self.GetParam("page")
	limit := self.GetParam("limit")
	if page == "" || limit == "" {
		self.JsonResponse(w, "", 401)
		return
	}
	pg, err := strconv.Atoi(page)
	if err != nil {
		self.JsonResponse(w, "", 401)
		return
	}
	lm, err := strconv.Atoi(limit)
	if err != nil {
		self.JsonResponse(w, "", 401)
		return
	}
	picpath := web.Settings["NOVELPIC"]
	ranks, err := models.NewNovelRankModel().GetRankList(pg, lm, picpath)
	if err != nil {
		self.JsonResponse(w, "", 500)
	} else {
		self.JsonResponse(w, ranks, 200)
	}
}
