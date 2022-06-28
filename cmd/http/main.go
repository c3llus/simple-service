package main

import (
	"fmt"
	"log"
	netHttp "net/http"

	"github.com/c3llus/cdk/app"
	server "github.com/c3llus/simple-service/server"
	"github.com/c3llus/simple-service/server/http"
)

func main() {
	httpApp, err := app.NewHTTP()
	if err != nil {
		log.Fatalf("NewHTTP")
	}

	server.Init(httpApp)

	httpServer := http.NewServer()
	if err := httpApp.RegisterServer(httpServer); err != nil {
		log.Fatal(fmt.Sprintf("[main][RegisterServer][err:%v]", err))
	}

	// run app
	if err := httpApp.Run(); err != nil && err != netHttp.ErrServerClosed {
		log.Fatal(fmt.Sprintf("[main][Run][err:%v]", err))
	}
}
