package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	/**
	 * Init server
	 */
	server := echo.New()
	server.Logger.SetLevel(log.DEBUG)
	server.Pre(middleware.RemoveTrailingSlash())
	server.Use(middleware.Logger())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	/**
	 * Register frontend
	 */
	RegisterFrontend(server)

	/**
	 * Init group API and register API
	 */
	v1 := server.Group("/api/v1")
	RegisterPostAPI(v1)

	/**
	 *  Start and graceful shutdown server
	 */
	go func() {
		port := "8080"
		if v, ok := os.LookupEnv("PORT"); ok {
			port = v
		}
		if err := server.Start(fmt.Sprintf("0.0.0.0:%s", port)); err != nil {
			server.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}
}

/**
 *  Register frontend to Server
 */
func RegisterFrontend(e *echo.Echo) {
	frontend := rice.MustFindBox("./frontend/build")
	fe := http.FileServer(frontend.HTTPBox())

	e.GET("/static/*", echo.WrapHandler(fe))

	// IMPORTANT STEP
	e.GET("/*", func(c echo.Context) error {
		index, err := frontend.Open("index.html")
		if err != nil {
			return err
		}
		content, err := ioutil.ReadAll(index)
		if err != nil {
			return err
		}
		return c.HTMLBlob(http.StatusOK, content)
	})
}
