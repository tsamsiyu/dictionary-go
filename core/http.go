package core

import (
  http "github.com/valyala/fasthttp"
  "dict/core/contracts"
)

type HttpApp struct {
  config HttpAppConfig
}

func (app *HttpApp) Init() {
  for _, i := range(app.config.Initializers) {
    i.Prepare(app)
  }
  for _, i := range(app.config.Initializers) {
    i.Run(app)
  }
}

type HttpAppConfig struct {
  Middlewares []contracts.Middleware
  Initializers []contracts.Initializer
}

type HttpResponse struct {

}

func CreateHttpApp(config HttpAppConfig) *HttpApp {
  app := new(HttpApp)
  app.config = config
  app.Init()
  return app
}

func (app *HttpApp) Process(request contracts.Request) contracts.Response {
		for _, m := range app.config.Middlewares {
			m.BeforeRequest(app, request)
		}
    response := new(HttpResponse)
		// Router().Handler(ctx)
		for _, m := range app.config.Middlewares {
			m.AfterRequest(app, request)
		}
    return response
}

type Middleware interface {
  BeforeRequest(request *http.RequestCtx)
  AfterRequest(request *http.RequestCtx)
}
