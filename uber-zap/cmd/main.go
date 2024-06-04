package main

import uberzap "github.com/x14n/goExperimental/uberZap"

func main() {
	uberzap.ZapConfig()
	uberzap.Logger.Info("hello")
}
