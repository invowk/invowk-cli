/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	browser "github.com/invowk/invowk-cli/internal/wui"
	"github.com/urfave/cli/v2"
)

// WebCommand represents the web command
func WebCommand() *cli.Command {
	return &cli.Command{
		Name:  "web",
		Usage: "A brief description of your command",
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				fmt.Println("gonna open " + c.Args().Get(0))
				browser.Open(c.Args().Get(0))
			}
			return nil
		},
	}
}
