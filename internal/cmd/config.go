/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/invowk/invowk-cli/internal/tui/server"
	"github.com/urfave/cli/v2"
	sqlite "modernc.org/sqlite"
)

// ConfigCommand represents the config command
func ConfigCommand() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "A brief description of your command",
		Action: func(c *cli.Context) error {
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

			cmd2 := exec.Command("python", "F:\\Repositories\\github\\invowk\\invowk-cli\\sample.py")
			cmd2.Stdout = os.Stdout
			cmd2.Stderr = os.Stderr

			fmt.Println("Gonna check pid")
			err := cmd2.Start()

			fmt.Printf("PID before calling run is %d", cmd2.Process.Pid)
			fmt.Println("")

			cmd2.Wait()

			if err != nil {
				log.Fatal(err)
			}

			return nil
		},
	}
}
