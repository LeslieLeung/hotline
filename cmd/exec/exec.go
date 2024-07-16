package exec

import (
	"fmt"
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
		panic("workflow file is required")
	}
	// open workflow file
	fbs, err := os.ReadFile(workflowFile)
	var wfs workflow.Spec
	err = yaml.Unmarshal(fbs, &wfs)
	if err != nil {
		panic(err)
	}
	// find the workflow
	var run *workflow.Run
	for _, wf := range wfs.Workflows {
		if wf.ID == args[0] {
			// execute the workflow
			run = workflow.NewRun(wf)
			err := run.Execute()
			if err != nil {
				fmt.Printf("error executing workflow: %s\n", err)
				os.Exit(1)
			}
			return
		}
	}
}

func init() {
	Cmd.Flags().StringVarP(&workflowFile, "file", "f", "", "workflow file")
}
