/*
Author: Aosen
Data: 2016-01-11
QQ: 316052486
Desc: 路由控制
*/
package routers

import (
	"net/http"
	"novel/controllers"
	"novel/utils"

	"github.com/gorilla/mux"
)

func Register(web *utils.Web, r *mux.Router) {
	// Bind to a port and pass our router in
	r.HandleFunc("/taglist/", web.Go(&controllers.TaglistHandler{}))
	r.HandleFunc("/novellist/", web.Go(&controllers.NovelListHandler{}))
	r.HandleFunc("/novelintroduction/", web.Go(&controllers.NovelIntroductionHandler{}))
	r.HandleFunc("/novelchapter/", web.Go(&controllers.NovelChapterHandler{}))
	r.HandleFunc("/novelcontent/", web.Go(&controllers.NovelContentHandler{}))
	r.HandleFunc("/novelpv/", web.Go(&controllers.NovelPVHandler{}))
	r.HandleFunc("/novelcollect/", web.Go(&controllers.NovelCollectHandler{}))
	r.HandleFunc("/novelsearch/", web.Go(&controllers.NovelSearchHandler{}))
	r.HandleFunc("/novelrank/", web.Go(&controllers.RankHandler{}))
	r.HandleFunc("/noveldownload/", web.Go(&controllers.DownloadHandler{}))
	r.HandleFunc("/novelrecommend/", web.Go(&controllers.NovelRecommendHandler{}))
	r.NotFoundHandler = http.HandlerFunc(web.Go(controllers.NewNotFoundHandler()))
}
