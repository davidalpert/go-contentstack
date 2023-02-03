package cmd

import (
	"fmt"
	"github.com/davidalpert/go-contentstack/internal/cfg"
	"github.com/davidalpert/go-contentstack/v1/management"
	"github.com/davidalpert/go-contentstack/v1/schema"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

type ContentTypeCreateOptions struct {
	*printers.PrinterOptions
	Config cfg.Config
	Client *management.Client
	UID    string
}

func NewContentTypeCreateOptions(s printers.IOStreams) *ContentTypeCreateOptions {
	return &ContentTypeCreateOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("yaml"),
	}
}

func NewCmdContentTypeCreate(s printers.IOStreams) *cobra.Command {
	o := NewContentTypeCreateOptions(s)
	var cmd = &cobra.Command{
		Use:   "create <uid>",
		Short: "create a sample global field with the given UID",
		Args:  cobra.ExactArgs(1),
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
func (o *ContentTypeCreateOptions) Complete(cmd *cobra.Command, args []string) error {
	o.UID = args[0]

	if err := cfg.ReadMergedInto(&o.Config); err != nil {
		return err
	}

	c, err := management.NewClient(&o.Config.ContentStack.ManagementApi)
	if err != nil {
		return err
	}
	o.Client = c

	return nil
}

// Validate the options
func (o *ContentTypeCreateOptions) Validate() error {
	if err := o.Config.Validate(); err != nil {
		return err
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *ContentTypeCreateOptions) Run() error {
	if o.FormatCategory() == "table" || o.FormatCategory() == "csv" {
		o.WithDefaultOutput("json")
	}

	t := schema.ContentType{
		Title: o.UID,
		UID:   o.UID,
		Schema: []schema.Field{
			schema.Field{
				DataType:    "string",
				DisplayName: "Title",
				Uid:         "title",
				FieldMetadata: schema.FieldMetadata{
					Description:  "title of the test entity",
					Default:      schema.BoolPtr(true),
					DefaultValue: "something amazing",
				},
				Mandatory: true,
				Multiple:  false,
			},
		},
		Description: fmt.Sprintf("test type %#v", o.UID),
	}

	r, err := o.Client.CreateContentType(&t)
	if err != nil {
		return err
	}

	return o.WriteOutput(r)
}
