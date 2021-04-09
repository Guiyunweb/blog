package service

import (
	"blog/app/dao"
	"blog/app/model"
	"github.com/gogf/gf/frame/g"
)

func init() {
	cacheMenu()
}

// cacheMenu 缓存菜单信息
func cacheMenu() {
	g.Log().Print("开始缓存菜单信息")
	menu, _ := dao.Menus.FindAll()
	_, err := g.Redis().Do("SET", "menus", menu)
	if err != nil {
		panic(err)
	}
}

func getMenus() []*model.Menus {
	result, err := g.Redis().DoVar("GET", "menus")
	if err != nil {
		panic(err)
	}

	var menus []*model.Menus
	if err = result.Struct(&menus); err != nil {
		panic(err)
	}
	if menus == nil {
		cacheMenu()
		return getMenus()
	}
	return menus
}
