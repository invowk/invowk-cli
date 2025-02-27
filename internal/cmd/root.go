/*
Copyright © 2022 Danilo Cominotti Marques <dcominottim@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var cfgFile string

// Execute initializes and runs the CLI application.
func Execute() {
	app := &cli.App{
		Name:  "invowk",
		Usage: "A code/script execution & sharing engine exposed as an user-extensible CLI",
		Description: `
invowk is a code execution & sharing engine exposed as an user-extensible CLI.

Why is it an 'execution engine'?
Because invowk helps you run code/scripts encapsulated as custom user-made 
commands! Such commands can be written as/in:
- shell scripts (e.g.: bash, zsh, cmd, powershell)
- dynamic languages (e.g.: python, js/node.js, ruby)
- compiled languages (e.g.: java, go, rust, c/c++)
They can be run against the host machine or a Docker container, all in the
most seamless fashion.

Why is it a 'sharing engine'?
Because invowk helps you import/export custom user-made commands so that you
can easily run other people's code/automations/ideas and vice-versa!
We even have several curated custom commands that you can try out-of-the-box.

Why is it an 'user-extensible CLI'?
Because invowk is a kind of 'meta-CLI': it is designed so that its users can 
extend it with custom commands that add the real value to the CLI and invowk
ecosystem! It also makes receiving (via flags or fancy interactive prompts)
and validating inputs a complete breeze.
`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Usage:       "config file (the default one is located in $HOME/.invowk/config.toml)",
				Destination: &cfgFile,
			},
			&cli.BoolFlag{
				Name:  "toggle",
				Usage: "Help message for toggle",
			},
		},
		Before: func(c *cli.Context) error {
			initConfig()
			return nil
		},
		Action: func(c *cli.Context) error {
			// Define the action for the root command here
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}
