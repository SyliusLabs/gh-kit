package cmd

import "github.com/spf13/cobra"

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge a pull request",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)

	mergeCmd.PersistentFlags().StringP("category", "c", "", "Category of the pull request")
	mergeCmd.PersistentFlags().StringP("strategy", "s", "merge", "Merge strategy to be used")
}
