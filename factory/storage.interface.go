package factory

// IFactory must implement methods to use Factory
type IFactory interface {
	GetOption(...interface{}) (IOption, error)
}

// IOption must implement methods to new Options
type IOption interface {
	evalOption(...interface{}) bool
	Do()
}
