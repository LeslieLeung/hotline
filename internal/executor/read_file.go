package executor

import (
	"github.com/leslieleung/hotline/internal/misc"
	"os"
)

// ReadFile reads the content of a file.
// Params:
// - path [string, required]: The path to the file to read.
// Output:
// - content [[]byte]: The content of the file.
type ReadFile struct{}

func (r *ReadFile) Execute(params map[string]interface{}) (map[string]interface{}, error) {
	path := misc.GetString(params, "path")
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"content": string(content),
	}, nil
}
