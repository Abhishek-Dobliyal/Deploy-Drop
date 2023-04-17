package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/kataras/tablewriter"
	"github.com/lensesio/tableprinter"
	"github.com/spf13/cobra"
)

var (
	Red     = color.New(color.FgRed, color.Bold).SprintfFunc()
	Cyan    = color.New(color.FgCyan, color.Bold).SprintfFunc()
	Yellow  = color.New(color.FgYellow, color.Italic).SprintfFunc()
	Printer = tableprinter.New(os.Stdout)

	rootCmd = &cobra.Command{
		Use:   "deploy-drop",
		Short: "A CLI for automatic removal of deployments on Github.",
		Long:  `Remove deployments of unused projects on Github. Uses official Github API & it's token to accomplish the task. `,
	}
)

// Execute executes the root command.
func Execute() error {
	Printer.BorderTop, Printer.BorderBottom = true, true
	Printer.BorderLeft, Printer.BorderRight = true, true
	Printer.CenterSeparator = "│"
	Printer.ColumnSeparator = "│"
	Printer.RowSeparator = "─"
	Printer.HeaderBgColor = tablewriter.BgHiMagentaColor
	Printer.HeaderFgColor = tablewriter.FgHiWhiteColor
	Printer.AutoFormatHeaders = true
	Printer.AutoWrapText = true
	Printer.RowLine = true
	return rootCmd.Execute()
}
