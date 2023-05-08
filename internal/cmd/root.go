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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "invowk",
	Short: "A code/script execution & sharing engined exposed as an user-extensible CLI",
	Long: `
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
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (the default one is located in $HOME/.invowk/config.toml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//rootCmd.SetUsageFunc(boa.UsageFunc)
	//rootCmd.SetHelpFunc(boa.HelpFunc)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		userCfgDir, err := os.UserConfigDir()
		cobra.CheckErr(err)
		// Search config in config directory with name ".invowk" (without extension).
		viper.AddConfigPath(filepath.Join(userCfgDir, ".invowk"))
		viper.SetConfigType("toml")
		viper.SetConfigName(".config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file: ", viper.ConfigFileUsed())
	}
}
