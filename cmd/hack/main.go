package main

import (
	"fmt"
	"hack/internal/app/config"
	"hack/internal/app/server"
	"hack/internal/app/store"
	"net/http"
)

// @title           hack API
// @version         1.0
// @description     API for Moscow City Hack 2022.
//
// @host      localhost:8080
// @BasePath  /
func main() {
	config := config.Get()
	//подключение к бд
	store, err := store.New(config)
	if err != nil {
		panic(err)
	}

	srv := server.NewServer(store, config)
	if err := srv.Start(fmt.Sprintf("%s:%d", config.Host, config.Port)); err != nil && err != http.ErrServerClosed {
		panic(err)
	}

}
