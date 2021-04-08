package api

import (
	"blog/app/model"
	"blog/app/system/index/internal/define"
	"blog/app/system/index/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

func Index(r *ghttp.Request) {
	var (
		data *define.ContentDetailReq
	)
	if err := r.Parse(&data); err != nil {
		service.View.Render(r, model.View{
			Title:   "首页",
			MainTpl: "index.tpl",
			Data:    data,
		})
	}
}