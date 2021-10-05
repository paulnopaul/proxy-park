package tools

import (
	"net/http"
	"proxy-park/internal/pkg/domain"
)

func FromHTTPRequest(request http.Request) domain.Request {
	return domain.Request{}
}

func FromHTTPResponse(response http.Response) domain.Response {
	return domain.Response{}
}

func FromHTTPFull(request http.Request, response http.Response) domain.Request {
	res := FromHTTPRequest(request)
	res.Response = FromHTTPResponse(response)
	return res
}
