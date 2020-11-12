package app

import "online-tests/domain"

type App struct {
	TestsService domain.TestsService
}

func NewApp(es domain.TestsService) *App {
	return &App{
		TestsService: es,
	}
}
