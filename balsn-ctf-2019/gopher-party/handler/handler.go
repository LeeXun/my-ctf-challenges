package handler

import (
	"log"
	"net/http"
	"time"

	"balsnctf/gopherparty/controller"
	"balsnctf/gopherparty/store"

	"github.com/go-redis/redis"
)

// New is a function
func New(s *store.Store, c *redis.Client) http.Handler {
	mux := http.NewServeMux()
	// Root
	mux.Handle("/", http.FileServer(http.Dir("static/")))

	// OauthGoogle
	mux.HandleFunc("/auth/google/login", controller.OauthGoogleLogin)
	mux.HandleFunc("/auth/google/callback", controller.OauthGoogleCallback)

	mux.HandleFunc("/register", controller.Register(s, c))

	return recoveryHandler(http.TimeoutHandler(mux, time.Second*3, "(O..O)!!!"))
}

func recoveryHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Panic Restarting...")
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, req)
	})
}
