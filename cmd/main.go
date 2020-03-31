package main

import (
	"os"
	"time"

	"todo-api-gin-gorm/pkg/config"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func globalFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug",
			Value:       true,
			Usage:       "Activate debug information",
			EnvVars:     []string{"SERVER_DEBUG"},
			Destination: &config.Server.Debug,
		},
		&cli.BoolFlag{
			Name:        "log-color",
			Value:       true,
			Usage:       "enable colored logging",
			EnvVars:     []string{"LOGS_COLOR"},
			Destination: &config.Logs.Color,
		},
		&cli.BoolFlag{
			Name:        "log-pretty",
			Value:       true,
			Usage:       "enable pretty logging",
			EnvVars:     []string{"LOGS_PRETTY"},
			Destination: &config.Logs.Pretty,
		},
		&cli.StringFlag{
			Name:        "log-level",
			Value:       "info",
			Usage:       "set logging level",
			EnvVars:     []string{"LOGS_LEVEL"},
			Destination: &config.Logs.Level,
		},
	}
}

func globalCommands() []*cli.Command {
	return []*cli.Command{
		Server(),
	}
}

func globalBefore() cli.BeforeFunc {
	return func(c *cli.Context) error {
		setupLogger()
		return nil
	}
}

func main() {
	app := &cli.App{
		Name:     "todo server",
		Usage:    "todo system service",
		Compiled: time.Now(),
		Flags:    globalFlags(),
		Commands: globalCommands(),
		Before:   globalBefore(),
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show the help, so what you see now",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("can't run app")
	}
}
