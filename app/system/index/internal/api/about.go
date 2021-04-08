package api

import (
	"blog/app/model"
	"blog/app/system/index/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

func About(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title:   "关于",
		MainTpl: "archives.tpl",
	})
}
