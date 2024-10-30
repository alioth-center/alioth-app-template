package router

import (
	"alioth-app-template/app/global"

	"github.com/alioth-center/infrastructure/network/http"
)

func serveBackend() {
	engine := http.NewEngine(global.Config.HttpEngineConfig.ServeURL)

	engine.AddEndPoints(ExampleRouterGroup...)

	engine.ServeAsync(global.Config.HttpEngineConfig.ServeAddr, make(chan struct{}, 1))
}

func init() {
	serveBackend()
}
