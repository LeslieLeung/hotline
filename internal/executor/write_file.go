package executor

import (
	"fmt"
	"github.com/leslieleung/hotline/internal/misc"
)

type WriteFile struct{}

func (w *WriteFile) Execute(params map[string]interface{}) (map[string]interface{}, error) {
	fmt.Printf("params: %+v\n", params)
	path := misc.GetString(params, "path")
	return nil, misc.SafeWriteFile(path, []byte(misc.GetString(params, "content")))
}
