package main

import (
	"context"
	"flag"
	"net"
	"net/http"

	"github.com/starudream/go-lib/app"
	"github.com/starudream/go-lib/log"
	"github.com/starudream/go-lib/router"

	. "github.com/starudream/github-asset-mirror/config"
	"github.com/starudream/github-asset-mirror/route"
)

var server *http.Server

func init() {
	flag.StringVar(&C.Addr, "addr", "0.0.0.0:80", "http server address")
	flag.StringVar(&C.Storage, "storage", "/storage", "file storage path")
	flag.StringVar(&C.Owner, "owner", "", "github owner")
	flag.StringVar(&C.Repo, "repo", "", "github repo")
	flag.StringVar(&C.Name, "name", "{{.repo}}-{{.os}}-{{.arch}}-{{.platform}}-{{.ver}}.zip", "github repo")
	flag.StringVar(&C.Proxy, "proxy", "", "available proxy: ghproxy fastgit")

	flag.Parse()

	if C.Owner == "" || C.Repo == "" {
		log.Fatal().Msg("owner and repo must be set")
	}

	route.Register()

	server = &http.Server{Addr: C.Addr, Handler: router.Handler()}
}

func main() {
	app.Add(serve)
	app.Defer(shutdown)
	err := app.Go()
	if err != nil {
		log.Fatal().Msgf("app init fail: %v", err)
	}
}

func serve(context.Context) error {
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return err
	}
	log.Info().Msgf("http server start at %s", server.Addr)
	return server.Serve(ln)
}

func shutdown() {
	_ = server.Shutdown(context.Background())
	log.Info().Msgf("http server shutdown")
}
