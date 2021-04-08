package api

import (
	"blog/app/model"
	"blog/app/system/index/internal/define"
	"blog/app/system/index/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

func Post(r *ghttp.Request) {
	var (
		data *define.ContentDetailReq
	)
	if err := r.Parse(&data); err != nil {
		service.View.Render(r, model.View{
			Title:   "友链",
			MainTpl: "links.tpl",
			Data:    data,
		})
	}

}
