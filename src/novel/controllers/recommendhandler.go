/*
推荐
*/

package controllers

import (
	"net/http"
	"novel/models"
	"novel/utils"
	"strconv"
)

type NovelRecommendHandler struct {
	BaseHandler
}

func (self *NovelRecommendHandler) Get(w http.ResponseWriter, r *http.Request, web *utils.Web) {
	tagid := self.GetParam("tagid")
	picpath := web.Settings["NOVELPIC"]
	recobj := models.NewNovelRecommendModel()
	if tagid == "" {
		reclist, err := recobj.GetList(picpath)
		if err != nil {
			self.JsonResponse(w, "", 500)
		} else {
			self.JsonResponse(w, reclist, 200)
		}
	} else {
		tag, err := strconv.Atoi(tagid)
		if err != nil {
			self.JsonResponse(w, "", 401)
		} else {
			reclist, err := recobj.GetMore(tag, picpath)
			if err != nil {
				self.JsonResponse(w, "", 500)
			} else {
				self.JsonResponse(w, reclist, 200)
			}
		}
	}
}
