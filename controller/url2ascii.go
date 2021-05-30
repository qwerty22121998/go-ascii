package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/qwerty22121998/go-ascii/dto"
	"github.com/qwerty22121998/go-ascii/service"
	"net/http"
)

type URLToAsciiController struct {
	BaseController

	convertService service.ConvertService
}

func NewURLToAsciiController(provider *service.Provider) *URLToAsciiController {
	return &URLToAsciiController{
		BaseController: DefaultBaseController(),
		convertService: provider.ConvertService,
	}
}

func (a *URLToAsciiController) parse(c echo.Context) (*dto.AsciiFromUrlRequest, error) {
	var req dto.AsciiFromUrlRequest

	if err := a.binder.BindQueryParams(c, &req); err != nil {
		return nil, err
	}
	return &req, nil
}

func (a *URLToAsciiController) ToImage(c echo.Context) error {
	req, err := a.parse(c)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, dto.AsciiFromUrlErrorResponse{
			AsciiFromUrlRequest: *req,
			Message:             err.Error(),
		})
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, dto.AsciiFromUrlErrorResponse{
			AsciiFromUrlRequest: *req,
			Message:             err.Error(),
		})
	}
	// process
	res, err := a.convertService.FromUrlToImage(req.Url, req.Size)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, dto.AsciiFromUrlErrorResponse{
			AsciiFromUrlRequest: *req,
			Message:             err.Error(),
		})
	}

	return c.Blob(http.StatusOK, "image/png", res)
}

func (a *URLToAsciiController) ToText(c echo.Context) error {

	req, err := a.parse(c)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, dto.AsciiFromUrlErrorResponse{
			AsciiFromUrlRequest: *req,
			Message:             err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, dto.AsciiFromUrlErrorResponse{
			AsciiFromUrlRequest: *req,
			Message:             err.Error(),
		})
	}

	// process
	res, err := a.convertService.FromUrlToString(req.Url, req.Size)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, dto.AsciiFromUrlErrorResponse{
			AsciiFromUrlRequest: *req,
			Message:             err.Error(),
		})
	}

	return c.String(http.StatusOK, res)
}
