package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger *log.Logger
	server *http.Server
}

func New(logger *log.Logger) *Server {
	// Создаём роутер и регистрируем хендлеры
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	// Настраиваем HTTP-сервер
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Возвращаем экземпляр сервера
	return &Server{
		logger: logger,
		server: httpServer,
	}
}

// Запускает сервер
func (s *Server) Start() error {
	s.logger.Printf("server is running on the port %s", s.server.Addr)
	return s.server.ListenAndServe()
}

// Останавливает сервер
func (s *Server) Stop() error {
	s.logger.Println("stopping the server...")
	return s.server.Shutdown(nil)
}
