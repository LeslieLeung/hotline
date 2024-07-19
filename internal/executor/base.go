package executor

import "fmt"

type Executor interface {
	BindParams(params map[string]interface{}) error
	Execute() (map[string]interface{}, error)
}

func GetExecutor(id string) (Executor, error) {
	var e Executor
	switch id {
	case "cmd":
		e = &Cmd{}
	case "dify_workflow":
		e = &DifyWorkflow{}
	case "print":
		e = &Print{}
	case "write-file":
		e = &WriteFile{}
	case "read-file":
		e = &ReadFile{}

	default:
		return nil, fmt.Errorf("executor %s not found", id)
	}
	return e, nil
}
