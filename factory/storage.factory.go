package factory

import (
	"fmt"
	"sync"
)

// EnumOptionNil default option for factory
const EnumOptionNil = "Not Implemented"

var (
	f    *factory
	once sync.Once

	listOptions []IOption
)

type factory struct{}

func setFactory() {
	once.Do(func() {
		f = &factory{}
	})
}

func addOption(option IOption) {
	listOptions = append(listOptions, option)
}

// GetFactory returns factory singleton
func GetFactory() IFactory {
	setFactory()
	return f
}

// GetOption implements IFactory method
// Returns storage depending on the parameter
func (f *factory) GetOption(params ...interface{}) (IOption, error) {
	for _, option := range listOptions {
		if option.evalOption(params...) {
			return option, nil
		}
	}
	return nil, fmt.Errorf("factory.GetFactory: Option %s not found", params)
}
