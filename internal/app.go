package internal

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/client/selector/static"
	micro_old_server "github.com/pr0head/micro-old-server/pkg"
	"time"
)

type Config struct {
	MicroSelector string
}

type Application struct {
	client micro_old_server.MicroOldServerService
}

func NewApplication() (app *Application) {
	cfg := &Config{}
	app = &Application{}

	var options []micro.Option

	if cfg.MicroSelector == "static" {
		options = append(options, micro.Selector(static.NewSelector()))
	}

	service := micro.NewService(options...)
	service.Init()

	app.client = micro_old_server.NewMicroOldServerService(micro_old_server.ServiceName, service.Client())

	return
}

func (app *Application) Run() {
	for {
		t := time.Now()
		req := &micro_old_server.TestRequest{Input: t.Format(time.RFC3339)}
		res, err := app.client.Test(context.Background(), req)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(res.Output)
		time.Sleep(1 * time.Second)
	}
	fmt.Println(1)
}
