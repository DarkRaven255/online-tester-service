package app

import "online-tester-service/domain"

type App struct {
	TestsService domain.TestsService
}

func NewApp(es domain.TestsService) *App {
	return &App{
		TestsService: es,
	}
}
