package cmd

import (
	"fmt"
	"github.com/davidalpert/go-contentstack/internal/cfg"
	"github.com/davidalpert/go-contentstack/internal/diagnostics"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
	"os"
)

// Execute builds the default root command and invokes it with os.Args
func Execute() {
	s := printers.DefaultOSStreams()
	// configure the logger here in the outer scope so that we can defer
	// any cleanup such as writing/flushing the stream
	logCleanupFn := diagnostics.ConfigureLogger(s)
	defer logCleanupFn()

	rootCmd := NewRootCmd(s)

	rootCmd.SetArgs(os.Args[1:]) // without program

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(s.Out, err)
		os.Exit(1)
	}
}

// NewRootCmd creates the 'root' command and configures it's nested children
func NewRootCmd(s printers.IOStreams) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "cs",
		Short:         "a CLI tool for interacting with the ContentStack APIs",
		SilenceUsage:  false,
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   true,
			DisableNoDescFlag:   false,
			DisableDescriptions: false,
			HiddenDefaultCmd:    false,
		},
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//	Run: func(cmd *cobra.Command, args []string) { },
		//  RunE: func(cmd *cobra.Command, args []string) error { },
		Aliases: []string{},
	}

	// Register subcommands
	rootCmd.AddCommand(NewCmdConfig(s))
	rootCmd.AddCommand(NewCmdContentType(s))
	rootCmd.AddCommand(NewCmdEnvironment(s))
	rootCmd.AddCommand(NewCmdGlobal(s))
	rootCmd.AddCommand(NewCmdLocale(s))
	rootCmd.AddCommand(NewCmdVersion(s))

	rootCmd.PersistentFlags().StringVarP(&cfg.File, "config", "c", cfg.File, "config file to use")

	return rootCmd
}

func init() {
}
