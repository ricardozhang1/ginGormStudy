package common

import (
	"net/http"
)

// 声明一个新的数据类型(函数类型)
type FilterHandle func(response http.ResponseWriter, request *http.Request) error

// Filter 拦截器结构体
type Filter struct {
	// 用来存储需要拦截的URL
	filterMap map[string]FilterHandle
}

// Filter 初始化函数
func NewFilter() *Filter {
	return &Filter{filterMap: make(map[string]FilterHandle)}
}

// 注册拦截器
func (f *Filter) RegisterFilterUri(uri string, handler FilterHandle) {
	f.filterMap[uri] = handler
}

// 根据Uri获取对应的handle
func (f *Filter) GetFilterHandle(uri string) FilterHandle {
	return f.filterMap[uri]
}


// 声明新的函数类型
type WebHandle func(response http.ResponseWriter, request *http.Request)

// 执行拦截器
func (f *Filter) Handle(webHandle WebHandle) func(response http.ResponseWriter, request *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		for path, handler := range f.filterMap {
			// request.RequestURI
			if path == request.URL.Path {
				// 执行拦截逻辑
				if err := handler(response, request); err != nil {
					_, _ =response.Write([]byte(err.Error()))
					return
				}
				// 跳出循环
				break
			}
		}
		// 执行正常注册的函数
		webHandle(response, request)
	}
}





