package helper

import "log"

// ErrorPanic 函数用于检查错误并在出现错误时触发 panic
func ErrorPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}
