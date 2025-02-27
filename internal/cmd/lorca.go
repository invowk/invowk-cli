/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/invowk/invowk-cli/internal/wui/lorca"
	"github.com/urfave/cli/v2"
)

// LorcaCommand represents the lorca command
func LorcaCommand() *cli.Command {
	return &cli.Command{
		Name:  "lorca",
		Usage: "A brief description of your command",
		Action: func(c *cli.Context) error {
			lorca.Render()
			return nil
		},
	}
}
