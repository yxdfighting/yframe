package yframe

import (
	"net/http"
	"strings"
)

//todo:
// 1. 检测路由是否发生冲突   2.检测路由是否满足基本规则  eg: /:test/:test/hello

//将router抽象出来，以便做功能增强
//该部分抽象的是route的匹配和对应的handler
//同时对于ServeHTTP修改入参为Context

//关于router设计思路主要是每一个router实例对应一个tria树，该树可以承载router上所有的路由
//我们知道相同的路由，method不同的话可以对应不同的handler
//所以需要有一个map，map的key是由method和路由路径进行拼接，value就是对应的handler
//所以该map的路由路径其实就是从tria中得到，因为对于动态的路由匹配，我们需要将用户输入的path插入到tria中，
//同时在tria中进行计算得到用户输入的path在handler这个map中对应的value
type router struct{
	//对应一个tria
	roots *node
	//同样对应该method有一个执行方法
	handler map[string] HandleFunc
}

func NewRouter() *router{
	return &router{
		roots: &node{},
		handler: make(map[string]HandleFunc),
	}
}

func parsePattern(pattern string)[]string{
	pathSlice := strings.Split(pattern,"/")

	result := []string{}
	for _,path := range pathSlice{
		if path != ""{
			result = append(result,path)
		}
		//如果路由满足 /*test  类似时，后续的部分就不写入列表
		if len(path) > 0 && path[0] == '*'{
			break
		}
	}
	return result
}

//根据用户的path插入对应的tria
func (r *router) addRoute(method string,pattern string,handler HandleFunc){
	//将pattern插入生成一个root node
	root := &node{}
	//将pattern分割为一个slice，例如/yxd/test/abc --->  [yxd,test,abc]
	pathSlice := parsePattern(pattern)
	root.insert(pattern,pathSlice,0)
	r.roots = root

	key := method + "-" + pattern
	r.handler[key] = handler
}

//根据用户的输入path和method确定，
//当前用户输入能匹配到tria的哪一个node,进而返回该tria对用的pattern
//当遇到 * 或者 :时，需要将对应path中匹配的部分返回出来
//eg：  yxdtest/:hello/test    用户输入path yxdtest/myhello/test  c.GetParam["hello"] = myhello
//yxdtest/*hello/test
func (r *router) getRoute(path string)(*node,map[string]string){
	searchPath := parsePattern(path)
	params := make(map[string]string)

	//得到path对应的node
	node := r.roots.search(searchPath,0)
	if node != nil{
		//处理 '*' ':'的场景,填充param
		parts := parsePattern(node.pattern)
		for idx,part := range parts{
			if part[0] == ':'{
				//todo 增加校验，如果出现/yxdtest/:test/:test重复场景，返回nil
				params[part[1:]] = searchPath[idx]
			}
			if part[0] == '*' && len(part) > 1{
				//yxdtest/*hello/test   此时c.GetParam["hello"] = "hello/test"
				params[part[1:]] = strings.Join(searchPath[idx:],"/")
				//如果匹配到 *，就无须继续查找
				return node,params
			}
		}
	}
	return node,params
}


func (r *router) handlerRoute(c *Context){
	node,params := r.getRoute(c.Path)
	//每次处理到一个handler，如果存在params,就将结果写入Context
	//如果多个handler同时写入c.Param会不会报错？？？
	if node != nil{
		key := c.Method + "-" + node.pattern
		c.Params = params
		if handler,ok := r.handler[key];ok{
			handler(c)
		}else{
			c.String(http.StatusNotFound,"404 NOT FOUND: %s\n",key)
		}
	}else{
		c.String(http.StatusNotFound,"404 NOT FOUND\n")
	}
}
