package cli

import "github.com/urfave/cli/v2"

func getGlobalFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "client-id",
			Usage:    "the client id",
			Required: true,
			EnvVars:  []string{"SBANKEN_CLIENT_ID"},
		},
		&cli.StringFlag{
			Name:     "client-secret",
			Usage:    "the client secret",
			Required: true,
			EnvVars:  []string{"SBANKEN_CLIENT_SECRET"},
		},
		&cli.StringFlag{
			Name:     "customer-id",
			Usage:    "customer id",
			Required: true,
			EnvVars:  []string{"SBANKEN_CUSTOMER_ID"},
		},
	}
}
