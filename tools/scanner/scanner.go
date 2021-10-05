package scanner

import (
	"net/http"
)

type Scanner interface {
	Scan(request http.Request) (string, error)
}
