package api

import (
	"alioth-app-template/app/service"
)

func init() {
	ExampleApi = exampleApiImpl{service: service.NewExampleService()}
}
