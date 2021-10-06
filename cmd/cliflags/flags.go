package cliflags

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// CLI Flags
const (
	// VerboseKey turns on debugging output
	verboseKey = "verbose"
)

type FlagManager struct {
	verbose bool
}

func NewFlagManager() *FlagManager {
	return &FlagManager{}
}

func (fl *FlagManager) ConfigureVerboseFlag(flags *pflag.FlagSet) {
	flags.BoolVar(&fl.verbose, verboseKey, false, "emit debug messages")
	viper.BindEnv(verboseKey)
	viper.BindPFlag(verboseKey, flags.Lookup(verboseKey))
}

func (fl *FlagManager) Verbose() bool {
	return viper.GetBool(verboseKey)
}
