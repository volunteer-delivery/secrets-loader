package main

import (
	"fmt"
	"log"
	"os"
//	"secrets-loader/src/loader"
//	"secrets-loader/src/logger"

	. "github.com/urfave/cli/v2"
)

//func initLoaderConfig(ctx *Context) *loader.Config {
//	config := loader.NewConfig()
//	config.Path = ctx.String("path")
//	config.Region = ctx.String("region")
//	config.Label = ctx.String("label")
//	return config
//}
//
//func initLogger(ctx *Context) *logger.Logger {
//	level, _ := logger.ParseLogLevel(ctx.String("log-level"))
//	return logger.NewLogger(level)
//}

func main() {
	app := &App{
		Name: "secrets-loader",
		Commands: []*Command{
			{
				Name:  "get",
				Usage: "Load secrets from AWS Parameters Store",
				Flags: []Flag{
					&StringFlag{
						Name:     "path",
						Aliases:  []string{"p"},
						Usage:    "Parameters Path to load from AWS",
						Required: true,
					},
					&StringFlag{
						Name:    "label",
						Aliases: []string{"l"},
						Usage:   "Filter parametrs by label",
					},
                    &StringFlag{
                        Name:    "region",
                        Aliases: []string{"r"},
                        Usage:   "AWS region",
                    },
				},
                Action: func(ctx *Context) error {
//                    loggerInstance := initLogger(ctx)
//                    loaderInstance := loader.NewLoader(loggerInstance, initLoaderConfig(ctx))
//
//                    for _, secret := range loaderInstance.Load() {
//                        loggerInstance.Data("%v=%v", secret.Name, secret.Value)
//                    }

                    fmt.Println("Hello from cli")

                    return nil
                },
			},
		},
		Flags: []Flag{
			&StringFlag{
				Name:  "log-level",
				Value: "none",
//				Action: func(ctx *Context, level string) error {
//					_, ok := logger.ParseLogLevel(ctx.String("log-level"))
//					if !ok {
//						return fmt.Errorf("invalid log level %v", level)
//					}
//					return nil
//				},
//				Usage: fmt.Sprintf("Set logging `level`. Available levels: %v", logger.LogLevelNames),
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
