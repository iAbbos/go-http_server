package usecase

import (
	"github.com/iAbbos/go-http_server/internal/entity"
	"github.com/iAbbos/go-http_server/internal/entity/types"
	"net/http"
)

func BadRequestError() *entity.Response {

	resp := entity.NewResponse()
	resp.SetVersion(types.VERSION_1_1)
	resp.SetStatus(http.StatusBadRequest, "Bad Request")

	return resp
}
