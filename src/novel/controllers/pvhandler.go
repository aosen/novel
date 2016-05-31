/*
修改pv数
put
/novelpv/
参数： novelid
返回字段:
novelid

2016-02-16
@aosen
*/

package controllers

import (
	"net/http"
	"novel/models"
	"novel/utils"
	"strconv"
)

type NovelPVHandler struct {
	BaseHandler
}

func (self *NovelPVHandler) Put(w http.ResponseWriter, r *http.Request, web *utils.Web) {
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
	pvobj := models.NewNovelPVModel()
	if ret, err := pvobj.PutPV(novelid); err != nil {
		self.JsonResponse(w, "", 401)
	} else {
		self.JsonResponse(w, ret, 200)
	}
}
