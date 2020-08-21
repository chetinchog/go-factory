package factory

import (
	"factory/interfaces"
	"fmt"
	"sync"
)

var (
	f    *factory
	once sync.Once

	listOptions []interfaces.IOption
)

type factory struct{}

func setFactory() {
	once.Do(func() {
		f = &factory{}
	})
}

// AddOption allows an IOption to add itself
func AddOption(option interfaces.IOption) {
	listOptions = append(listOptions, option)
}

// GetFactory returns factory singleton
func GetFactory() interfaces.IFactory {
	setFactory()
	return f
}

// GetOption implements IFactory method
// Returns storage depending on the parameter
func (f *factory) GetOption(params ...interface{}) (interfaces.IOption, error) {
	for _, option := range listOptions {
		if option.EvalOption(params...) {
			return option, nil
		}
	}
	return nil, fmt.Errorf("factory.GetFactory: Option %s not found", params)
}
