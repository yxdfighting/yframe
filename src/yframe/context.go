package yframe

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//将所有上下文都封装在Context中
type Context struct{
	Writer http.ResponseWriter
	Req *http.Request

	Path string
	Params map[string]string
	Method string
	StatusCode int
}

//w http.ResponseWriter r *http.Request
//将这两个封装后，加入到Context中，
//同时将req 和 writer衍生的常用属性放在Context中
func NewContext(w http.ResponseWriter,r *http.Request) *Context{
	return &Context{
		w,r,r.URL.Path,nil,r.Method,0,
	}
}

func (c *Context) Query(key string) string{
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Param(key string) string{
	return c.Params[key]
}

func (c *Context) PostForm(key string) string{
	return c.Req.Form.Get(key)
}

func (c *Context) SetStatusCode(code int){
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) String(statusCode int,format string,values ...interface{}){
	c.SetHeader("Content-Type","text/plain")
	c.Writer.WriteHeader(statusCode)

	if _, err := c.Writer.Write([]byte(fmt.Sprintf(format,values...)));err != nil{
		http.Error(c.Writer,err.Error(),500)
	}
}

func (c *Context) SetHeader(key string,value string){
	c.Writer.Header().Set(key,value)
}

func (c *Context) JSON(statusCode int,obj interface{}){
	c.SetHeader("Content-Type","application/json")
	c.Writer.WriteHeader(statusCode)

	var res []byte
	var err error

	if res,err = json.Marshal(obj); err != nil{
		http.Error(c.Writer,err.Error(),500)
	}

	if _, err := c.Writer.Write(res);err != nil{
		http.Error(c.Writer,err.Error(),500)
	}
}