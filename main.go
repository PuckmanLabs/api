package main

import (
	"flag"

	"github.com/PuckmanLabs/api/server"
	"github.com/PuckmanLabs/api/worker"
)

func main() {
	session := server.NewSession("puckmanlabs_api")

	// Setup flags
	isWorker := flag.Bool("worker", false, "Run in worker mode")

	// Collect all flags
	flag.Parse()

	if !*isWorker {
		apiServer := server.NewServer(session)
		apiServer.Run()
	} else {
		worker.Setup(session.DB())
	}
}
