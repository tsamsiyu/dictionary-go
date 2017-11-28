package api

import (
	// "os"
	// http "github.com/valyala/fasthttp"
	// "dict/api/config"
	"dict/api/middlewares"
	"dict/api/initializers"
	"dict/core"
	"dict/core/contracts"
)

// Start is entry point of api application
func Start() {
	config := core.HttpAppConfig {
		Initializers: []contracts.Initializer {
			new(initializers.Logger),
		},
		Middlewares: []contracts.Middleware {
			new(middlewares.Transport),
		},
	}

	app := core.CreateHttpApp(config)
}
