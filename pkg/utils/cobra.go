package utils

import (
	"github.com/spf13/pflag"
)

func AddStringFlag(
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

func AddBoolFlag(
	holder *bool,
	flags *pflag.FlagSet,
	namespaceFlag string,
	short string,
	defaultValue bool,
	description string,
) {
	flags.BoolVarP(
		holder,
		namespaceFlag,
		short,
		defaultValue,
		description,
	)
}

func AddStringSliceFlag(
	holder *[]string,
	flags *pflag.FlagSet,
	namespaceFlag string,
	short string,
	defaultValue []string,
	description string,
) {
	flags.StringSliceVarP(
		holder,
		namespaceFlag,
		short,
		defaultValue,
		description,
	)
}
