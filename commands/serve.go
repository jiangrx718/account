package commands

import (
	"account/gopkg/graceful"
	"account/gopkg/utils"

	"github.com/urfave/cli/v2"

	"account/server/http/handlers"

	"net/http"
	_ "net/http/pprof"
)

func Serve(ctx *cli.Context) error {
	go func() {
		_ = http.ListenAndServe(":8999", nil)
	}()

	srv := utils.NewHttpServer(":8080")

	srv.RegisterHandler(handlers.NewHandler)

	graceful.Start(srv)

	graceful.Wait()
	
	return nil
}
