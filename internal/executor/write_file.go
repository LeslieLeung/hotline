package executor

import (
	"fmt"
	"github.com/leslieleung/hotline/internal/misc"
	"github.com/leslieleung/hotline/internal/ui"
)

// WriteFile writes content to a file.
// Params:
// - path [string, required]: The path of the file to write.
// - content [[]byte, required]: The content to write to the file.
// Output:
// - None
type WriteFile struct {
	Path    string
	Content string
}

var _ Executor = (*WriteFile)(nil)

func (w *WriteFile) BindParams(params map[string]interface{}) error {
	ui.Debugf("[write file] params: %+v\n", params)
	path := misc.GetString(params, "path")
	content := misc.GetString(params, "content")

	if path == "" {
		return fmt.Errorf("missing required parameter 'path'")
	}

	w.Path = path
	w.Content = content
	return nil

}

func (w *WriteFile) Execute() (map[string]interface{}, error) {
	return nil, misc.SafeWriteFile(w.Path, []byte(w.Content))
}
