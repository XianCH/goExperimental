package helper

import (
	"encoding/json"
	"runtime"
)

// Error 实现了自定义错误接口
type customError struct {
	msg     string
	code    int
	callers []CallerInfo
	wrapped []error
}

type Error interface {
	Caller() []CallerInfo
	Wrapped() []error
	Code() int
	error

	private()
}

type CallerInfo struct {
	FuncName string
	FileName string
	FileLine int
}

func (e *customError) Error() string {
	return e.msg
}

func (e *customError) Caller() []CallerInfo {
	return e.callers
}

func (e *customError) Wrapped() []error {
	return e.wrapped
}

func (e *customError) Code() int {
	return e.code
}

func (e *customError) private() {}

func captureCallers() []CallerInfo {
	pc := make([]uintptr, 10)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	var callers []CallerInfo

	for {
		frame, more := frames.Next()
		if frame.Function == "runtime.goexit" || frame.Function == "runtime.main" {
			break
		}
		callers = append(callers, CallerInfo{
			FuncName: frame.Function,
			FileName: frame.File,
			FileLine: frame.Line,
		})
		if !more {
			break
		}
	}
	return callers
}

func New(msg string) error {
	return &customError{
		msg:     msg,
		callers: captureCallers(),
	}
}

func NewWithCode(code int, msg string) error {
	return &customError{
		msg:     msg,
		code:    code,
		callers: captureCallers(),
	}
}

func Wrap(err error, msg string) error {
	return &customError{
		msg:     msg,
		callers: captureCallers(),
		wrapped: []error{err},
	}
}

func WrapWithCode(code int, err error, msg string) error {
	return &customError{
		msg:     msg,
		code:    code,
		callers: captureCallers(),
		wrapped: []error{err},
	}
}

func FromJson(jsonStr string) (Error, error) {
	var e customError
	err := json.Unmarshal([]byte(jsonStr), &e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func ToJson(err error) string {
	e, ok := err.(*customError)
	if !ok {
		return ""
	}
	jsonStr, _ := json.Marshal(e)
	return string(jsonStr)
}
