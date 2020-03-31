package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"todo-api-gin-gorm/pkg/config"
	"todo-api-gin-gorm/pkg/models"
	"todo-api-gin-gorm/pkg/routers"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
)

var (
	defaultHostAddr = ":8080"
)

// Server provides the sub-command to start the API server.
func Server() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Start the todo service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "db-driver",
				Value:       "sqlite3",
				Usage:       "Database driver selection",
				EnvVars:     []string{"DATABASE_DRIVER"},
				Destination: &config.Database.Driver,
			},
			&cli.StringFlag{
				Name:        "db-name",
				Value:       "vote",
				Usage:       "Name for database connection",
				EnvVars:     []string{"DATABASE_NAME"},
				Destination: &config.Database.Name,
			},
			&cli.StringFlag{
				Name:        "db-username",
				Value:       "root",
				Usage:       "Username for database connection",
				EnvVars:     []string{"DATABASE_USERNAME"},
				Destination: &config.Database.Username,
			},
			&cli.StringFlag{
				Name:        "db-password",
				Value:       "root",
				Usage:       "Password for database connection",
				EnvVars:     []string{"DATABASE_PASSWORD"},
				Destination: &config.Database.Password,
			},
			&cli.StringFlag{
				Name:        "db-host",
				Value:       "localhost:3306",
				Usage:       "Host for database connection",
				EnvVars:     []string{"DATABASE_HOST"},
				Destination: &config.Database.Host,
			},
			&cli.StringFlag{
				Name:        "addr",
				Value:       defaultHostAddr,
				Usage:       "Address to bind the server",
				EnvVars:     []string{"SERVER_ADDR"},
				Destination: &config.Server.Addr,
			},
			&cli.StringFlag{
				Name:        "root",
				Value:       "/",
				Usage:       "Root folder of the app",
				EnvVars:     []string{"SERVER_ROOT"},
				Destination: &config.Server.Root,
			},
			&cli.BoolFlag{
				Name:        "pprof",
				Value:       false,
				Usage:       "Enable pprof debugging server",
				EnvVars:     []string{"SERVER_PPROF"},
				Destination: &config.Server.Pprof,
			},
			&cli.StringFlag{
				Name:        "sessionkey",
				Value:       "secret",
				Usage:       "Session Key",
				EnvVars:     []string{"Session_Key"},
				Destination: &config.Session.Key,
			},
		},
		Action: func(c *cli.Context) error {
			idleConnsClosed := make(chan struct{})

			// load global script
			log.Info().Msg("Initial module engine.")
			if err := models.NewEngine(); err != nil {
				log.Fatal().Err(err).Msg("Failed to initialize ORM engine.")
			}

			server := &http.Server{
				Addr:         config.Server.Addr,
				Handler:      routers.Load(),
				ReadTimeout:  5 * time.Second,
				WriteTimeout: 10 * time.Second,
			}

			go func(srv *http.Server) {
				sigint := make(chan os.Signal, 1)

				// interrupt signal sent from terminal
				signal.Notify(sigint, os.Interrupt)
				// sigterm signal sent from kubernetes
				signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
				defer signal.Stop(sigint)

				<-sigint

				log.Info().Msg("received an interrupt signal, shut down the server.")
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				defer cancel()
				// We received an interrupt signal, shut down.
				if err := srv.Shutdown(ctx); err != nil {
					// Error from closing listeners, or context timeout:
					log.Error().Err(err).Msg("HTTP server Shutdown")
				}
				close(idleConnsClosed)
			}(server)

			var (
				g errgroup.Group
			)

			g.Go(func() error {
				log.Info().Msgf("Starting shorten server on %s", config.Server.Addr)
				return server.ListenAndServeTLS("", "")
			})

			if err := g.Wait(); err != nil {
				log.Fatal().Err(err)
			}

			<-idleConnsClosed

			return nil
		},
	}
}
