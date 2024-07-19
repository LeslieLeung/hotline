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
type Print struct {
	Message string
}

var _ Executor = (*Print)(nil)

func (c *Print) BindParams(params map[string]interface{}) error {
	message := misc.GetString(params, "message")

	c.Message = message
	return nil
}

func (c *Print) Execute() (map[string]interface{}, error) {
	fmt.Printf("%+v\n", c.Message)
	return nil, nil
}
