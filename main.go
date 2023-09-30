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

	file, err := os.OpenFile("kvsLog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatalln(err)
    }

    log.SetOutput(file)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT env variable not set. Create a .env file and put `PORT=\"80\"` inside it")
	}

	log.Println("Starting Server at port", port)
	err = http.ListenAndServe("127.0.0.1:"+port, router)

	if err != nil {
		log.Fatalln("Error Starting the server at port %v\n", port)
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
