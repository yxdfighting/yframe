package yframe

//定义路由分组
//参考gin使用方式
//router.Group("/api",MyMiddelware())
//v := router.Group("/v1")
//v.GET("/",func(c *gin.Context))
//对于group中的url，可以运行相同的中间件，可以统一进行管理

//所以我们可以将yframe抽象为一个最顶层的group
//每一个group需要操作router


type group struct{
	groupPath  string
	//分组嵌套的话需要有parent支持
	parent *group
	//所有group共享一个router实例
	router *router
}


func (g *group) Group(prefix string) *group{
	newGroup := &group{
		groupPath: g.groupPath + prefix,
		parent: g,
		router: g.router,
	}
	return newGroup
}

//每当通过group添加路由时，相当于在tria树上增加对应的记录，，而path则是由当前group的parent的path+pre
func (g *group) addRoute(method string,pre string,handler HandleFunc){
	pattern := g.groupPath + pre
	g.router.addRoute(method,pattern,handler)
}

func (g *group) GET(pattern string,handler HandleFunc){
	g.addRoute("GET",pattern,handler)
}

func (g *group) POST(pattern string,handler HandleFunc){
	g.addRoute("POST",pattern,handler)
}
