# caddy-static

Generate static-page file listings from [caddy's](https://github.com/mholt/caddy) [browse plugin](https://caddyserver.com/docs/browse).

## Usage

```sh
$ caddy-render -root /path/to/files -public /var/www

```

## Features

- Support for [minify](https://caddyserver.com/docs/http.minify) via `-minify`
