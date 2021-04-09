package boot

import (
	_ "blog/packed"
	"blog/router"
	"github.com/gogf/gcache-adapter/adapter"
	"github.com/gogf/gf/frame/g"
)

func init() {

	router.Init()

	// 将数据库缓存设置为Redis
	adapter := adapter.NewRedis(g.Redis())
	g.DB().GetCache().SetAdapter(adapter)
}
