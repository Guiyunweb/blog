package model

// View 视图渲染内容对象
type View struct {
	Title       string      // 页面标题
	Keywords    string      // 页面Keywords
	Description string      // 页面Description
	Error       string      // 错误信息
	MainTpl     string      // 自定义MainTpl展示模板文件
	Redirect    string      // 引导页面跳转
	ContentType string      // 内容模型
	Page        Page        // 分页信息
	Data        interface{} // 页面参数
}

type Page struct {
	Page  int
	Size  int
	Total int
}
