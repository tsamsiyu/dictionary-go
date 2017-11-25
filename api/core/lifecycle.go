package core

import http "github.com/valyala/fasthttp"

// Middleware represents data structure that can be used in app lifecycle
type Middleware interface {
  BeforeHandle(ctx *http.RequestCtx)
  AfterHandle(ctx *http.RequestCtx)
}
