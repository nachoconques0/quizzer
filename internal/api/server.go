package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/cors"
)

type Server struct {
	*http.Server
	quizSvc Service
}

func NewServer(port string, quizSvc Service) (*Server, error) {
	if port == "" {
		return nil, errors.New("server port can not be empty")
	}
	s := &Server{
		quizSvc: quizSvc,
	}
	s.Server = &http.Server{
		Addr:         ":" + port,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		Handler:      LoggingMiddleware(s.router()),
	}
	return s, nil
}

func (s *Server) Run() error {
	c := cors.New(cors.Options{
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
	})
	s.Handler = c.Handler(s.Handler)

	go func() {
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error while starting the http server %v", err)
		}
	}()
	log.Printf("Server has started %s", s.Addr)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := s.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete.")
	return nil
}

func LoggingMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (s *Server) router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /question", s.quizSvc.GetQuestion)
	router.HandleFunc("POST /answer", s.quizSvc.SubmitAnswers)
	return router
}
