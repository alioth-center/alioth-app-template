package api

import (
	"alioth-app-template/app/entity"
	"alioth-app-template/app/service"

	"github.com/alioth-center/infrastructure/network/http"
)

var ExampleApi exampleApiImpl

type exampleApiImpl struct {
	service *service.ExampleService
}

func (impl exampleApiImpl) Ping() http.Chain[*entity.ExampleRequest, *entity.ExampleResponse] {
	return http.NewChain(impl.service.Ping)
}
