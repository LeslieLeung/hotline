package executor

import (
	"fmt"
	"github.com/leslieleung/hotline/internal/misc"
	"github.com/leslieleung/hotline/internal/ui"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Cmd runs a command.
// Params:
// - command [string, required]: The command to run.
// - split_output [string]: The delimiter to split the output by.
// Output:
// - stdout [string]: The standard output of the command.
// - items [[]string]: The output split by the delimiter.
type Cmd struct {
	Command     string
	SplitOutput string
}

var _ Executor = (*Cmd)(nil)

func (c *Cmd) BindParams(params map[string]interface{}) error {
	ui.Debugf("[cmd] params: %+v\n", params)
	command := misc.GetString(params, "command")
	splitOutput := misc.GetString(params, "split_output")

	if command == "" {
		return fmt.Errorf("missing required parameter 'command'")
	}

	c.Command = command
	c.SplitOutput = splitOutput
	return nil
}

func (c *Cmd) Execute() (map[string]interface{}, error) {
	cmd := exec.Command(getUserShell(), "-c", c.Command)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	items := make([]string, 0)
	if c.SplitOutput != "" {
		items = strings.Split(string(output), c.SplitOutput)
	}
	return map[string]interface{}{
		"stdout": string(output),
		"items":  items,
	}, nil
}

func getUserShell() string {
	switch runtime.GOOS {
	case "windows":
		if os.Getenv("COMSPEC") != "" {
			return os.Getenv("COMSPEC")
		}
		return "/cmd.exe"
	case "darwin":
		if os.Getenv("SHELL") != "" {
			return os.Getenv("SHELL")
		}
		return "/bin/zsh" // macOS's default shell since Catalina
	default:
		if os.Getenv("SHELL") != "" {
			return os.Getenv("SHELL")
		}
		return "/bin/sh"
	}
}
