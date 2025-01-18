package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use:     "apollocli",
	Version: version,
	Short:   "apollocli is a cli tool for the apollo project",
	Long:    "apollocli is a cli tool for the apollo project - api service, data, monitoring",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Interact with the API service",
}

func Execute() {

	inspectCmd.Flags().BoolVarP(&onlyDigits, "digits", "d", false, "Count only digits")

	apiCmd.AddCommand(reverseCmd)
	apiCmd.AddCommand(inspectCmd)
	apiCmd.AddCommand(healthCmd)

	rootCmd.AddCommand(apiCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing the apollo cli '%s'\n", err)
		os.Exit(1)
	}
}

var reverseCmd = &cobra.Command{
	Use:     "reverse",
	Aliases: []string{"rev"},
	Short:   "Reverse a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := reverse(args[0])
		fmt.Println(res)
	},
}

var onlyDigits bool
var inspectCmd = &cobra.Command{
	Use:     "inspect",
	Aliases: []string{"insp"},
	Short:   "Inspects a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		i := args[0]
		res, kind := inspect(i, onlyDigits)

		pluralS := "s"
		if res == 1 {
			pluralS = ""
		}
		fmt.Printf("'%s' has %d %s%s.\n", i, res, kind, pluralS)
	},
}

var healthCmd = &cobra.Command{
	Use:     "health",
	Aliases: []string{"healthcheck", "status"},
	Short:   "Check the health status of the API service",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
