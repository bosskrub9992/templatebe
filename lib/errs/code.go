package errs

type Code int

const (
	Unknown Code = iota
	Bind
	Validate
)
