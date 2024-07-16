package workflow

type Spec struct {
	Workflows []Workflow `yaml:"workflows"`
}

type Workflow struct {
	Name    string   `yaml:"name"`
	ID      string   `yaml:"id"`
	Inputs  []Input  `yaml:"inputs"`
	Outputs []Output `yaml:"outputs"`
	Steps   []Step   `yaml:"steps"`
}

type Input struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default"`
	Description string `yaml:"description"`
}

type Output struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type Step struct {
	Name   string                 `yaml:"name"`
	ID     string                 `yaml:"id"`
	Uses   string                 `yaml:"uses"`
	With   map[string]interface{} `yaml:"with"`
	Output map[string]interface{} `yaml:"output"`
}
