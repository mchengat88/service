package handlers

import (
	"github.com/mchengat/service/business/mid"
	"github.com/mchengat/service/foundation/web"
	"log"
	"net/http"
	"os"
)

// API constructs a http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log*log.Logger) *web.App{
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics(log))

	check := check{
		log: log,
	}
	app.Handle(http.MethodGet, "/readiness", check.readiness)
	return app
}
