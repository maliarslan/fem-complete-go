package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/maliarslan/fem-complete-go/internal/app"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend port")
	flag.Parse()
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/health", HealthCheck)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("Running the application on %d\n", port)

	err = server.ListenAndServe()

	if err != nil {
		app.Logger.Fatal(err)
	}

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}
