/// 注册
package register

import (
	"echo_shop/app/context"
)

func Register(c *context.AppContext) error {
	return c.String(200, "register")
}
