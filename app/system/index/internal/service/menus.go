package service

import (
	"blog/app/dao"
	"blog/app/model"
)

func getMenus() ([]*model.Menus, error) {
	return dao.Menus.Cache(0, "menu").FindAll()
}
