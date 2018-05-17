# Session middleware for [Echo](https://github.com/labstack/echo)

[![Build][Build-Status-Image]][Build-Status-Url] [![Coverage][Coverage-Image]][Coverage-Url] [![ReportCard][reportcard-image]][reportcard-url] [![GoDoc][godoc-image]][godoc-url] [![License][license-image]][license-url]

## Quick Start

### Download and install

```bash
$ go get -u -v github.com/go-session/echo-session
```

### Create file `server.go`

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/go-session/echo-session"
	"github.com/labstack/echo"
	"gopkg.in/session.v2"
)

func main() {
	e := echo.New()

	e.Use(echosession.New(
		session.SetCookieName("session_id"),
		session.SetSign([]byte("sign")),
	))

	e.GET("/", func(ctx echo.Context) error {
		store := echosession.FromContext(ctx)
		store.Set("foo", "bar")
		err := store.Save()
		if err != nil {
			return err
		}
		return ctx.Redirect(302, "/foo")
	})

	e.GET("/foo", func(ctx echo.Context) error {
		store := echosession.FromContext(ctx)
		foo, ok := store.Get("foo")
		if !ok {
			return ctx.String(http.StatusNotFound, "not found")
		}
		return ctx.String(http.StatusOK, fmt.Sprintf("foo:%s", foo))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
```

### Build and run

```bash
$ go build server.go
$ ./server
```

### Open in your web browser

<http://localhost:8080>

    foo:bar


## MIT License

    Copyright (c) 2018 Lyric

[Build-Status-Url]: https://travis-ci.org/go-session/echo-session
[Build-Status-Image]: https://travis-ci.org/go-session/echo-session.svg?branch=master
[Coverage-Url]: https://coveralls.io/github/go-session/echo-session?branch=master
[Coverage-Image]: https://coveralls.io/repos/github/go-session/echo-session/badge.svg?branch=master
[reportcard-url]: https://goreportcard.com/report/github.com/go-session/echo-session
[reportcard-image]: https://goreportcard.com/badge/github.com/go-session/echo-session
[godoc-url]: https://godoc.org/github.com/go-session/echo-session
[godoc-image]: https://godoc.org/github.com/go-session/echo-session?status.svg
[license-url]: http://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/npm/l/express.svg