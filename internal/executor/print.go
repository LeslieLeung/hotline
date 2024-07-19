package executor

import (
	"fmt"
	"github.com/leslieleung/hotline/internal/misc"
)

// Print prints a message to the console.
// Params:
// - message [string, required]: The message to print.
// Output:
// - None
type Print struct{}

func (c *Print) Execute(params map[string]interface{}) (map[string]interface{}, error) {
	fmt.Printf("%+v\n", misc.GetString(params, "message"))
	return nil, nil
}
