package cli

import (
	"github.com/spf13/cobra"
)

// RootCmd main cobra command
var RootCmd = &cobra.Command{
	Use:   "what-anime",
	Short: "Find the anime scene by image using your terminal",
	Long:  `Complete documentation is available at https://github.com/irevenko/what-anime-cli`,
}
