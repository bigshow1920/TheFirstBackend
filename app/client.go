package app

import (
	"context"
	"database/sql"
	"myapp/config"
	"myapp/handler"
	"myapp/service"
)

func NewClient(ctx context.Context, config config.Config) (*handler.PlayerHandler, error) {
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		return nil, err
	}
	playerService := service.NewPLayerService(conn)
	playerHandler := handler.NewPlayerHandler(playerService)
	return playerHandler, nil
}
