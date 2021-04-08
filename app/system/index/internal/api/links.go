package api

import (
	"blog/app/model"
	"blog/app/system/index/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

func Links(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title:   "友链",
		MainTpl: "links.tpl",
	})
}
