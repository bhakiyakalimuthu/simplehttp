package app

import (
	"log"
	"net/http"
	"github.com/go-chi/chi"
)

var _ http.Handler = (*ServerOne)(nil)
type ServerOne struct {}


func (s *ServerOne) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`hello server one `))
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *ServerOne) RouteOne(){
	log.Fatal(http.ListenAndServe(":8080", s))
}


type ServerTwo struct{}

func (s *ServerTwo) RouteTwo(){
	http.HandleFunc("/hello",Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Hello(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`hello server two`))
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type ServerMux struct{}

func (s *ServerMux) RouteMux(){
	handler :=http.NewServeMux()
	handler.HandleFunc("/hello",Hello)
	log.Fatal(http.ListenAndServe(":8080",handler))
}

type ServerChi struct{}

func (s *ServerChi) RouteChi(){
	handler := chi.NewRouter()
}