package version

import (
	"fmt"
	"os"

	annot "github.com/object88/tugboat-annotator"
	"github.com/object88/tugboat-annotator/cmd/common"
	"github.com/spf13/cobra"
)

func CreateCommand(cmmn *common.CommonArgs) *cobra.Command {
	c := &cobra.Command{
		Use:   "version",
		Short: "version will report the version of the tool",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(os.Stdout, annot.Version{}.String())
			return nil
		},
	}

	return c
}
