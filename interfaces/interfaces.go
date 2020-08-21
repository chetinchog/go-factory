package interfaces

// IFactory must-implement methods to use Factory
type IFactory interface {
	// GetOption returns the corresponding option
	GetOption(...interface{}) (IOption, error)
}

// IOption must-implement methods to use Option
type IOption interface {
	// evalOption option selection logic
	EvalOption(...interface{}) bool

	// evalOption option function logic
	Do()
}
