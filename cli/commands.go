package cli

import (
	"github.com/spf13/cobra"
)

// SearchByFile Command
var SearchByFile = &cobra.Command{
	Use:   "file",
	Short: "Search for the anime scene by existing image file",
	Long:  `what-anime file <PATH_TO_IMAGE>`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		SearchByImageFile(args[0])
	},
}

// SearchByLink command
var SearchByLink = &cobra.Command{
	Use:   "link",
	Short: "Search for the anime scene by existing image url",
	Long:  `what-anime link <IMAGE_URL>`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		SearchByImageLink(args[0])
	},
}

// AddCommands launches all commands
func AddCommands() {
	RootCmd.AddCommand(SearchByFile)
	RootCmd.AddCommand(SearchByLink)
}
