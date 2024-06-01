package usecase

import (
	"github.com/iAbbos/go-http_server/internal/entity"
	"net/http"
)

func InternalError() *entity.Response {
	resp := entity.NewResponse()
	resp.SetVersion("HTTP/1.1")
	resp.SetStatus(http.StatusInternalServerError, "Internal Server Error")

	return resp
}
