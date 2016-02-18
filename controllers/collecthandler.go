/*
修改收藏量
put
/novelcollect/
参数: novelid
返回字段:
novelid

2016-02-17
@aosen
*/

package controllers

import (
	"net/http"
	"novel/models"
	"strconv"

	"github.com/aosen/goutils"
)

type NovelCollectHandler struct {
	BaseHandler
}

func (self *NovelCollectHandler) Put(w http.ResponseWriter, r *http.Request, web *goutils.Web) {
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
	collectobj := models.NewNovelCollectModel()
	if ret, err := collectobj.PutCollect(novelid); err != nil {
		self.JsonResponse(w, "", 401)
	} else {
		self.JsonResponse(w, ret, 200)
	}
}
