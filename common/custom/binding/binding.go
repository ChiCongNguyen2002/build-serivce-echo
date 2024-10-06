package binding

import (
	validate2 "BuildService/common/custom/validate"
	"github.com/labstack/echo/v4"
)

var (
	validate     = validate2.NewValidate()
	customBinder = &CustomBinder{}
)

type CustomBinder struct{}

func (cb *CustomBinder) Bind(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}

	if err := validate.ValidateStruct(i); err != nil {
		return err
	}

	return nil
}

func GetBinding() *CustomBinder {
	return customBinder
}
