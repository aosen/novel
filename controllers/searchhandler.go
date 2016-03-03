/*
Author: Aosen
Date: 2016-02-01
Desc: 处理前端发过来的搜索请求
*/
package controllers

import (
	"net/http"
	"novel/models"
	"novel/utils"

	"github.com/aosen/search"
	"github.com/aosen/search/scorer"
)

type NovelSearchHandler struct {
	BaseHandler
}

func (self *NovelSearchHandler) Handle(w http.ResponseWriter, r *http.Request, web *utils.Web, text string) {
	searcher := web.Searcher
	searchresult := searcher.Search(search.SearchRequest{
		Text: text,
		RankOptions: &search.RankOptions{
			SearchScorer: scorer.NewBM25Scorer(),
		},
	})
	//获取docid列表
	var docids []int
	for _, doc := range searchresult.Docs {
		docids = append(docids, int(doc.DocId))
	}
	res, err := models.NewBaseModel().GetNovels(docids)
	if err != nil {
		self.JsonResponse(w, "", 200)
	} else {
		self.JsonResponse(w, res, 200)
	}
}

func (self *NovelSearchHandler) Get(w http.ResponseWriter, r *http.Request, web *utils.Web) {
	if text := self.GetParam("wd"); text == "" {
		self.JsonResponse(w, "", 401)
		return
	} else {
		self.Handle(w, r, web, text)
	}
}
