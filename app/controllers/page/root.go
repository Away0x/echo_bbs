package page

import (
	"echo_shop/pkg/errno"
	"errors"

	"echo_shop/pkg/validate"

	"github.com/labstack/echo/v4"
)

type UserForm struct {
	validate.BaseForm
	// Name  string `validate:"required|minLen:7"`
	Email string `validate:"required|minLength:7"`
	// Age      int       `validate:"required|int|min:1|max:99"`
	// CreateAt int       `validate:"min:1"`
	// Safe     int       `validate:"-"`
	// UpdateAt time.Time `validate:"required"`
	// Code     string    `validate:"customValidator"` // 使用自定义验证器
}

// Root -
func Root(c echo.Context) error {
	req := new(UserForm)

	if err := c.Bind(&req); err != nil {
		return errno.ParamsErr.SetErrors(err)
	}
	if err := c.Validate(req); err != nil {
		return errno.ParamsErr.SetErrors(err)
	}

	return errors.New("sss")
}
