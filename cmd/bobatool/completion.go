// Code generated. DO NOT EDIT.

package main

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completionCmd)
}

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Emits bash a completion for bobatool",
	Long: `Enable bash completion like so:
		Linux:
			source <(bobatool completion)
		Mac:
			brew install bash-completion
			bobatool completion > $(brew --prefix)/etc/bash_completion.d/bobatool`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenBashCompletion(os.Stdout)
	},
}
