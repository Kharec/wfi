package main

import (
	"errors"
	"github.com/Kharec/waitforit/internal/wfi"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "wfi",
		HelpName:  "wait for it",
		Usage:     "try to join a host on a port until connection succeeded",
		UsageText: "wfi <protocol> <host> <port>",
		ArgsUsage: "[proto host port]",
		Action: func(cCtx *cli.Context) error {
			if cCtx.Args().Len() < 3 {
				return errors.New("Not enough arguments")
			}
			WaitForIt(cCtx.Args())
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func WaitForIt(args cli.Args) bool {
	for {
		if wfi.TryPort(args.Get(0), args.Get(1), args.Get(2)) {
			log.Println("connection succeeded")
			break
		}
	}
	return true
}
