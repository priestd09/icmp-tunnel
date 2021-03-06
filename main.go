package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	Run(os.Args)
}

// Run command line tool
func Run(args []string) {

	app := cli.NewApp()

	app.Name = "icmp-tunnel"
	app.Usage = "Transparently tunnel your IP traffic through ICMP echo and reply packets"

	verboseFlag := cli.BoolFlag{
		Name:  "verbose, v",
		Usage: "Enable verbose logging",
	}

	app.Commands = []cli.Command{
		{
			Name:   "server",
			Usage:  "Runs server",
			Action: RunServer,
			Flags:  []cli.Flag{verboseFlag},
		},
		{
			Name:   "client",
			Usage:  "Runs client",
			Action: RunClient,
			Flags:  []cli.Flag{verboseFlag},
		},
	}

	app.Run(args)

}
