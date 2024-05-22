package cache_14n

import (
	"fmt"
	"testing"

	"github.com/x14n/goExperimental/day7_go/cache-14n/lru"
)

func TestCacheGet(t *testing.T) {
	cache := &cache{
		lru: lru.New(10, nil),
	}

	cache.add("key1", ByteView{b: []byte("key4")})
	fmt.Println(cache.get("key1"))
	// expect := "123"
	// if v, ok := cache.get("key1"); ok {
	// 	t.Errorf("can't get key1")
	// } else if .(string) != expect {
	// }
}
