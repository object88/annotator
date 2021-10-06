package main

import (
	"context"
	"fmt"
	"os"

	"github.com/object88/tugboat-annotator/cmd"
)

func main() {
	rootCmd := cmd.InitializeCommands()
	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		fmt.Println(err)
		// if ece, ok := err.(*errs.ExitCodeError); ok {
		// 	os.Exit(ece.ExitCode)
		// }
		os.Exit(-1)
	}
}
