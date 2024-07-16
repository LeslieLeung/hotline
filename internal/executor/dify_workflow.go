package executor

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/leslieleung/dify-connector/pkg/dify"
	"github.com/leslieleung/hotline/internal/misc"
)

type DifyWorkflow struct{}

func (c *DifyWorkflow) Execute(params map[string]interface{}) (map[string]interface{}, error) {
	fmt.Printf("params: %+v\n", params)
	d := dify.New(misc.GetString(params, "host"), misc.GetString(params, "api_key"))
	inputs := params["inputs"].([]interface{})
	formatInput := make(map[string]interface{})
	for _, input := range inputs {
		input := input.(map[interface{}]interface{})
		inputKey := input["name"].(string)
		inputValue := input["value"]
		formatInput[inputKey] = inputValue
	}
	res, err := d.WorkflowRun(dify.WorkflowRunRequest{
		Inputs:       formatInput,
		ResponseMode: dify.ResponseModeBlocking,
		User:         uuid.New().String(),
	})
	if err != nil {
		return nil, err
	}
	return res.Data.Outputs, nil
}
