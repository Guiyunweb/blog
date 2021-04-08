package boot

import (
	"blog/app/system/admin"
	"blog/app/system/index"
	_ "blog/packed"
)

func init() {
	admin.Init()
	index.Init()
}
