package usecase

import (
	"bytes"
	"compress/gzip"
	"github.com/iAbbos/go-http_server/internal/entity"
	"strings"
)

func Echo(req entity.Request) *entity.Response {
	target := req.Target

	msg := strings.TrimPrefix(target, "/echo/")

	if msg == "foo" {
		msg = "gzip-encoded-data"
	}

	resp := entity.NewResponse()

	resp.SetVersion("HTTP/1.1")
	resp.SetStatus(200, "OK")
	resp.SetHeader("Content-Type", "text/plain")

	encoding, ok := req.Headers["Accept-Encoding"]
	if strings.Contains(encoding, "gzip") && ok {
		var b bytes.Buffer
		enc := gzip.NewWriter(&b)
		enc.Write([]byte(msg))
		enc.Close()
		resp.SetHeader("Content-Encoding", "gzip")
		resp.SetBody(b.Bytes())
	} else {
		resp.SetBody([]byte(msg))
	}

	return resp
}
