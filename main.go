// main package ...
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Println(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
}

func run() error {
	e := echo.New()

	e.HTTPErrorHandler = withLogs(log.Printf, e.DefaultHTTPErrorHandler)

	e.GET("/ok", ok)
	e.GET("/not-found", notFound)
	e.GET("/not-http-err", notHTTPErr)

	if err := e.Start("127.0.0.1:3000"); err != nil {
		return fmt.Errorf("start echo http server: %w", err)
	}
	return nil
}

// This could also be an `errorfer` interface, I changed it to a function
// for simplicity (log.Printf can be provided upon call).
type errorLogger func(msg string, args ...interface{})

func withLogs(log errorLogger, base echo.HTTPErrorHandler) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		base(err, c)

		he, ok := err.(*echo.HTTPError)
		if !ok {
			he = &echo.HTTPError{
				Code:     http.StatusInternalServerError,
				Message:  http.StatusText(http.StatusInternalServerError),
				Internal: err,
			}
		}
		log("%s\n", he)
	}
}

func ok(c echo.Context) error {
	res := map[string]string{"it": "works!"}
	c.JSON(http.StatusOK, res)
	return nil
}

func notFound(c echo.Context) error {
	sqlErr := errors.New("some sql error")
	repoErr := errors.Wrap(sqlErr, "some repo error")
	svcErr := errors.Wrap(repoErr, "service error")

	return &echo.HTTPError{
		Code:     http.StatusNotFound,
		Message:  "something isn't found",
		Internal: svcErr,
	}
}

func notHTTPErr(c echo.Context) error {
	sqlErr := errors.New("some sql error")
	repoErr := errors.Wrap(sqlErr, "some repo error")
	svcErr := errors.Wrap(repoErr, "service error")
	return svcErr
}
