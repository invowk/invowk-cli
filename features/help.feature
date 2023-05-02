Feature: Help flag
  In order to know invowk's features
  As an invowk user
  I must see info about possible usages when executing the root command with the commonly used '-h' flag

  Background:
    Given an open TTY

  Scenario: Show help information
    When I successfully run "--help"
    Then the terminal output should contain:
    """
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

    Usage:
      invowk [command]

    Available Commands:
      completion  Generate the autocompletion script for the specified shell
      config      A brief description of your command
      help        Help about any command
      hterm       A brief description of your command
      lorca       A brief description of your command
      prettymd    A brief description of your command
      tui         A brief description of your command
      web         A brief description of your command

    Flags:
          --config string   config file (default is $HOME/.config/invowk/invowk.toml)
      -h, --help            help for invowk
      -t, --toggle          Help message for toggle

    Use "invowk [command] --help" for more information about a command.
    """

  Scenario: Show Markdown text
    When I successfully run "prettymd"
    Then the TUI output should contain the content from "prettymd.txt", which renders to:
    """
       Dang, we have run into an issue!

      We have failed to start our super powered TUI Server due to weird
      conditions.

      ## Things you can try to fix and retry

      • Run this command

        $ invowk fix

        and try again what you doing before.
    """