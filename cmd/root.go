package cmd

import (
	"fmt"
	"github.com/A1esandr/korean/internal/app"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "korea",
		Short: "Tool for korean language",
		Run: func(cmd *cobra.Command, args []string) {
			Word("first")
		},
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Word(word string) {
	w := app.ReadCsv("csv/" + word + ".csv")
	app.Check(w)
}
