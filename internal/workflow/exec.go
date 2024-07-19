package workflow

import (
	"fmt"
	"github.com/leslieleung/hotline/internal/executor"
	"github.com/leslieleung/hotline/internal/ui"
	"github.com/spf13/pflag"
	"os"
	"regexp"
	"strings"
)

type Run struct {
	Workflow Workflow

	inputs   map[string]interface{}
	inputMap map[string]*string
	steps    map[string]interface{}
}

func NewRun(workflow Workflow) *Run {
	return &Run{
		Workflow: workflow,
		inputs:   make(map[string]interface{}),
		inputMap: make(map[string]*string),
		steps:    make(map[string]interface{}),
	}
}

func (r *Run) parseFlags() error {
	fs := pflag.NewFlagSet("workflow", pflag.ContinueOnError)
	for _, input := range r.Workflow.Inputs {
		r.inputMap[input.Name] = new(string)
		fs.StringVar(r.inputMap[input.Name], input.Name, input.Default, input.Description)
	}
	fs.ParseErrorsWhitelist.UnknownFlags = true
	err := fs.Parse(os.Args)
	if err != nil {
		return err
	}

	for key, value := range r.inputMap {
		r.inputs[key] = *value
	}

	// validate inputs
	for _, input := range r.Workflow.Inputs {
		if input.Required && r.inputs[input.Name] == "" {
			return fmt.Errorf("input %s is required", input.Name)
		}
	}

	ui.Debugf("Inputs args: [%+v]", r.inputs)
	return nil
}

func (r *Run) Execute() error {
	// parse flags
	err := r.parseFlags()
	if err != nil {
		return err
	}

	// execute steps
	return r.executeSteps()
}

func (r *Run) executeSteps() error {
	for _, step := range r.Workflow.Steps {
		exec, err := executor.GetExecutor(step.Uses)
		if err != nil {
			return err
		}

		r.fillVariables(step.With)
		ui.Debugf("Executing step [name: %s, id: %s, uses: %s, params: %+v]",
			step.Name, step.ID, step.Uses, step.With)
		err = exec.BindParams(step.With)
		if err != nil {
			return err
		}
		output, err := exec.Execute()
		if err != nil {
			return err
		}
		ui.Debugf("Step executed successfully [name: %s, id: %s, uses: %s, output: %+v]",
			step.Name, step.ID, step.Uses, output)
		// store output to steps cache
		r.steps[step.ID] = make(map[string]interface{})
		r.steps[step.ID].(map[string]interface{})["outputs"] = output
	}

	return nil
}

func (r *Run) fillVariables(with interface{}) interface{} {
	variables := map[string]interface{}{
		"steps":  r.steps,
		"inputs": r.inputs,
	}

	placeholderRegex := regexp.MustCompile(`\${{\s*(.+?)\s*}}`)

	switch withTyped := with.(type) {
	case string:
		matches := placeholderRegex.FindAllStringSubmatch(withTyped, -1)
		for _, match := range matches {
			value, err := getValueByPath(variables, match[1])
			if err != nil {
				ui.Errorf("error getting value by path: %s\n", err)
				continue
			}
			withTyped = strings.ReplaceAll(withTyped, match[0], fmt.Sprint(value))
		}
		return withTyped
	case map[string]interface{}:
		for key, value := range withTyped {
			withTyped[key] = r.fillVariables(value)
		}
	case []interface{}:
		for i, value := range withTyped {
			withTyped[i] = r.fillVariables(value)
		}
	case map[interface{}]interface{}:
		for key, value := range withTyped {
			withTyped[key] = r.fillVariables(value)
		}
	default:
		ui.Errorf("unsupported type: %T\n", with)
	}
	return with
}

func getValueByPath(m map[string]interface{}, key string) (interface{}, error) {
	keys := strings.Split(key, ".")
	var ok bool
	for i, k := range keys {
		if i == len(keys)-1 {
			if val, exists := m[k]; exists {
				return val, nil
			}
			return nil, fmt.Errorf("key %s not found", key)
		}
		if m, ok = m[k].(map[string]interface{}); !ok {
			return nil, fmt.Errorf("path %s is not a map", strings.Join(keys[:i+1], "."))
		}
	}
	return nil, fmt.Errorf("key %s not found", key)
}
