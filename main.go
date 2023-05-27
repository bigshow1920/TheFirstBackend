package main

import (
	"context"
	"log"
	"myapp/config"

	"myapp/app"

	_ "github.com/lib/pq"
)

func main() {
	configs, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	ctx := context.Background()
	err = app.NewApp(ctx, configs)

}
