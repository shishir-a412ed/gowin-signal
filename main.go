package main

import (
	"fmt"
	"log"
	"os"
//	"syscall"

	"github.com/urfave/cli/v2"
)

func main() {
	var pid int
	var signal string

	app := &cli.App{
		Name:    "gowin-signal",
		Version: "1.0.0",
		Usage:   "CLI tool for sending a signal to windows processes",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "pid",
				Usage:       "process ID",
				Destination: &pid,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "signal",
				Usage:       "signal to be send to process ID",
				Destination: &signal,
				Required:    true,
			},
		},
		Action: func(c *cli.Context) error {
			p, err := os.FindProcess(pid)
			if err != nil {
				return err
			}
			return p.Process.Signal(os.Interrupt)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("gowin-signal completed successfully.")
}
