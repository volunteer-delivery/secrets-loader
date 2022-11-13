package main

import (
    "fmt"
    "log"
    "os"

    "github.com/urfave/cli/v2"
)

func main() {
    app := &cli.App{
        Name:  "secrets-loader",
        Usage: "Load secrets from AWS Parameters Store",
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name: "path",
                Aliases: []string{"p"},
                Usage: "Parameters Path to load from AWS",
            },
            &cli.StringFlag{
                Name: "region",
                Aliases: []string{"r"},
                Usage: "AWS region",
            },
            &cli.StringFlag{
                Name: "label",
                Aliases: []string{"l"},
                Usage: "Filter parametrs by label",
            },
        },
        Action: func(ctx *cli.Context) error {
            fmt.Println("boom! I say!")
            if len(ctx.String("path")) > 0 {
                fmt.Println(fmt.Sprintf("--path %v", ctx.String("path")))
            }
            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}