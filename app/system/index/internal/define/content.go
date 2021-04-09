package define

type ContentDetailReq struct {
	Id   uint `v:"min:1#请选择查看的内容"`
	Page uint `v:"min:1#请选择查看的内容"`
}

type ContentPageReq struct {
	Page int `v:"min:1#请选择选择分页"`
}
