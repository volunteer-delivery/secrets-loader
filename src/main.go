package main

import (
    "fmt"
    "log"
    "os"
    "secrets-loader/src/loader"

    . "github.com/urfave/cli/v2"
)

func initLoaderConfig(ctx *Context) *loader.Config {
    config := loader.NewConfig()
    config.Path = ctx.String("path")
    config.Region = ctx.String("region")
    config.Label = ctx.String("label")
    return config
}

func main() {
    app := &App{
        Name:  "secrets-loader",
        Usage: "Load secrets from AWS Parameters Store",
        Flags: []Flag{
            &StringFlag{
                Name: "path",
                Aliases: []string{"p"},
                Usage: "Parameters Path to load from AWS",
                Required: true,
            },
            &StringFlag{
                Name: "region",
                Aliases: []string{"r"},
                Usage: "AWS region",
            },
            &StringFlag{
                Name: "label",
                Aliases: []string{"l"},
                Usage: "Filter parametrs by label",
            },
        },
        Action: func(ctx *Context) error {
            instance := loader.NewLoader(initLoaderConfig(ctx))
            secrets := instance.Load()

            for _, secret := range secrets {
                fmt.Println(fmt.Sprintf("%v=%v", secret.Name, secret.Value))
            }

            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}