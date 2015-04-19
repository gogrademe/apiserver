//go:generate go run gen/generate.go

package envconfig

import (
	"fmt"
	"os"
	"strings"
)

var environment = make(map[string]*ConfigVar)

// ConfigVar represents a value from the environment.
type ConfigVar struct {
	Name        string
	Description string
	Value       Value  // value as set
	Default     string // default value (as text); for description message
}

// NewVar retrieves a variable from the environment that is of type Value.
func NewVar(value Value, name string, description string) *ConfigVar {
	envVar := &ConfigVar{
		Name:        strings.ToUpper(name),
		Description: description,
		Value:       value,
		Default:     value.String(),
	}
	_, defined := environment[name]
	if defined {
		panic("env: " + name + " already defined.")
	}

	actual := os.Getenv(envVar.Name)
	if actual != "" {
		envVar.Value.Set(actual)
	}

	environment[name] = envVar

	return envVar
}

// Var retrieves a ConfigVar by name from the ConfigVar map.
func Var(name string) *ConfigVar {
	if v, ok := environment[name]; ok {
		return v
	}
	return nil
}

// Vars retrieve all ConfigVars from the ConfigVar map.
func Vars() []*ConfigVar {
	vars := make([]*ConfigVar, len(environment))
	for _, v := range environment {
		vars = append(vars, v)
	}
	return vars
}

// PrintDefaults prints, to stderr, the default values of all defined ConfigVars.
func PrintDefaults() {
	for _, v := range environment {
		fmt.Fprintf(os.Stderr, "%s=%q: %s\n", v.Name, v.Default, v.Description)
	}
}
