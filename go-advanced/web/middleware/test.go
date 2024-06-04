package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// type Handler interface {
//     ServeHTTP(ResponseWriter, *Request)
// }
//
// type HandlerFunc func(ResponseWriter, *Request)
//
// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
//     f(w, r)
// }

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		log.Println(timeElapsed)
	})
}

func Test01() {
	http.Handle("/", timeMiddleware(http.HandlerFunc(hello)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
