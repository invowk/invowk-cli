/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"github.com/invowk/invowk-cli/internal/issue"
	"github.com/spf13/cobra"
)

// htermCmd represents the config command
var prettymdCmd = &cobra.Command{
	Use:   "prettymd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		issue.Handle(errors.New("hola"), issue.NewErrPrettyPrinter(issue.Get(issue.FileNotFoundId)))
	},
}

func init() {
	rootCmd.AddCommand(prettymdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// htermCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// htermCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
