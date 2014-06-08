package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

// Register VM specific commands
func init() {
	app.Commands = append(app.Commands, cli.Command{
		Name:  "vm",
		Usage: "Virtual Machine commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				ShortName: "ls",
				Usage:     "List VMs",
				Action:    vmList,
			}, {
				Name:   "on",
				Action: vmOn,
				Usage:  "Power on Virtual Machine given by its managed object reference",
				Before: func(c *cli.Context) error {
					fmt.Println(c.Args())
					if c.Args().Get(1) == "" {
						// return errors.New("missing VM obj reference")
					}
					return nil
				},
			},
		},
	})
}
