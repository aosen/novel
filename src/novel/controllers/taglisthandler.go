/*
Author: Aosen
Date: 2016-02-01
Desc:
first / second *注：如果获取全部分类目录 无需传first second
返回字段：
firstid  secondid
firstname(一级分类名称)  secondname（二级分类名称）
*/

package controllers

import (
	"net/http"
	"novel/models"
	"novel/utils"
)

//请求获取分类列表
type TaglistHandler struct {
	BaseHandler
}

func (self *TaglistHandler) Get(w http.ResponseWriter, r *http.Request, web *utils.Web) {
	tagobj := models.NewTagListModel()
	if taglist, err := tagobj.GetTagList(); err != nil {
		self.JsonResponse(w, "", 500)
	} else {
		self.JsonResponse(w, taglist, 200)
	}
}
