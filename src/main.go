package main

import (
	"fmt"
	y "ginLearning/src/yframe"
	"net/http"
)

func main(){
	yframe := y.NewYFrameWork()

	yframe.GET("/yxdtest/:test/hello",yxdtestmao)
	yframe.GET("/yxdtest1/*test",yxdtestsnow)
	yframe.Run(":8888")
}

func yxdtestmao(c *y.Context){
	fmt.Printf("params: %v\n",c.Params)
	c.JSON(http.StatusOK, fmt.Sprintf("hello world,%v",c.Params["test"]))
}

func yxdtestsnow(c *y.Context){
	fmt.Printf("params: %v\n",c.Params)
	c.JSON(http.StatusOK, fmt.Sprintf("hello world,%v",c.Params["test"]))
}
