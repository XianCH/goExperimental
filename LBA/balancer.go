package lba

type Balancer interface {
	NextServer() string
}
