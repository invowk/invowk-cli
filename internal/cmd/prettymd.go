/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"

	"github.com/invowk/invowk-cli/internal/issue"
	"github.com/urfave/cli/v2"
)

// PrettyMdCommand represents the prettymd command
func PrettyMdCommand() *cli.Command {
	return &cli.Command{
		Name:  "prettymd",
		Usage: "A brief description of your command",
		Action: func(c *cli.Context) error {
			issue.Handle(errors.New("hola"), issue.NewErrPrettyPrinter(issue.Get(issue.FileNotFoundId)))
			return nil
		},
	}
}
