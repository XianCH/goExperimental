package client

import "flag"

var (
	address = flag.String("address", "localhost:50051", "address of the server")
	name    = flag.String("name", defaultName, "name to greet")
)
