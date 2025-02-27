/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/invowk/invowk-cli/internal/tui/bubble/textinput"
	"github.com/urfave/cli/v2"
)

// TuiCommand represents the tui command
func TuiCommand() *cli.Command {
	return &cli.Command{
		Name:  "tui",
		Usage: "A brief description of your command",
		Action: func(c *cli.Context) error {
			textinput.Bubble()
			return nil
		},
	}
}
