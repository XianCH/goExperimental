package main

type Error interface {
	Caller() []CallerInfo
	Wraped() []error
	Code() int
	error

	private()
}

type CallerInfo struct {
	FuncName string
	FileName string
	FileLine int
}

func New(msg string) error

func NewWithCode(code int, msg string) error

func Wrap(err error, msg string) error

func WrapWithCode(Code int, err error, msg string) error

func FromJson(json string) (Error, error)
func ToJson(err error) string
