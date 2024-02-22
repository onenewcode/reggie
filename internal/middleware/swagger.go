package middleware

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

func InitSwagger(r *route.RouterGroup) {
	url := swagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition

	r.GET("/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
}
