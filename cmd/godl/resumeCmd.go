package main

import (
	"github.com/sftfjugg/godl/internal/errorHandle"
	"github.com/sftfjugg/godl/internal/executioner"
	"github.com/sftfjugg/godl/internal/resume"
	"github.com/sftfjugg/godl/internal/tool"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(resumeCmd)
}

var resumeCmd = &cobra.Command{
	Use:     "resume",
	Short:   "resume downloading task",
	Example: `godl resume URL or file name`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		errorHandle.ExitWithError(resumeTask(args))
	},
}

func resumeTask(args []string) error {
	task := ""
	if tool.IsVaildURL(args[0]) {
		task = tool.GetFilenameFrom(args[0])
	} else {
		task = args[0]
	}

	state, err := resume.Resume(task)
	if err != nil {
		return errors.WithStack(err)
	}
	return executioner.Do(state.URL, state, conc)
}
