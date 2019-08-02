package commands

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/instrumenta/conftest/pkg/commands/execute"
	"github.com/instrumenta/conftest/pkg/commands/pull"
	"github.com/instrumenta/conftest/pkg/commands/push"
	"github.com/instrumenta/conftest/pkg/commands/test"
	"github.com/instrumenta/conftest/pkg/commands/update"
	"github.com/instrumenta/conftest/pkg/constants"
	"github.com/instrumenta/conftest/pkg/report"
)

// NewDefaultCommand creates the default command
func NewDefaultCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "conftest <subcommand>",
		Short:   "Test your configuration files using Open Policy Agent",
		Version: fmt.Sprintf("Version: %s\nCommit: %s\nDate: %s\n", constants.Version, constants.Commit, constants.Date),
	}

	cmd.PersistentFlags().StringP("policy", "p", "policy", "path to the Rego policy files directory. For the test command, specifying a specific .rego file is allowed.")
	cmd.PersistentFlags().BoolP("debug", "", false, "enable more verbose log output")
	cmd.PersistentFlags().BoolP("trace", "", false, "enable more verbose trace output for rego queries")
	cmd.PersistentFlags().StringP("namespace", "", "main", "namespace in which to find deny and warn rules")
	cmd.PersistentFlags().BoolP("no-color", "", false, "disable color when printing")
	cmd.PersistentFlags().StringP("output", "o", "", fmt.Sprintf("output format for conftest results - valid options are: %s", report.ValidOutputs()))

	cmd.SetVersionTemplate(`{{.Version}}`)

	viper.BindPFlag("policy", cmd.PersistentFlags().Lookup("policy"))
	viper.BindPFlag("debug", cmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("trace", cmd.PersistentFlags().Lookup("trace"))
	viper.BindPFlag("namespace", cmd.PersistentFlags().Lookup("namespace"))
	viper.BindPFlag("no-color", cmd.PersistentFlags().Lookup("no-color"))
	viper.BindPFlag("output", cmd.PersistentFlags().Lookup("output"))

	viper.SetEnvPrefix("CONFTEST")
	viper.SetConfigName("conftest")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	cmd.AddCommand(test.NewTestCommand(
		os.Exit,
		test.GetOutputManager,
	))
	cmd.AddCommand(update.NewUpdateCommand())
	cmd.AddCommand(push.NewPushCommand())
	cmd.AddCommand(pull.NewPullCommand())

	if viper.GetBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.ErrorLevel)
	}

	return cmd
}
