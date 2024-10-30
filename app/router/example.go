package router

import (
	"alioth-app-template/app/api"
	"alioth-app-template/app/entity"
	"github.com/alioth-center/infrastructure/network/http"
)

var exampleRouter = http.NewRouter("/example")

var ExampleRouterGroup = []http.EndPointInterface{
	http.NewEndPointBuilder[*entity.ExampleRequest, *entity.ExampleResponse]().
		SetHandlerChain(api.ExampleApi.Ping()).
		SetAllowMethods(http.POST).
		SetRouter(exampleRouter.Group("/ping")).
		Build(),
}
