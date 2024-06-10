package main

import uberzap "github.com/x14n/goExperimental/uber-zap"

func main() {
	uberzap.ZapConfig()
	uberzap.Logger.Info("hello")
}
