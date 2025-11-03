package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/maliarslan/fem-complete-go/internal/app"
	"github.com/maliarslan/fem-complete-go/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend port")
	flag.Parse()
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	// Defer ensures a function is called at the end of the program's execution,
	// after all other operations are complete,
	// which is useful for closing database connections
	defer app.DB.Close()

	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
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
