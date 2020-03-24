package application

import (
	"flag"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	_ "huaweiApi/pkg/utils/log"
)

const (
	FlagLogLevel   = "log-level"
	FlagConfigFile = "config-file"
	FlagServiceId  = "services-id"
)

func GetCommand() *cobra.Command {

	setFlags()
	return &cobra.Command{
		Use:   "services",
		Short: "this is a services that provides example RPC and RESTful APIs",
		Long:  `this is a services that provides example RPC API and RESTful APIs .`,
		Run:   NewApp().run,
	}
}

func setFlags() {

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.String(FlagLogLevel, "info", "log level(options: trace, debug, info, warn|warning, error, fatal, panic)")
	pflag.String(FlagConfigFile, "", "config file for json format")
	pflag.Uint16(FlagServiceId, 0, "distinguish between different services when highly available.")
}
