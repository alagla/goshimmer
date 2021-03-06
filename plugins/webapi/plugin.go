package webapi

import (
	"context"
	"time"

	"github.com/iotaledger/hive.go/daemon"
	"github.com/iotaledger/hive.go/events"
	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/hive.go/node"
	"github.com/labstack/echo"
)

var PLUGIN = node.NewPlugin("WebAPI", node.Enabled, configure, run)
var log = logger.NewLogger("WebAPI")

var Server = echo.New()

func configure(plugin *node.Plugin) {
	Server.HideBanner = true
	Server.HidePort = true
	Server.GET("/", IndexRequest)

	daemon.Events.Shutdown.Attach(events.NewClosure(func() {
		log.Info("Stopping Web Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		if err := Server.Shutdown(ctx); err != nil {
			log.Errorf("Couldn't stop server cleanly: %s", err.Error())
		}
	}))
}

func run(plugin *node.Plugin) {
	log.Info("Starting Web Server ...")

	daemon.BackgroundWorker("WebAPI Server", func() {
		log.Info("Starting Web Server ... done")

		if err := Server.Start(":8080"); err != nil {
			log.Info("Stopping Web Server ... done")
		}
	})
}
