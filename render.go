package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/mholt/caddy/caddyhttp/httpserver"
	"github.com/pkg/errors"
)

type CaddyClient struct {
	addr string
	http.Client
}

func NewCaddyClient() *CaddyClient {
	return &CaddyClient{
		addr:   makeAddr(),
		Client: http.Client{},
	}
}

func makeAddr() string {
	scheme := "http"
	if httpserver.Port == httpserver.HTTPSPort {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s:%s", scheme, httpserver.Host, httpserver.Port)
}

func (c *CaddyClient) Get(uri string) (*http.Response, error) {
	u, err := url.Parse(c.addr + uri)
	if err != nil {
		return nil, err
	}
	return c.Client.Get(u.String())
}

var (
	cl      *CaddyClient
	rootDir string
)

func renderPublic() (err error) {
	cl = NewCaddyClient()
	dir, err := filepath.EvalSymlinks(httpserver.Root)
	if err != nil {
		return err
	}
	rootDir = dir
	return filepath.Walk(rootDir, render)
}

func render(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	uri := strings.TrimPrefix(path, rootDir)
	// we need to download the index.html files as well
	if info.IsDir() {
		uri = uri + "/"
	}

	resp, err := cl.Get(uri)
	if err != nil {
		return errors.Wrap(err, "error requesting file")
	}
	defer resp.Body.Close()

	uriDir, uriFile := filepath.Split(uri)
	publicDir := filepath.Join(Public, uriDir)

	// create dir in public/
	if err := os.MkdirAll(publicDir, os.ModeDir|0755); err != nil {
		return err
	}

	if info.IsDir() {
		uriFile = "index.html"
	}

	// open file in public/
	publicFile := filepath.Join(publicDir, uriFile)
	fd, err := os.OpenFile(publicFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	// write contents of response
	if _, err := io.Copy(fd, resp.Body); err != nil {
		return err
	}

	log.Println(uri, "::", resp.StatusCode)
	return nil
}
