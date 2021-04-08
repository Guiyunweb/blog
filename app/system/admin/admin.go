package admin

import "github.com/gogf/gf/frame/g"

func Init() {
	s := g.Server()
	s.AddStaticPath("/admin", "public/html")
}
