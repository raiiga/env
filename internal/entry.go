package internal

import (
	"errors"
	"os"
	"reflect"
	"strconv"
)

const (
	EntryTag       = "env"
	EntryColon     = ":"
	EntrySeparator = ","

	Name     = "name"
	Required = "required"
	Default  = "default"
)

type Entry struct {
	Name     string
	Default  string
	Required bool
}

func NewEntry(params map[string]string) *Entry {
	required, _ := strconv.ParseBool(params[Required])
	return &Entry{Name: params[Name], Default: params[Default], Required: required}
}

func (e *Entry) Assign(value reflect.Value) error {
	if e.Name == "" {
		return errors.New("env: property '" + Name + "' is required")
	}

	env := os.Getenv(e.Name)

	if env == "" && e.Default == "" && e.Required {
		return errors.New("env: variable '" + e.Name + "' is required")
	}

	if env == "" && e.Default != "" {
		env = e.Default
	}

	value.SetString(env) // TODO: multi-type support
	return nil
}
