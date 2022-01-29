package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

func isCmdPath(cmd *cobra.Command, path string) bool {
	return strings.Contains(cmd.CommandPath(), path)
}

var GetFoodDynamic = func(cmd *cobra.Command) [][]string {
	if !isCmdPath(cmd, "get food") {
		return nil
	}

	return [][]string{
		[]string{"apple", "Green apple"},
		[]string{"tomato", "Red tomato"},
	}
}

var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get something",
	Aliases: []string{"eat"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var getFoodCmd = &cobra.Command{
	Use:   "food",
	Short: "Get some food",
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		for _, v := range args {
			if verbose {
				cmd.Println("Here you go, take this:", v)
			} else {
				cmd.Println(v)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getFoodCmd)
	getCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose log")
}
