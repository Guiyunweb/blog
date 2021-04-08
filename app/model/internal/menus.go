// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// Menus is the golang structure for table menus.
type Menus struct {
	Id         int         `orm:"id,primary"  json:"id"`         //
	CreateTime *gtime.Time `orm:"create_time" json:"createTime"` //
	UpdateTime *gtime.Time `orm:"update_time" json:"updateTime"` //
	Icon       string      `orm:"icon"        json:"icon"`       //
	Name       string      `orm:"name"        json:"name"`       //
	ParentId   int         `orm:"parent_id"   json:"parentId"`   //
	Priority   int         `orm:"priority"    json:"priority"`   //
	Target     string      `orm:"target"      json:"target"`     //
	Team       string      `orm:"team"        json:"team"`       //
	Url        string      `orm:"url"         json:"url"`        //
}
