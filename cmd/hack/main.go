package main

import (
	"context"
	"fmt"
	v1 "hack/api/v1"
	"hack/internal/app/config"
	"hack/internal/app/server"
	"hack/internal/app/store/sqlstore"
	"hack/internal/app/websocket"
	"hack/internal/app/worker"
	"net/http"

	"github.com/labstack/gommon/log"
)

// @title           hack API
// @version         1.0
// @description     API for Moscow City Hack 2022.
//
// @host      51.250.44.134
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
	log.Infof("[INIT] Init worker")
	ctx := context.Background()
	worker.Init(ctx)
	log.Infof("[INIT] Run worker")
	go worker.Run(ctx)
	log.Infof("[INIT] END")
	if err := srv.Start(fmt.Sprintf("%s:%d", config.Host, config.Port)); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
