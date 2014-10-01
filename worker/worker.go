package worker

import (
	"log"

	"github.com/jrallison/go-workers"
	"github.com/kelseyhightower/envconfig"
	"github.com/nu7hatch/gouuid"
	"labix.org/v2/mgo"
)

var DB *mgo.Database

type Configuration struct {
	RedisServer   string
	RedisDatabase string
	RedisPool     string
	ProcID        string
}

func (config *Configuration) GetProcID() string {
	if config.ProcID == "" {
		u4, err := uuid.NewV4()
		if err != nil {
			panic(err)
		}
		config.ProcID = u4.String()
	}
	return config.ProcID
}

func Setup(_DB *mgo.Database) {
	var config Configuration

	err := envconfig.Process("worker", &config)
	if err != nil {
		log.Fatal("envconfig.Process: ", err.Error())
	}

	workers.Configure(map[string]string{
		// location of redis instance
		"server": config.RedisServer,
		// instance of the database
		"database": config.RedisDatabase,
		// number of connections to keep open with redis
		"pool": config.RedisPool,
		// unique process id for this instance of workers (for proper recovery of inprogress jobs on crash)
		"process": config.GetProcID(),
	})

	DB = _DB

	workers.Run()
}
