package server

import (
	"log"
	"net/http"
	"time"
)

// Logging - функция логгирования запросов к серверу,
// она оборачивает мультиплексор, к которому привязаны хендлеры,
// внутри засекаем время начала выполнения запроса
// и логируем вызываемые методы и URI
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("Method: %s  URI: %s  time: %s", req.Method, req.RequestURI, time.Since(start))
	})
}
