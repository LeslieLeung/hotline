package executor

import (
	"github.com/google/uuid"
	"github.com/leslieleung/dify-connector/pkg/dify"
	"github.com/leslieleung/hotline/internal/misc"
	"github.com/leslieleung/hotline/internal/ui"
)

// DifyWorkflow executes a Dify workflow.
// Params:
// - host [string, required]: The host of the Dify instance.
// - api_key [string, required]: The API key of the Dify workflow.
// - inputs [map[string]interface{}]: The inputs of the workflow.
// Output:
// - [map[string]interface{}]: The outputs of the workflow.
type DifyWorkflow struct {
	Host   string `json:"host"`
	ApiKey string
	Inputs map[string]interface{}
}

var _ Executor = (*DifyWorkflow)(nil)

func (c *DifyWorkflow) BindParams(params map[string]interface{}) error {
	ui.Debugf("[Dify Workflow] params: %+v\n", params)
	c.Host = misc.GetString(params, "host")
	c.ApiKey = misc.GetString(params, "api_key")
	inputs := params["inputs"].([]interface{})
	formatInput := make(map[string]interface{})
	for _, input := range inputs {
		input := input.(map[interface{}]interface{})
		inputKey := input["name"].(string)
		inputValue := input["value"]
		formatInput[inputKey] = inputValue
	}
	c.Inputs = formatInput
	return nil
}

func (c *DifyWorkflow) Execute() (map[string]interface{}, error) {
	d := dify.New(c.Host, c.ApiKey)
	res, err := d.WorkflowRun(dify.WorkflowRunRequest{
		Inputs:       c.Inputs,
		ResponseMode: dify.ResponseModeBlocking,
		User:         uuid.New().String(),
	})
	if err != nil {
		return nil, err
	}
	return res.Data.Outputs, nil
}
