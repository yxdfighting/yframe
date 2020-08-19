package yframe

import (
	"fmt"
	"net/http"
	"strings"
)

//定义一个方法
//方法入参出参与ServeHTTP相同，这样就可以承接对应的路由匹配方法
type HandleFunc func(c *Context)

type yFrameWork struct{
	*group
	router *router
}

func NewYFrameWork() *yFrameWork{
	return &yFrameWork{
		router: NewRouter(),
	}
}

func (y *yFrameWork) ServeHTTP(w http.ResponseWriter, r *http.Request){
	context := NewContext(w,r)
	y.router.handlerRoute(context)
}

func (y *yFrameWork) AddRoute(method string,pattern string,handler HandleFunc){
	key := method + "-" + pattern
	y.router.handler[key] = handler
	y.router.addRoute(method,pattern,handler)
}

func (y *yFrameWork) GET(pattern string,handler HandleFunc){
	y.router.handler["GET-"+pattern] = handler
	y.router.addRoute("GET",pattern,handler)
}

func (y *yFrameWork) POST(pattern string,handler HandleFunc){
	y.router.handler["POST-"+pattern] = handler
	y.router.addRoute("POST",pattern,handler)
}

func (y *yFrameWork) Run(host string){
	hostPort := strings.Split(host,":")
	if len(hostPort) != 2{
		fmt.Printf("host and port not illegal, %v",host)
		return
	}

	if hostPort[0] == ""{
		fmt.Printf("server is listening, localhost%v",host)
	}else{
		fmt.Printf("server is listening, %v",host)
	}

	if err := http.ListenAndServe(host,y);err != nil{
		fmt.Printf("ListenAndServe Error,%v",err)
		return
	}
}



