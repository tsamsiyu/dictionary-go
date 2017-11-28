package core

import (
  log "github.com/sirupsen/logrus"
  http "github.com/valyala/fasthttp"
)

// API represents api object structure
type API struct {
  Logger log.Logger
}

// Middleware represents objects that is called before and after http requests
type Middleware interface {
  BeforeHandle(ctx *http.RequestCtx)
  AfterHandle(ctx *http.RequestCtx)
}

// Initializer represents objects that is called before application is started
type Initializer interface {
  InitializeApplication(api *API)
}
