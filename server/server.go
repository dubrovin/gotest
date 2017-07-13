package server

import (
	"github.com/dubrovin/gotest/appconf"
	"github.com/valyala/fasthttp"
	"log"
)

// https://github.com/throttled/throttled

// Server -
type Server struct {
	Conf    *appconf.Config
	API     *API
	ErrChan chan error
}

// NewServer -
func NewServer(conf *appconf.Config, api *API) (*Server, error) {
	return &Server{
		Conf:    conf,
		API:     api,
		ErrChan: make(chan error, 100),
	}, nil
}

// ListenAndServe -
func (s *Server) ListenAndServe() {
	log.Println("Listen and server addr = ", s.Conf.ListenAddress)
	s.ErrChan <- fasthttp.ListenAndServe(s.Conf.ListenAddress, s.API.router.HandleRequest)
}

// ReadErrChan -
func (s *Server) ReadErrChan() {
	for err := range s.ErrChan {
		log.Println("handlers server error: ", err)
	}
}

// Run -
func (s *Server) Run() {
	go s.ListenAndServe()
	go s.ReadErrChan()
}
