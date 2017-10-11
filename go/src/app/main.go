package main

import (
	"app/controller"

	"flag"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	Engine *echo.Echo
}

func New() *Server {
	return &Server{Engine: echo.New()}
}

func (s *Server) init() {
}

func (s *Server) middleware() {
	s.Engine.Use(middleware.Logger())
	s.Engine.Use(middleware.Recover())
	s.Engine.Use(middleware.RemoveTrailingSlash())
}

func (s *Server) route() {
	sample := controller.Sample{}
	foo := controller.Foo{}
	sample.SetupSample(s.Engine)
	foo.SetupFoo(s.Engine)
}

func (s *Server) run(addr string) {
	s.Engine.Logger.Fatal(s.Engine.Start(addr))
}

func main() {
	var (
		addr = flag.String("addr", ":8080", "addr to bind")
	)
	flag.Parse()
	s := New()
	s.init()
	s.middleware()
	s.route()
	s.run(*addr)
}
