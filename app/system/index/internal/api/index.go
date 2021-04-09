package api

import (
	"blog/app/dao"
	"blog/app/model"
	"blog/app/system/index/internal/define"
	"blog/app/system/index/internal/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Index(r *ghttp.Request) {
	var (
		data *define.ContentPageReq
	)
	if err := r.Parse(&data); err != nil {
		list, _ := dao.Posts.Page(data.Page, 10).FindAll(g.Map{
			"type": 0,
		})
		total, _ := dao.Posts.Count(g.Map{
			"type": 0,
		})
		service.View.Render(r, model.View{
			Title:   "首页",
			MainTpl: "index.tpl",
			Data:    list,
			Page: model.Page{
				Page:  data.Page,
				Size:  10,
				Total: total,
			},
		})
	} else {
		service.View.Render500(r)
	}
}
