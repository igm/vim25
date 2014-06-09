package main

import "github.com/codegangsta/cli"

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
			}, {
				Name:        "import",
				Action:      vmImport,
				Usage:       "Import Virtual Machine using OVF",
				Description: "Import Virtual Machine using OVF file",
				Flags: []cli.Flag{
					cli.StringFlag{"resourcepool, r", "", "resource pool obj reference"},
					cli.StringFlag{"datastore, d", "", "datastore obj reference"},
					cli.StringFlag{"folder, f", "", "folder obj reference"},
				},
			},
		},
	})
}
