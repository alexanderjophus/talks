package testing

import (
	"net/http"
)

// A Server that says hello
type Server struct {
	router *http.ServeMux
}

// NewServer is created
func NewServer() Server {
	s := Server{router: http.NewServeMux()}
	s.router.HandleFunc("/", helloWorld)
	return s
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
