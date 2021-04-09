package service

import (
	"blog/app/dao"
	"blog/app/model"
	"blog/app/system/index/internal/define"
	"github.com/gogf/gf/frame/g"
)

func GetPostPage(data *define.ContentPageReq) model.View {

	total, _ := dao.Posts.Cache(0).Count(g.Map{
		"type": 0,
	})

	list, _ := dao.Posts.Cache(0).Page(data.Page, 10).FindAll(g.Map{
		"type": 0,
	})

	menus, _ := getMenus()

	return model.View{
		Page: model.Page{
			Page:  data.Page,
			Total: total,
		},
		Menu: menus,
		Data: list,
	}

}
