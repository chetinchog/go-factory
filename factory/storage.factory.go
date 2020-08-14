package factory

import (
	"fmt"
	"sync"
)

// EnumOptionNil an option for factory
const EnumOptionNil = "Not Implemented"

var (
	f    *factory
	once sync.Once

	listOptions map[string]IOption
)

type factory struct{}

func setFactory() {
	once.Do(func() {
		f = &factory{}
	})
}

func addOption(key string, option IOption) {
	listOptions[key] = option
}

// IFactory methods access
type IFactory interface {
	GetOption(string) (IOption, error)
}

// GetFactory returns factory singleton
func GetFactory() IFactory {
	setFactory()
	return f
}

// GetOption implements IFactory interface
// Returns storage depending on the parameter
func (f *factory) GetOption(key string) (IOption, error) {
	if option := listOptions[key]; option != nil {
		return option, nil
	}
	return nil, fmt.Errorf("factory.GetFactory: Option '%s' not found", key)
}
