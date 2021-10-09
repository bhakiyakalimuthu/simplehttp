package app

import (
	"github.com/go-chi/chi"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var _ http.Handler = (*ServerOne)(nil)
// ServerOne using http handleFunc
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


// ServerTwo using http handleFunc
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

// ServerThree using http handleFunc
type ServerThree struct{}

func (s *ServerThree) RouteThree(){
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter,_ *http.Request){
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`hello server three `))
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})))
}

// ServerMux using http mux
type ServerMux struct{}

func (s *ServerMux) RouteMux(){
	handler :=http.NewServeMux()
	handler.HandleFunc("/hello",HelloMux)
	log.Fatal(http.ListenAndServe(":8080",handler))
}

func HelloMux(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`hello server mux`))
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// ServerChi using chi router
type ServerChi struct{}

func (s *ServerChi) RouteChi(){
	handler := chi.NewRouter()
	handler.HandleFunc("/hello",HelloChi)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func HelloChi(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	_,err := w.Write([]byte(`Hello server chi`))
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}


// ServerJS Julienschmidt httprouter
type ServerJS struct{}

func(s *ServerJS)RouteJS(){
	handler := httprouter.New()
	handler.GET("/hello",HelloJS)
	log.Fatal(http.ListenAndServe(":8080",handler))
}

func HelloJS(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	w.WriteHeader(http.StatusOK)
	_,err := w.Write([]byte(`Hello server js`))
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
	}
}