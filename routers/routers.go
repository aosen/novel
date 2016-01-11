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

	"github.com/aosen/utils"
	"github.com/gorilla/mux"
)

func Register(web *utils.Web, r *mux.Router) {
	// Bind to a port and pass our router in
	//r.HandleFunc("/cut/", g.Go(&controllers.CutHandler{}))
	r.NotFoundHandler = http.HandlerFunc(web.Go(controllers.NewNotFoundHandler()))
}
