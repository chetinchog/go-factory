package factory

// IOption interface
type IOption interface {
	newOption() IOption
	Find() string
}
