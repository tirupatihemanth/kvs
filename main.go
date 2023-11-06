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

	// Loads environment variables such as port number from the .env file
	// If running inside docker container docker will setup PORT env variable as per Dockerfile configuration.
	godotenv.Load(".env")
	kvMap = &KVMap{make(map[string]string), sync.RWMutex{}}

	// Disabling persistence for the coding assingment 2 to ensure our throughput, latency plots are reliable.
	// kvMap.LoadFromFile(PERSIST_FILE_NAME)
	// go scheduleSaving()
}

func main() {
	router := chi.NewRouter()
	configureMiddleware(router)
	configureRoutes(router)

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatalln("PORT env variable not set. Create a .env file and put `PORT=\"8080\"` inside it")
	}
	file, err := os.OpenFile("kvsLog"+"_"+port+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	log.SetOutput(file)

	log.Println("Starting Server at port", port)
	// 127.0.0.1 will not work on docker
	err = http.ListenAndServe("0.0.0.0:"+port, router)

	if err != nil {
		log.Fatalln("Error Starting the server at port %v\n", port)
	}
}

// Route PUT, GET, DELETE requests to appropriate handlers.
func configureRoutes(router *chi.Mux) {
	router.Get("/", middleware_key(getKeyHandler))
	router.Put("/", middleware_key(putKeyHandler))
	router.Delete("/", middleware_key(delKeyHandler))
}

// Allowing only GET, PUT, DELETE. Not allowing "POST", "OPTIONS".
func configureMiddleware(router *chi.Mux) {

	// CORS Middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}
