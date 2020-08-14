package factory

// IOption interface
type IOption interface {
	evalOption(...interface{}) bool
	getOption() IOption
	Do()
}
