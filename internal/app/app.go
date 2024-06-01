package app

import (
	tspsrv "github.com/iAbbos/go-http_server/internal/delivery/tcp/server"
	configpkg "github.com/iAbbos/go-http_server/internal/pkg/config"
)

type App struct {
	Config *configpkg.Config
}

func NewApp(config *configpkg.Config) (*App, error) {
	return &App{
		Config: config,
	}, nil
}

func (a *App) Run() error {
	server := tspsrv.NewServer(a.Config)
	err := server.Run()
	if err != nil {
		return err
	}
	return nil
}
