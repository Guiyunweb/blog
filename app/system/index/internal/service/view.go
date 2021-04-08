package service

import (
	"blog/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gmode"
)

// View 视图管理服务

var View = viewService{}

type viewService struct{}

func (s *viewService) Render(r *ghttp.Request, data ...model.View) {
	var (
		viewObj  = model.View{}
		viewData = make(g.Map)
	)
	if len(data) > 0 {
		viewObj = data[0]
	}
	if viewObj.Title == "" {
		viewObj.Title = g.Cfg().GetString(`setting.title`)
	}
	if viewObj.Keywords == "" {
		viewObj.Keywords = g.Cfg().GetString(`setting.keywords`)
	}
	if viewObj.Description == "" {
		viewObj.Description = g.Cfg().GetString(`setting.description`)
	}
	viewData = gconv.Map(viewObj)
	// 去除空数据
	for k, v := range viewData {
		if g.IsEmpty(v) {
			delete(viewData, k)
		}
	}

	r.Response.WriteTpl(g.Cfg().GetString("viewer.Layout"), viewData)
	// 开发模式下，在页面最下面打印所有的模板变量
	if gmode.IsDevelop() {
		r.Response.WriteTplContent(`{{dump .}}`, viewData)
	}
	// 退出当前业务函数执行
	r.Exit()
}
