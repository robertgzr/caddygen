package main // import "github.com/robertgzr/caddy-render"

import (
	"flag"
	"log"

	"github.com/oklog/run"
	"github.com/pkg/errors"
)

const (
	DefaultPublic = "./public"
)

var (
	Public = DefaultPublic
	Minify = true
)

func main() {
	flag.StringVar(&Public, "public", DefaultPublic, "Root path of the generated static-site")
	flag.BoolVar(&Minify, "minify", Minify, "Wether to minify the generated content")
	flag.Parse()

	var g run.Group
	{
		caddyserver, err := runCaddy()
		if err != nil {
			log.Fatal(err)
		}
		g.Add(func() (err error) {
			caddyserver.Wait()
			return nil
		}, func(err error) {
			if stopErr := caddyserver.Stop(); stopErr != nil {
				log.Fatal(errors.Wrap(stopErr, "error stopping caddy"))
			}
		})
	}
	{
		g.Add(func() error {
			return renderPublic()
		}, func(err error) {
			return
		})
	}

	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}
