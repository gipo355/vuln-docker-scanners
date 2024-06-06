package utils

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func AddFlag(
	holder *string,
	flags *pflag.FlagSet,
	namespaceFlag string,
	short string,
	defaultValue string,
	description string,
) {
	flags.StringVarP(
		holder,
		namespaceFlag,
		short,
		defaultValue,
		description,
	)
}

func AddPersistentFlag(
	holder *string,
	flags *pflag.FlagSet,
	namespaceFlag string,
	short string,
	defaultValue string,
	description string,
) {
	flags.StringVarP(
		holder,
		namespaceFlag,
		short,
		defaultValue,
		description,
	)
}

func MarkFlagRequired(cmd *cobra.Command, flag string) error {
	return fmt.Errorf("error marking flag required: %w", cmd.MarkFlagRequired(flag))
}
