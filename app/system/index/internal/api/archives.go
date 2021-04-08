package api

import (
	"blog/app/model"
	"blog/app/system/index/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

func Archives(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title:   "归档",
		MainTpl: "archives.tpl",
	})
}
