package index

import (
	"blog/app/system/index/internal/api"
	"github.com/gogf/gf/frame/g"
)

func Init() {
	s := g.Server()
	s.BindHandler("/", api.Index)
	s.BindHandler("/page/:page", api.Index)
	s.BindHandler("/archives", api.Archives)
	s.BindHandler("/links", api.Links)
	s.BindHandler("/about", api.About)
	s.BindHandler("/post/:id", api.Post)
}
