package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

var kvMap *KVMap

func init() {
	godotenv.Load(".env")
	kvMap = &KVMap{make(map[string]string), sync.RWMutex{}}
}

func main() {
	router := chi.NewRouter()
	configureMiddleware(router)

	v1Router := chi.NewRouter()
	configureRoutes(v1Router)
	router.Mount("/v1", v1Router)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT env variable not set. Create a .env file and put `PORT=\"80\"` inside it")
	}

	log.Printf("Starting Server at port %v\n", port)
	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		log.Fatalf("Error Starting the server at port %v\n", port)
	}
}

func configureRoutes(v1Router *chi.Mux) {
	v1Router.Get("/kvs", middleware_key(getKeyHandler))
	v1Router.Put("/kvs", middleware_key(putKeyHandler))
	v1Router.Delete("/kvs", middleware_key(delKeyHandler))
}

func configureMiddleware(router *chi.Mux) {

	// CORS Middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "PUT", "DELETE"}, // Not allowing "POST", "OPTIONS"
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}
