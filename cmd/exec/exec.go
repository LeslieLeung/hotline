package exec

import (
	"github.com/leslieleung/hotline/internal/ui"
	"github.com/leslieleung/hotline/internal/workflow"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"os"
)

var Cmd = &cobra.Command{
	Use:   "exec workflow [flags]",
	Short: "Execute a workflow",
	Run:   run,
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		UnknownFlags: true,
	},
}

var workflowFile string

func run(_ *cobra.Command, args []string) {
	if workflowFile == "" {
		ui.ErrorfExit("workflow file is required")
	}
	// open workflow file
	fbs, err := os.ReadFile(workflowFile)
	var wfs workflow.Spec
	err = yaml.Unmarshal(fbs, &wfs)
	if err != nil {
		ui.ErrorfExit("error reading workflow file: %s", err)
	}
	// find the workflow
	var run *workflow.Run
	for _, wf := range wfs.Workflows {
		if wf.ID == args[0] {
			// execute the workflow
			run = workflow.NewRun(wf)
			err := run.Execute()
			if err != nil {
				ui.ErrorfExit("error executing workflow: %s", err)
			}
			return
		}
	}
}

func init() {
	Cmd.Flags().StringVarP(&workflowFile, "file", "f", "", "workflow file")
}
