package cli

import (
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func getGlobalFlags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "client-id",
			Usage:   "the client id",
			Aliases: []string{"clid"},
			EnvVars: []string{"SBANKEN_CLIENT_ID"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "client-secret",
			Usage:   "the client secret",
			Aliases: []string{"s"},
			EnvVars: []string{"SBANKEN_CLIENT_SECRET"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "customer-id",
			Usage:   "customer id",
			Aliases: []string{"cuid"},
			EnvVars: []string{"SBANKEN_CUSTOMER_ID"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "style",
			Usage: "set output style",
		}),
		&cli.StringFlag{
			Name:    "config",
			Usage:   "path to YAML config",
			Aliases: []string{"c"},
			EnvVars: []string{"SBANKEN_CONFIG"},
		},
	}
}
