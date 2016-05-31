/*
小说文件下载，文件格式
Json格式, 格式如下:
{
    "title" : "XXXXXXXXXXX", //小说标题
    "chaptercontent": [
    {
        "chapterid": 123,
        "subtitle": "XXXXXXX",
        "content": "XXXXXXXXXXXXXX",
    },
    ]
}
*/

package controllers

import (
	"net/http"
	"novel/models"
	"novel/utils"
	"strconv"
)

type DownloadHandler struct {
	BaseHandler
}

func (self *DownloadHandler) Head(w http.ResponseWriter, r *http.Request, web *utils.Web) {
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
	downtext, err := models.NewNovelDownloadModel().GetNovelText(novelid)
	if err != nil {
		self.JsonResponse(w, "", 500)
	} else {
		self.JsonResponse(w, downtext, 200)
	}
}

func (self *DownloadHandler) Get(w http.ResponseWriter, r *http.Request, web *utils.Web) {
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
	downtext, err := models.NewNovelDownloadModel().GetNovelText(novelid)
	if err != nil {
		self.JsonResponse(w, "", 500)
	} else {
		self.JsonResponse(w, downtext, 200)
	}
}
