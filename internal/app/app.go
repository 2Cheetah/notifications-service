package app

import (
	"github.com/2Cheetah/notifications-service/internal/repository"
	"github.com/2Cheetah/notifications-service/internal/server"
)

type App struct {
	Server     server.Server
	Service    server.Service
	NovuClient repository.NovuClient
}
