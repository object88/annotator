package common

import (
	"strings"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/object88/tugboat-annotator/cmd/cliflags"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type CommonArgs struct {
	FlagMan *cliflags.FlagManager
	Logger  logr.Logger
}

func NewCommonArgs() *CommonArgs {
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	ca := &CommonArgs{
		FlagMan: cliflags.NewFlagManager(),
	}

	return ca
}

func (ca *CommonArgs) Evaluate() error {
	apLog, err := zap.NewProduction()
	if ca.FlagMan.Verbose() {
		apLog, err = zap.NewDevelopment()
	}
	if err != nil {
		return err
	}

	ca.Logger = zapr.NewLogger(apLog)
	return nil
}
