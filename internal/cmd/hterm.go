/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/invowk/invowk-cli/internal/wui/gotty"
	"github.com/urfave/cli/v2"
)

// HtermCommand represents the hterm command
func HtermCommand() *cli.Command {
	return &cli.Command{
		Name:  "hterm",
		Usage: "A brief description of your command",
		Action: func(c *cli.Context) error {
			gotty.Main()
			return nil
		},
	}
}
