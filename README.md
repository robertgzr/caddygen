# caddygen

Generate static-page file listings from [caddy's](https://github.com/caddyserver/caddy) [browse plugin](https://caddyserver.com/docs/browse).

## Usage

```sh
$ caddygen -root /path/to/files -public /var/www
```

Build and run as a container

```
$ make IMAGE=<image> container
```

## Features

- Support for [minify](https://caddyserver.com/docs/http.minify) via `-minify`
