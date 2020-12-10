package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "umigame",
	Short: "single-play umigame game",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Green("root command"))
	},
}
