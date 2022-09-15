package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "spell",
		Description: "A simple, fast, and secure web framework for Go.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "Enable verbose logging",
				Value:   false,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "build",
				Usage: "Build spell application",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "out",
						Aliases: []string{"o"},
						Usage:   "Output directory",
						Value:   "dist/public",
					},
					&cli.BoolFlag{
						Name:    "minify",
						Aliases: []string{"m"},
						Usage:   "Minify the output",
						Value:   true,
					},
				},
				Action: func(ctx *cli.Context) error {
					log.Println("Building spell application...")

					out := ctx.String("out")
					minify := ctx.Bool("minify")
					files := ctx.Args().Slice()

					return Build(out, minify, files...)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
