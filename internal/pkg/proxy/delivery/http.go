package delivery

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"proxy-park/internal/pkg/domain"
)

type proxyHandler struct {
	uCase domain.RequestUsecase
}

func NewProxyHandler(router *mux.Router, uCase domain.RequestUsecase) {
	newHandler := proxyHandler{
		uCase: uCase,
	}
	router.HandleFunc("*", newHandler.ServeHTTProxy).Methods("GET")
}

func (h *proxyHandler) ServeHTTProxy(w http.ResponseWriter, r *http.Request) {
	if r.Method == "CONNECT" {
		w.Write([]byte("haha https"))
	}

	r.RequestURI = ""
	r.Header.Del("Proxy-Connection")

	cl := http.Client{}
	resp, err := cl.Do(r)
	if err != nil {
		http.Error(w, "internal thing", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for key, valuesArr := range resp.Header {
		for _, value := range valuesArr {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
