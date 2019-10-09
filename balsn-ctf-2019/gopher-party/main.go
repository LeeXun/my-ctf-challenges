package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"balsnctf/gopherparty/handler"
	"balsnctf/gopherparty/store"

	"github.com/go-redis/redis"
)

func main() {
	// This server is running on t2.nano
	runtime.GOMAXPROCS(1)
	s := store.New()

	c := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	pong, err := c.Ping().Result()
	fmt.Println(pong, err)

	// We create a simple server using http.Server and run.
	server := &http.Server{
		Addr:    ":8000",
		Handler: handler.New(s, c),
	}
	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
