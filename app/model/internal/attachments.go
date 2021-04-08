// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// Attachments is the golang structure for table attachments.
type Attachments struct {
	Id         int         `orm:"id,primary"  json:"id"`         //
	CreateTime *gtime.Time `orm:"create_time" json:"createTime"` //
	UpdateTime *gtime.Time `orm:"update_time" json:"updateTime"` //
	FileKey    string      `orm:"file_key"    json:"fileKey"`    //
	Height     int         `orm:"height"      json:"height"`     //
	MediaType  string      `orm:"media_type"  json:"mediaType"`  //
	Name       string      `orm:"name"        json:"name"`       //
	Path       string      `orm:"path"        json:"path"`       //
	Size       int64       `orm:"size"        json:"size"`       //
	Suffix     string      `orm:"suffix"      json:"suffix"`     //
	ThumbPath  string      `orm:"thumb_path"  json:"thumbPath"`  //
	Type       int         `orm:"type"        json:"type"`       //
	Width      int         `orm:"width"       json:"width"`      //
}
