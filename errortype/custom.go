package errortype

type CustomError interface {
	Error() string
	Type() string
	Status() int
}
