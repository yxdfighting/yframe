package yframe

//中间件主要是完成对应的处理
//处理的输入是Context对象，在获取到用户req之后完成Context的初始化，然后通过对应的中间件进行处理Context
//对于多个中间件，按照顺序进行执行，中间件的func作用于group
//对于中间件，可以区分在handler前执行，还是在handler后执行
//todo
type middleware struct{
	middlewareType string
	middlewareFunc []HandleFunc
}

