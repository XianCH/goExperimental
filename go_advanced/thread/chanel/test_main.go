package main

var done = make(chan bool)

var msg string

func aGoroutine() {
	msg = "hello,world"
	<-done
}

func testFunc() {
	go aGoroutine()
	done <- true
	println(msg)
}
