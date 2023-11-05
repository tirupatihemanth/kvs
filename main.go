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
	kvMap.LoadFromFile(PERSIST_FILE_NAME)
	//go scheduleSaving()
}

func main() {
	router := chi.NewRouter()
	configureMiddleware(router)
	configureRoutes(router)

	file, err := os.OpenFile("kvsLog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	log.SetOutput(file)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT env variable not set. Create a .env file and put `PORT=\"65432\"` inside it")
	}

	log.Println("Starting Server at port", port)
	// 127.0.0.1 will not work on docker
	err = http.ListenAndServe("0.0.0.0:"+port, router)

	if err != nil {
		log.Fatalln("Error Starting the server at port %v\n", port)
	}
}

func configureRoutes(router *chi.Mux) {
	router.Get("/", middleware_key(getKeyHandler))
	router.Put("/", middleware_key(putKeyHandler))
	router.Delete("/", middleware_key(delKeyHandler))
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
