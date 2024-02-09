package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func main() {
	echoApp := echo.New()
	echoApp.GET("", HelloHandler)

	fmt.Println("Listening on :8000")
	echoApp.Logger.Fatal(echoApp.Start(":8000"))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func HelloHandler(c echo.Context) error {
	return Render(c, http.StatusOK, button("John"))
}
