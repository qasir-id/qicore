package main

import (
	"fmt"
	"log"
	"os"

	"github.com/qasir-id/qicore/scaffold"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0-rc"
	app.Usage = "Generate scaffold project layout for Qasir Team"
	app.Commands = []*cli.Command{
		{
			Name:    "qsr-service",
			Aliases: []string{"i"},
			Usage:   " qsr-service -n 'pos-service-account'",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("service name : ", c.String("name"))
				currDir, _ := os.Getwd()
				err := scaffold.New(false).Generate(scaffold.DataFlag{
					Path: currDir,
					Name: c.String("name"),
				})
				if err == nil {
					fmt.Println("Success Created. Please excute `make up` to start service.")
				}
				return err
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
