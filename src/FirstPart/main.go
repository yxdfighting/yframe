package FirstPart

import (
	"fmt"
	"net/http"
)

/****************************************************ListenAndServe nil********************************************************************
//主要介绍go的原生http包如何处理一个请求
//http.ListenAndServe(":9527",nil);如果该方法第二个参数是nil，表示使用的默认Mux，
// 默认Mux由Handle和Handle方法增加handler
func main(){
	http.HandleFunc("/yxdtest/helloworld",HwHandler)
	fmt.Printf("Listen Port: %v\n",9527)
	//ListenAndServe starts an HTTP server with a given address and handler.
	// The handler is usually nil, which means to use DefaultServeMux.
	// Handle and HandleFunc add handlers to DefaultServeMux:
	if err := http.ListenAndServe(":9527",nil); err != nil{
		fmt.Printf("ListenAndServe error, %v",err)
		return
	}
}

func HwHandler(w http.ResponseWriter,r *http.Request){
	fmt.Printf("headerMsg: %v \n urlMsg: %v",r.Header,r.URL.Path)
	w.WriteHeader(http.StatusOK)
	if _,err := w.Write([]byte("yxdTest HelloWorld Successful")); err != nil{
		fmt.Printf("Write Msg error,%v",err)
	}
}

 ****************************************************************************************************/
//通过ListenAndServe 参数传入一个struct
//该struct实现了Handler interface，所以后续Mux会由该struct接管
//当前其实相当于一个简单的demo 实现了静态路由匹配
type YanDemo struct{
}

func NewYanDemo() *YanDemo {
	return &YanDemo{}
}

func (y *YanDemo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/yxdtest/helloworld":
		w.WriteHeader(http.StatusOK)
		if _,err := w.Write([]byte("yxdTest HelloWorld Successful")); err != nil{
			fmt.Printf("Write Msg error,%v",err)
		}
	case "/yxdtest/handsome":
		w.WriteHeader(http.StatusOK)
		if _,err := w.Write([]byte("yxdTest handsome forever")); err != nil{
			fmt.Printf("Write Msg error,%v",err)
		}
	}
}

func main(){
	yanInstance := NewYanDemo()
	fmt.Printf("Listen Port: %v\n",9527)
	if err := http.ListenAndServe(":9527",yanInstance); err != nil{
		fmt.Printf("ListenAndServe error, %v",err)
		return
	}
}