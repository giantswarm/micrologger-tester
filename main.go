package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name: "version",
			Action: func(c *cli.Context) error {
				fmt.Println("No compiled version")
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {

		logger, err := micrologger.New(micrologger.Config{})
		if err != nil {
			panic(microerror.Mask(err))
		}

		i := 0
		for range time.Tick(1 * time.Second) {
			logger.Log("level", "info", "message", fmt.Sprintf("Generating entry %v", i))
			i++
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
