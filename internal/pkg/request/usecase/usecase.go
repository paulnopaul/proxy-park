package usecase

import (
	"net/http"
	"proxy-park/internal/pkg/domain"
	"proxy-park/tools/scanner"
	param_miner "proxy-park/tools/scanner/param-miner"
)

type usecase struct {
	requestRepo domain.RequestRepository
	scanner     scanner.Scanner
}

func NewRequestUsecase(requestRepo domain.RequestRepository) domain.RequestUsecase {
	return &usecase{
		requestRepo: requestRepo,
		scanner:     param_miner.NewParamMiner(),
	}
}

func (u usecase) SaveRequest(request http.Request) error {
	return u.requestRepo.SaveRequest(request)
}

func (u usecase) GetRequest(id string) (http.Request, error) {
	return u.requestRepo.GetRequest(id)
}

func (u usecase) GetAllRequests() ([]http.Request, error) {
	return u.requestRepo.GetAllRequests()
}

func (u usecase) RepeatRequest(id string) (http.Response, error) {
	req, err := u.GetRequest(id)
	if err != nil {
		return http.Response{}, err
	}

	client := http.Client{}
	response, err := client.Do(&req)
	if err != nil {
		return http.Response{}, err
	}
	return *response, nil
}

func (u usecase) Scan(id string) (string, error) {
	req, err := u.GetRequest(id)
	if err != nil {
		return "", err
	}
	return u.scanner.Scan(req)
}
