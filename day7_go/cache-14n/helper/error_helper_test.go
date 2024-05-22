package helper

import ()

//
// func TestErrorHelper(t *testing.T) {
// 	// 创建一个缓冲区，用于捕获日志输出
// 	var logBuffer bytes.Buffer
// 	log.SetOutput(&logBuffer)
//
// 	// 设置模拟的调用栈信息
// 	callers := []error.CallerInfo{
// 		{FuncName: "main.main", FileName: "main.go", FileLine: 10},
// 		{FuncName: "main.foo", FileName: "main.go", FileLine: 20},
// 		{FuncName: "main.bar", FileName: "main.go", FileLine: 30},
// 	}
//
// 	// 创建一个自定义错误
// 	customErr := error.NewCustomError("test error", 123, callers, nil)
//
// 	// 调用 ErrorHelper
// 	error.ErrorHelper(customErr)
//
// 	// 检查输出
// 	expectedOutput := `Error: test error
// Callers:
//   main.go:10 main.main
//   main.go:20 main.foo
//   main.go:30 main.bar
// `
// 	if logBuffer.String() != expectedOutput {
// 		t.Errorf("expected:\n%s\nbut got:\n%s", expectedOutput, logBuffer.String())
// 	}
// }
