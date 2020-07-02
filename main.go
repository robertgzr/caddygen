package main // import "github.com/robertgzr/caddy-render"

import (
	"flag"
	"log"
	"os"

	"github.com/oklog/run"
)

const (
	DefaultPublic   = "./public"
	DefaultTemplate = ":default:"
)

var (
	Public   = DefaultPublic
	Minify   = true
	Template = DefaultTemplate
)

func main() {
	flag.StringVar(&Public, "public", DefaultPublic, "Root path of the generated static-site")
	flag.BoolVar(&Minify, "minify", Minify, "Wether to minify the generated content")
	flag.StringVar(&Template, "template", DefaultTemplate, "Path to a template file to pass to http.browse")
	flag.Parse()

	if Template[:1] == ":" && Template[len(Template)-1:] == ":" {
		tplPath, err := WriteTempTemplate(Template[1 : len(Template)-1])
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(tplPath)
		Template = tplPath
	}

	var g run.Group
	caddyserver, err := runCaddy()
	if err != nil {
		log.Fatalf("caddygen: %v", err)
	}
	g.Add(func() error {
		caddyserver.Wait()
		return nil
	}, func(error) { _ = caddyserver.Stop() })
	g.Add(func() error {
		return renderPublic()
	}, func(error) { log.Println("done") })

	if err := g.Run(); err != nil {
		log.Fatalf("caddygen: %v", err)
	}
}
