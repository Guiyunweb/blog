package router

import (
	"blog/app/system/admin"
	"blog/app/system/index"
	"github.com/gogf/gf/frame/g"
)

func Init() {

	s := g.Server()
	admin.Init(s)
	index.Init(s)

}
