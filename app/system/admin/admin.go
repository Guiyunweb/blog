package admin

import (
	"github.com/gogf/gf/net/ghttp"
)

func Init(s *ghttp.Server) {
	s.AddStaticPath("/admin", "public/html")
}
