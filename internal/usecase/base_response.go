package usecase

import (
	"github.com/iAbbos/go-http_server/internal/entity"
	"github.com/iAbbos/go-http_server/internal/entity/types"
	"net/http"
)

func BaseResponse() *entity.Response {
	resp := entity.NewResponse()
	resp.SetVersion(types.VERSION_1_1)
	resp.SetStatus(http.StatusOK, "OK")

	return resp
}
