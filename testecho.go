package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main()  {
	//创建一个ECHO服务实例
	e := echo.New()
	//ECHO里的URL路由以及对应的url处理函数
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK,"Hello,ECHO!")
	})
	//在1323号端口开启服务
	e.Logger.Fatal(e.Start(":1323"))
}


