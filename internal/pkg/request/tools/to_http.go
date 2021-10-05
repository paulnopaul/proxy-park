package tools

import (
	"net/http"
	"net/url"
	"proxy-park/internal/pkg/domain"
)

func ToHTTPRequest(request domain.Request) http.Request {
	res := http.Request{}
	for key, value := range request.Headers {
		res.Header.Add(key, value)
	}
	res.Method = request.Method
	res.URL, _ = url.Parse(request.Host + request.Path)
	return res
}
