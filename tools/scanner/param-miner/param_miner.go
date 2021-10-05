package param_miner

import (
	"net/http"
	"proxy-park/tools/scanner"
)

type paramMiner struct {
	paramsList []string
}

func NewParamMiner() scanner.Scanner {
	return &paramMiner{}
}

func (m *paramMiner) Scan(request http.Request) (string, error) {
	return "", nil
}
