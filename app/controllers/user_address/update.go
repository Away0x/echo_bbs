package useraddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/app/request"
)

// Update 编辑收货地址
func Update(c *context.AppContext, u *models.User) error {
	address := new(models.UserAddress)
	if err := c.ModelByID("user_address", &address); err != nil {
		return err.HTML()
	}

	req := new(request.AddressStoreForm)
	if err := c.BindAndValidate(req); err != nil {
		c.ErrorFlash(err)
		return c.RedirectByName("user_addresses.edit", address.ID)
	}

	if err := models.AssignAndUpdate(true, address, req); err != nil {
		c.ErrorFlash(c.WE(err, "收货地址更新失败"))
		return c.RedirectByName("user_addresses.edit", address.ID)
	}

	return c.RedirectByName("user_addresses.index")
}
