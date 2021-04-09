package api

import (
	"blog/app/system/index/internal/define"
	"blog/app/system/index/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

func Index(r *ghttp.Request) {
	var (
		data *define.ContentPageReq
	)
	if err := r.Parse(&data); err != nil {
		// 获取文章列表
		service.View.Render(r, service.GetPostPage(data))
	} else {
		service.View.Render500(r)
	}
}
