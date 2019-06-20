package echoinit

import (
	"echo_shop/config"
	"echo_shop/pkg/constants"
	pongo2utils "echo_shop/pkg/pongo2"
	"fmt"

	"echo_shop/pkg/flash"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
)

// SetupRender -
func SetupRender(e *echo.Echo) {
	pongo2utils.Setup(&pongo2utils.Config{
		Echo:        e,
		PublicPath:  config.String("APP.PUBLIC_DIR"),
		MixFilePath: config.String("APP.PUBLIC_DIR") + "/mix-manifest.json",
	})

	render := pongo2utils.NewRenderer()

	// template dir
	render.AddDirectory(config.String("APP.TEMPLATE_DIR"))

	// template global var
	globalVar := pongo2.Context{
		"APP_NAME":    config.String("APP.NAME"),
		"APP_RUNMODE": config.String("APP.RUNMODE"),
	}

	// render action
	render.UseContextProcessor(func(echoCtx echo.Context, pongoCtx pongo2.Context) {
		pongoCtx.Update(globalVar)

		other := pongo2.Context{}

		// url
		urlPath := echoCtx.Request().URL.Path
		other["route_path"] = urlPath

		// csrf
		csrf := echoCtx.Get(constants.CsrfContexntName)
		if c, ok := csrf.(string); ok && c != "" {
			other["csrf_token"] = c
			other["csrf_field"] = fmt.Sprintf(`<input type="hidden" name="%s" value="%s">`, constants.CsrfValueName, c)
			other["csrf_meta"] = fmt.Sprintf(`<meta name="%s" content="%s">`, constants.CsrfValueName, c)
		}

		// method 重写
		other["delete_method_field"] = fmt.Sprintf(`<input type="hidden" name="%s" value="DELETE">`, constants.MethodOverrideFromFormKeyName)
		other["put_method_field"] = fmt.Sprintf(`<input type="hidden" name="%s" value="PUT">`, constants.MethodOverrideFromFormKeyName)
		other["patch_method_field"] = fmt.Sprintf(`<input type="hidden" name="%s" value="PATCH">`, constants.MethodOverrideFromFormKeyName)

		// flash
		messageFlash := flash.NewMessageFlash(echoCtx).Read()
		other["messages"] = messageFlash
		oldvalueFlash := flash.NewOldValueFlash(echoCtx).Read()
		other["old_value"] = oldvalueFlash
		errorsFlash := flash.NewErrorsFlash(echoCtx).Read()
		other["errors"] = errorsFlash
		other["all_errors"] = flash.GetAllErrors(errorsFlash)

		pongoCtx.Update(other)
	})
	e.Renderer = render

	// register filters
	// pongo2.RegisterFilter("demo", demo)

	// register tags
	pongo2.RegisterTag("static", pongo2utils.StaticTag) // 获取静态文件地址
	pongo2.RegisterTag("mix", pongo2utils.MixTag)       // 配合 laravel-mix 使用
	pongo2.RegisterTag("route", pongo2utils.RouteTag)   // 通过 route name 生成 route path
}
