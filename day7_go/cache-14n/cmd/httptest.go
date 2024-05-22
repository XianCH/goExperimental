package main

//
// import (
// 	"fmt"
// 	"log"
// 	"net/http"
//
// 	cache14n "github.com/x14n/goExperimental/day7_go/cache-14n"
// )
//
// var db = map[string]string{
// 	"Tom":  "630",
// 	"Jack": "589",
// 	"Sam":  "567",
// }
//
// func main() {
// 	cache14n.NewGroup("scores", 2<<10, cache14n.GetterFunc(
// 		func(key string) ([]byte, error) {
// 			log.Println("[SlowDB] search key", key)
// 			if v, ok := db[key]; ok {
// 				return []byte(v), nil
// 			}
// 			return nil, fmt.Errorf("%s not exist", key)
// 		}))
//
// 	addr := "localhost:9999"
// 	peers := cache14n.NewHTTPPool(addr)
// 	log.Println("geecache is running at", addr)
// 	log.Fatal(http.ListenAndServe(addr, peers))
// }
