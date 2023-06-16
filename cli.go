package main

import (
	"strconv"

	"github.com/spf13/cobra"
)

var (
	flagDebug bool
)

var rootCmd = &cobra.Command{
	Use: "raf-utils ...",
}

var cmdRetab = &cobra.Command{
	Use:   "retab <file> <from-indent> <to-indent> <tab-space-count>",
	Short: "Retabs things",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		err := Retab(args[0], RetabConfig{
			FromIndent:    fromInt(args[1]),
			ToIndent:      fromInt(args[2]),
			TabSpaceCount: fromInt(args[3]),
		})
		checkErrorDie(err)
	},
}

var cmdYaml = &cobra.Command{
	Use: "yaml ...",
}

var cmdYamlLoc = &cobra.Command{
	Use:   "loc <file> <path>",
	Short: "Gets the line of a path",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := yamlLoc(args[0], args[1])
		checkErrorDie(err)
	},
}

var cmdYamlPos = &cobra.Command{
	Use:   "pos <file> <x> <y>",
	Short: "Gets the path of a cursor pos",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		x, err := strconv.Atoi(args[1])
		checkErrorDie(err)

		y, err := strconv.Atoi(args[2])
		checkErrorDie(err)

		err = yamlPos(args[0], x, y)
		checkErrorDie(err)
	},
}

func init() {
	rootCmd.AddCommand(cmdRetab)
	rootCmd.AddCommand(cmdYaml)
	cmdYaml.AddCommand(cmdYamlLoc)
	cmdYaml.AddCommand(cmdYamlPos)
	rootCmd.PersistentFlags().BoolVar(&flagDebug, "debug", false, "Enable debug output")
}
