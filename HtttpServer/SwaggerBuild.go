package htttpserver

import (
	swaggers "github.com/nttlong/wx/swagger3"
)

type SwaggerBuild struct {
	server  *HtttpServer
	BaseUri string
	info    SwaggerInfo
	swagger *swaggers.Swagger
	err     error
}
