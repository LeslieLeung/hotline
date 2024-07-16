package executor

import (
	"fmt"
	"github.com/leslieleung/hotline/internal/misc"
)

type Print struct{}

func (c *Print) Execute(params map[string]interface{}) (map[string]interface{}, error) {
	fmt.Printf("%+v\n", misc.GetString(params, "message"))
	return nil, nil
}
