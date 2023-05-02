/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"database/sql"
	"fmt"
	"github.com/cominotti/invowk/internal/tui/server"
	"github.com/spf13/cobra"
	"log"
	"os/exec"

	sqlite "modernc.org/sqlite"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")

		sqlcon, _ := sql.Open("sqlite", "")

		sqlcon.Close()

		s := sqlite.Error{}

		_ = s.Error()

		fmt.Println("Getting TUI server")
		router := server.NewHttpTuiServer()
		fmt.Println("Starting TUI server")
		router.Start()

		fmt.Println("Calling run.sh")
		cmd2 := exec.Command("./internal/cmd/run.sh")
		fmt.Println("Gonna check pid")
		err := cmd2.Start()

		fmt.Printf("PID before calling run is %d", cmd2.Process.Pid)
		fmt.Println("")

		cmd2.Wait()

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
