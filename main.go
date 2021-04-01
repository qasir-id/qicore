package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/qasir-id/qicore/scaffold"
	"github.com/urfave/cli/v2"
)

func main() {
	_ = godotenv.Load()
	app := cli.NewApp()
	app.Version = "1.0.0-rc"
	app.Usage = "Generate scaffold project layout for Qasir Team"
	app.Commands = []*cli.Command{
		{
			Name:    "service",
			Aliases: []string{"i"},
			Usage:   "service -n 'service name'",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("service name : ", c.String("name"))
				currDir, _ := os.Getwd()
				err := scaffold.New(false).Generate(scaffold.DataFlag{
					Path:          currDir,
					Name:          c.String("name"),
					SubStrService: "service",
				})
				if err == nil {
					fmt.Println("Success Created. Please execute `go run main.go` to start service.")
				}
				return err
			},
		}, {
			Name:    "gateway",
			Aliases: []string{"i"},
			Usage:   "gateway -n 'gateway name'",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("gateway name : ", c.String("name"))
				currDir, _ := os.Getwd()
				err := scaffold.New(false).Generate(scaffold.DataFlag{
					Path:          currDir,
					Name:          c.String("name"),
					SubStrService: "gateway",
				})
				if err == nil {
					fmt.Println("Success Created. Please execute `go run main.go` to start service.")
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
