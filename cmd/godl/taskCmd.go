package main

import (
	"github.com/sftfjugg/godl/internal/errorHandle"
	"github.com/sftfjugg/godl/internal/resume"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(taskCmd)
}

var taskCmd = &cobra.Command{
	Use:     "task",
	Short:   "show current downloading task",
	Example: `godl task`,
	Run: func(cmd *cobra.Command, args []string) {
		errorHandle.ExitWithError(task())
	},
}

func task() error {
	err := resume.TaskPrint()
	return errors.WithStack(err)
}
