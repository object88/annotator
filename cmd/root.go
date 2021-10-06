package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/object88/tugboat-annotator/cmd/common"
	"github.com/object88/tugboat-annotator/cmd/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// InitializeCommands sets up the cobra commands
func InitializeCommands() *cobra.Command {
	common, rootCmd := createRootCommand()

	rootCmd.AddCommand(
		version.CreateCommand(common),
	)

	return rootCmd
}

func createRootCommand() (*common.CommonArgs, *cobra.Command) {
	var ca *common.CommonArgs

	var start time.Time
	cmd := &cobra.Command{
		Use:   "tugboat-annotator",
		Short: "tugboat-annotator uses AWS secrets to access databaes",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			start = time.Now()
			if err := ca.Evaluate(); err != nil {
				return fmt.Errorf("failed to evaluate common arguments: %w", err)
			}

			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, _ []string) error {
			duration := time.Since(start)

			segments := []string{}
			var f func(c1 *cobra.Command)
			f = func(c1 *cobra.Command) {
				parent := c1.Parent()
				if parent != nil {
					f(parent)
				}
				segments = append(segments, c1.Name())
			}
			f(cmd)

			ca.Logger.Info("Executed command.", "command", strings.Join(segments, " "), "duration", duration)
			return nil
		},
	}

	viper.SetEnvPrefix("tugboat-annotator")

	ca = common.NewCommonArgs()

	flags := cmd.PersistentFlags()

	ca.FlagMan.ConfigureVerboseFlag(flags)

	return ca, cmd
}
