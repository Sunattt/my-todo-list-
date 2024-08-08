package server

import (
	"context"
	"net/http"
	"time"
)
//с помощью него будем запускать сервер  
type Server struct {
	httpServer *http.Server
}

//запуск работы 
func (s *Server) Run(port string, handler http.Handler) error {
	//инкапсулируем значения 	 
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

//остановка работы 
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
