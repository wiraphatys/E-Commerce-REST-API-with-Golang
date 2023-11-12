package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/modules/middlewares/middlewaresHandlers"
	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/modules/middlewares/middlewaresRepositories"
	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/modules/middlewares/middlewaresUsecases"
	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/modules/monitor/monitorHandlers"
	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/modules/users/usersHandlers"
	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/modules/users/usersRepositories"
	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/modules/users/usersUsecases"
)

type IModuleFactory interface {
	MonitorModule()
	UsersModule()
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

func (m *moduleFactory) UsersModule() {
	repository := usersRepositories.UsersRepository(m.s.db)
	usecase := usersUsecases.UsersUsecase(m.s.cfg, repository)
	handler := usersHandlers.UsersHandler(m.s.cfg, usecase)

	router := m.r.Group("/users")

	router.Post("/signup", handler.SignUpCustomer)
	router.Post("/signin", handler.SignIn)
}
