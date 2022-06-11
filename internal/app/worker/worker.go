package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"hack/internal/app/config"
	"hack/internal/app/model"
	"hack/internal/app/store"
	"hack/internal/app/websocket"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

type Worker struct {
	cfg      config.Config
	store    store.Store
	recordCh chan model.File
	ws       *websocket.WS
}

func New(cfg config.Config, s store.Store, ws *websocket.WS) *Worker {
	return &Worker{
		cfg:      cfg,
		store:    s,
		recordCh: make(chan model.File),
		ws:       ws,
	}
}

func (w *Worker) Init(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)

	for i := 0; i < w.cfg.NumOfWorkers; i++ {
		g.Go(func() error {
			var err error
			for ch := range w.recordCh {
				err = w.handle(ctx, ch)
			}
			return err
		})
	}
	return nil
}

func (w *Worker) handle(ctx context.Context, o model.File) error {

	// TODO Открыть файл и запустить скрипт обработки модели
	return nil
}

func (w *Worker) check(ctx context.Context, f model.File) error {
	log.Println("starting check")
	// TODO Открыть файл результата и записать в базу
	a := &model.PredictResult{}

	if a.Status == "SUCCESS" || a.Status == "INVALID" {
		f.Status = a.Status
		err := w.store.File().Update(ctx, f)
		if err != nil {
			return err
		}
		w.ws.Clients.Range(func(key, value interface{}) bool {
			result := model.PredictResult{
				Status:  a.Status,
				Message: a.Message,
				Data:    a.Data,
			}

			response, err := json.Marshal(result)
			if err != nil {
				log.Printf("key %s, error %s", key, err.Error())
				return false
			}

			value.(*websocket.Client).WriteMessage(string(response))
			return false
		})
	}
	return nil
}

func (w *Worker) Add(o model.File) {
	w.recordCh <- o
	log.Println("send Order to chan")
}

func (w *Worker) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			fileList, err := w.store.File().GetList(ctx)
			if err != nil && err != sql.ErrNoRows {
				log.Println(err)
				return err
			}
			for _, f := range fileList {
				log.Println("found new query row")
				go w.check(ctx, f)
			}
			log.Println("sleep")
			time.Sleep(3 * time.Second)
		}
	}
}
