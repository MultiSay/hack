package main

import (
	"fmt"
	v1 "hack/api/v1"
	"hack/internal/app/config"
	"hack/internal/app/server"
	"hack/internal/app/store/sqlstore"
	"hack/internal/app/websocket"
	"hack/internal/app/worker"
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
	store, err := sqlstore.New(config)
	if err != nil {
		panic(err)
	}

	ws := websocket.NewWS()
	worker := worker.New(config, store, ws)
	api := v1.New(store, config, worker)
	srv := server.NewServer(store, config, api, ws)
	if err := srv.Start(fmt.Sprintf("%s:%d", config.Host, config.Port)); err != nil && err != http.ErrServerClosed {
		panic(err)
	}

}
