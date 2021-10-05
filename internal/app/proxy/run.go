package proxy

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"proxy-park/internal/pkg/proxy/delivery"
	"proxy-park/internal/pkg/request/repository/mongo"
	"proxy-park/internal/pkg/request/usecase"
	"proxy-park/tools/mongo_helper"
)

func Run(host string, port int) error {
	collection, err := mongo_helper.GetMongoCollection("mongo://localhost", "proxy", "requests")
	if err != nil {
		return err
	}
	mongoRepo := mongo.NewMongoRepo(collection)
	requestUsecase := usecase.NewRequestUsecase(mongoRepo)

	r := mux.NewRouter()

	delivery.NewProxyHandler(r, requestUsecase)

	server := http.Server{
		Addr:    host + ":" + fmt.Sprint(port),
		Handler: r,
	}
	return server.ListenAndServe()
}
