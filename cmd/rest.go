package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.warungpintar.co/sales-platform/brook/ports/rest"
)

// restCmd represents the rest command
var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "Running rest service",
	Run: func(cmd *cobra.Command, args []string) {
		if err := rest.Application(rest.GetDefaultConfig(cfg)); err != nil {
			panic(fmt.Errorf("cannot start rest server: %w", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(restCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
