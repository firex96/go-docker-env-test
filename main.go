package main

import (
	"log"
	"net/http"
	"strconv"
	"github.com/go-chi/chi"
	"github.com/swaggo/http-swagger" // http-swagger middleware
	_ "./docs"
)

const port = "9999"

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	http.HandleFunc("/fib", fibHandler)

	log.Printf("listening on port %s", port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))

	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:" + port + "/swagger/doc.json"), //The url pointing to API definition"
	))

	http.ListenAndServe(":" + port, r)
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /fib
func fibHandler(w http.ResponseWriter, r *http.Request) {
	ns, ok := r.URL.Query()["num"]

	if !ok || len(ns) < 0 {
		log.Println("param 'num' is missing")

		return
	}

	n0, err := strconv.Atoi(ns[0])

	if err != nil {
		log.Printf("invalid number: %v", err)

		return
	}

	fb := strconv.Itoa(fib(n0))
	w.Write([]byte(fb + "\n"))
}

func fib(n int) int {
	a, b := 1, 1

	for i := 0; i < n-2; i++ {
		a, b = b, a+b
	}

	return b
}
