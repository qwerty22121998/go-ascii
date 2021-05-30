package controller

import "github.com/labstack/echo/v4"

type BaseController struct {
	binder echo.DefaultBinder
}

func DefaultBaseController() BaseController {
	return BaseController{
		binder: echo.DefaultBinder{},
	}
}
