package executor

import (
	"fmt"
	"github.com/leslieleung/hotline/internal/misc"
	"github.com/leslieleung/hotline/internal/ui"
	"os"
)

// ReadFile reads the content of a file.
// Params:
// - path [string, required]: The path to the file to read.
// Output:
// - content [[]byte]: The content of the file.
type ReadFile struct {
	Path string
}

var _ Executor = (*ReadFile)(nil)

func (r *ReadFile) BindParams(params map[string]interface{}) error {
	ui.Debugf("[read file] params: %+v\n", params)
	path := misc.GetString(params, "path")

	if path == "" {
		return fmt.Errorf("missing required parameter 'path'")
	}

	r.Path = path
	return nil
}

func (r *ReadFile) Execute() (map[string]interface{}, error) {
	content, err := os.ReadFile(r.Path)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"content": string(content),
	}, nil
}
