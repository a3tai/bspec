package commands

import (
	"github.com/spf13/cobra"
)

var (
	// Version information (set by build process)
	Version   = "1.1.1"
	BuildDate = "unknown"
	GitCommit = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  `Display the version, build date, and git commit information for the BSpec CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		short, _ := cmd.Flags().GetBool("short")

		if short {
			cmd.Println(Version)
		} else {
			cmd.Printf("BSpec CLI version %s\n", Version)
			cmd.Printf("Build date: %s\n", BuildDate)
			cmd.Printf("Git commit: %s\n", GitCommit)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().BoolP("short", "s", false, "Show only the version number")
}