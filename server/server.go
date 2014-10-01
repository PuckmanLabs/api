package server

import (
	"log"

	"github.com/go-martini/martini"
	"github.com/kelseyhightower/envconfig"
	"github.com/martini-contrib/render"
)

type Configuration struct {
}

/*
Wrap the Martini server struct.
*/
type Server *martini.ClassicMartini

func NewServer(session *DatabaseSession) Server {
	var config Configuration

	err := envconfig.Process("api", &config)
	if err != nil {
		log.Fatal("envconfig.Process: ", err.Error())
	}

	// Create the server and set up middleware.
	m := Server(martini.Classic())
	m.Use(render.Renderer(render.Options{
		IndentJSON: true,
	}))
	m.Use(session.Database())

	m.NotFound(func(r render.Render) {
		r.JSON(404, map[string]string{
			"error": "Not found",
		})
	})

	return m
}
