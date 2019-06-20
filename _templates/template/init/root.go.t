---
to: cmd/root.go
---

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "<%= name %>",
	Short: "A <%= name %> project",
	Long:  `The long description`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

// Execute the current command
func Execute() {
       if err := rootCmd.Execute(); err != nil {
               fmt.Println(err)
               os.Exit(1)
       }
}
