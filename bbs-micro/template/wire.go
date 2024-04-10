package main

import "github.com/google/wire"

var thirdPartySet = wire.NewSet()

var interactiveSvcProvider = wire.NewSet()

func InitApp() *App {
	wire.Build()
	return new(App)
}
