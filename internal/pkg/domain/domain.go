package domain

import "net/http"

type Response struct {
	Headers map[string]string
	Cookie  string
	Method  string
	Code    int
	Body    string
}

type Request struct {
	Host     string
	Headers  map[string]string
	Cookie   string
	Method   string
	Path     string
	Body     string
	Response Response
}

type RequestRepository interface {
	SaveRequest(http.Request) error
	GetAllRequests() ([]http.Request, error)
	GetRequest(string) (http.Request, error)
}

type RequestUsecase interface {
	SaveRequest(http.Request) error
	GetRequest(string) (http.Request, error)
	GetAllRequests() ([]http.Request, error)
	RepeatRequest(string) (http.Response, error)
	Scan(string) (string, error)
}
