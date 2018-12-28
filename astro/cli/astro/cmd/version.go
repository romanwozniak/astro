package cmd

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	// When a release happens, the value of this variable will be overwritten
	// by the linker to match the release version.
	version = "dev"
	commit  = currentGitHash()
	date    = currentDate()
)

// currentGitHash returns the git hash of the current working directory or an
// empty string if the hash can't be determined (i.e. git is not installed).
func currentGitHash() string {
	git := exec.Command("git", "rev-parse", "--short", "HEAD")

	output, err := git.Output()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(output))
}

// currentDate returns the current date in a format suitable for printing in
// the version output.
func currentDate() string {
	return time.Now().Format(time.RFC3339)
}

var versionCmd = &cobra.Command{
	Use:                   "version",
	DisableFlagsInUseLine: true,
	Short:                 "Print astro version",
	RunE: func(cmd *cobra.Command, args []string) error {
		versionString := []string{
			"astro version",
			version,
		}

		if commit != "" {
			versionString = append(versionString, fmt.Sprintf("(%s)", commit))
		} else {
			versionString = append(versionString, "(hash unknown)")
		}

		versionString = append(versionString, fmt.Sprintf("built %s", date))

		println(strings.Join(versionString, " "))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
