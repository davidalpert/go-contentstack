package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/davidalpert/go-contentstack/internal/cfg"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
	"strings"
)

type ConfigSetupOptions struct {
	*printers.PrinterOptions
	Config *cfg.Config
}

func NewConfigSetupOptions(s printers.IOStreams) *ConfigSetupOptions {
	return &ConfigSetupOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("text"),
		Config:         &cfg.Config{},
	}
}

func NewCmdConfigSetup(s printers.IOStreams) *cobra.Command {
	o := NewConfigSetupOptions(s)
	var cmd = &cobra.Command{
		Use:   "setup",
		Short: "set up configuration",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Complete(cmd, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			return o.Run()
		},
	}

	o.AddPrinterFlags(cmd.Flags())

	return cmd
}

// Complete the options
func (o *ConfigSetupOptions) Complete(cmd *cobra.Command, args []string) error {
	if err := cfg.ReadMergedInto(o.Config); err != nil {
		return err
	}

	return nil
}

// Validate the options
func (o *ConfigSetupOptions) Validate() error {
	return o.PrinterOptions.Validate()
}

// Run the command
func (o *ConfigSetupOptions) Run() error {
	if err := askOneWithStreams(o.IOStreams, &survey.Input{
		Message: "ContentStack Region (EU | AZURE-NA | NA)",
		Default: o.Config.ContentStack.RegionCode,
	}, &o.Config.ContentStack.RegionCode, survey.WithValidator(survey.Required)); err != nil {
		return err
	}

	if err := askOneWithStreams(o.IOStreams, &survey.Input{
		Message: "ContentStack Content Delivery API Host",
		Default: o.Config.ContentStack.DeliveryApi.Host,
	}, &o.Config.ContentStack.DeliveryApi.Host, survey.WithValidator(survey.Required)); err != nil {
		return err
	}

	if err := askOneWithStreams(o.IOStreams, &survey.Input{
		Message: "ContentStack Content Management API Host",
		Default: o.Config.ContentStack.ManagementApi.Host,
	}, &o.Config.ContentStack.ManagementApi.Host, survey.WithValidator(survey.Required)); err != nil {
		return err
	}

	if err := askOneWithStreams(o.IOStreams, &survey.Input{
		Message: "ContentStack Content Management API Key",
		Default: o.Config.ContentStack.ManagementApi.Key,
	}, &o.Config.ContentStack.ManagementApi.Key, survey.WithValidator(survey.Required)); err != nil {
		return err
	}

	if err := askOneWithStreams(o.IOStreams, &survey.Input{
		Message: "ContentStack Content Management API Token",
		Default: o.Config.ContentStack.ManagementApi.Token,
	}, &o.Config.ContentStack.ManagementApi.Token, survey.WithValidator(survey.Required)); err != nil {
		return err
	}

	if err := o.Config.WriteToFile(cfg.File); err != nil {
		return err
	}
	_, err := fmt.Fprintf(o.Out, "configuration written to %s\n", cfg.File)

	return err
}

func validateMapping(ans interface{}) error {
	switch v := ans.(type) {
	case string:
		if strings.Count(v, ":") != 1 {
			return fmt.Errorf("expected a 'key:value' format")
		}
	default:
		return fmt.Errorf("expected a string")
	}

	return nil
}
