package main

import (
	"github.com/pkg/errors"

	"github.com/caddyserver/caddy"
	"github.com/caddyserver/caddy/caddyhttp/httpserver"

	// plug plugins
	_ "github.com/hacdias/caddy-minify"
	_ "github.com/caddyserver/caddy/caddyhttp/browse"
	_ "github.com/caddyserver/caddy/caddyhttp/root"
	_ "github.com/caddyserver/caddy/onevent"
)

func runCaddy() (*caddy.Instance, error) {
	caddy.AppName = "CaddyRenderStatic"
	caddy.AppVersion = "0.1.0"
	caddy.SetDefaultCaddyfileLoader("default", caddy.LoaderFunc(defaultLoader))

	caddyfile, err := caddy.LoadCaddyfile("http")
	if err != nil {
		return nil, errors.Wrap(err, "error loading caddyfile")
	}
	caddyserver, err := caddy.Start(caddyfile)
	if err != nil {
		return nil, errors.Wrap(err, "error starting caddyserver")
	}
	return caddyserver, nil
}

func defaultLoader(serverType string) (caddy.Input, error) {
	if serverType != "http" {
		return nil, errors.New("no http server!")
	}

	contents := httpserver.Host + ":" + httpserver.Port
	contents = contents + "\ntls off\n"
	contents = contents + "browse / " + Template + "\n"
	if Minify {
		contents = contents + "minify\n"
	}
	return caddy.CaddyfileInput{
		Contents:       []byte(contents),
		Filepath:       caddy.DefaultConfigFile,
		ServerTypeName: serverType,
	}, nil
}
