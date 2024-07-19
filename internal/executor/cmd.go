package executor

import (
	"github.com/leslieleung/hotline/internal/misc"
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
type Cmd struct{}

func (c *Cmd) Execute(params map[string]interface{}) (map[string]interface{}, error) {
	command := misc.GetString(params, "command")
	splitOutput := misc.GetString(params, "split_output")

	cmd := exec.Command(getUserShell(), "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	items := make([]string, 0)
	if splitOutput != "" {
		items = strings.Split(string(output), splitOutput)
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
