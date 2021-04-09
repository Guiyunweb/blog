package index

import (
	"blog/app/system/index/internal/api"
	"blog/app/system/index/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

func Init(s *ghttp.Server) {

	// 前台系统自定义错误页面
	s.BindStatusHandler(401, func(r *ghttp.Request) {
		service.View.Render401(r)
	})
	s.BindStatusHandler(403, func(r *ghttp.Request) {
		service.View.Render403(r)
	})
	s.BindStatusHandler(404, func(r *ghttp.Request) {
		service.View.Render404(r)
	})
	s.BindStatusHandler(500, func(r *ghttp.Request) {
		service.View.Render500(r)
	})

	s.BindHandler("/", api.Index)
	s.BindHandler("/page/:page", api.Index)
	s.BindHandler("/archives", api.Archives)
	s.BindHandler("/links", api.Links)
	s.BindHandler("/about", api.About)
	s.BindHandler("/post/:id", api.Post)
}
