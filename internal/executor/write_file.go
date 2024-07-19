package executor

import (
	"github.com/leslieleung/hotline/internal/misc"
	"github.com/leslieleung/hotline/internal/ui"
)

// WriteFile writes content to a file.
// Params:
// - path [string, required]: The path of the file to write.
// - content [[]byte, required]: The content to write to the file.
// Output:
// - None
type WriteFile struct{}

func (w *WriteFile) Execute(params map[string]interface{}) (map[string]interface{}, error) {
	ui.Debugf("[write file] params: %+v\n", params)
	path := misc.GetString(params, "path")
	return nil, misc.SafeWriteFile(path, []byte(misc.GetString(params, "content")))
}
