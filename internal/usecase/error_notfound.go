package usecase

import (
	"github.com/iAbbos/go-http_server/internal/entity"
)

func NotFoundError() *entity.Response {
	resp := entity.NewResponse()
	resp.SetVersion("HTTP/1.1")
	resp.SetStatus(404, "Not Found")

	return resp
}
