package cache_14n

// input key to pick node
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// using group to get cache
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
