package handlers

import (
	"log/slog"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type HTTPHandler func(e echo.Context) error

func Make(h HTTPHandler) echo.HandlerFunc {
	return func(e echo.Context) error {
		if err := (h(e)); err != nil {
			slog.Error("Http Handler error", "err", err, "path", e.Request().URL.Path)
			return err
		}
		return nil
	}
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
