package main

import (
	"fmt"
	"github.com/labstack/echo"
)
//import "github.com/labstack/echo/middleware"


func HandleBot() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header
			fmt.Printf("[HandleBot] header: %v\n", header)
			userAgentValue := header.Get("User-Agent")
			fmt.Printf("[HandleBot] user agent: %v \n", userAgentValue)
			return next(c)
			//if config.Skipper(c) {
			//	return next(c)
			//}

			//defer func() {
			//	if r := recover(); r != nil {
			//		err, ok := r.(error)
			//		if !ok {
			//			err = fmt.Errorf("%v", r)
			//		}
			//		stack := make([]byte, 10)
			//		//length := runtime.Stack(stack, !config.DisableStackAll)
			//		length := 10
			//		//if !config.DisablePrintStack {
			//		if false {
			//			c.Logger().Printf("[PANIC RECOVER] %v %s\n", err, stack[:length])
			//		}
			//		c.Error(err)
			//	}
			//}()
			//return next(c)
		}
	}
	//return HandleMainBot
}
