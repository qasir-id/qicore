package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	qodes "gitlab.com/qasir/web/project/qasircore.git/transport/grpc/codes"
	qhttp "gitlab.com/qasir/web/project/qasircore.git/transport/http"
)

type GetPing struct{}

func (h *GetPing) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	resp := &qhttp.Response{
		Code:    qodes.Success,
		Message: qodes.StatusMessage[qodes.Success],
		Data: map[string]interface{}{
			"ping": "pong",
		},
		TraceID: qhttp.GetTraceID(ctx),
	}

	return resp.JSON(c)
}

func NewGetPing() *GetPing {
	return &GetPing{}
}
