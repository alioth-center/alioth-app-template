package service

import (
	"alioth-app-template/app/entity"
	"github.com/alioth-center/infrastructure/network/http"
)

type ExampleService struct{}

func NewExampleService() *ExampleService {
	return &ExampleService{}
}

func (srv *ExampleService) Ping(ctx http.Context[*entity.ExampleRequest, *entity.ExampleResponse]) {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetResponse(&entity.ExampleResponse{Hello: ctx.Request().Hello})
}
