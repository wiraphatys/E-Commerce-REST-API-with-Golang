package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/modules/middlewares/middlewaresHandlers"
	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/modules/middlewares/middlewaresRepositories"
	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/modules/middlewares/middlewaresUsecases"
	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/modules/monitor/monitorHandlers"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	r   fiber.Router
	s   *server
	mid middlewaresHandlers.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server, mid middlewaresHandlers.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresUsecase(repository)
	handler := middlewaresHandlers.MiddlewaresHandler(s.cfg, usecase)
	return handler
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.s.cfg)

	m.r.Get("/", handler.HealthCheck)
}
