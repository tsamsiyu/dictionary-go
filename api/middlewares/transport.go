package middlewares

import (
  "dict/core/contracts"
  // "dict/core"
  // http "github.com/valyala/fasthttp"
)

// Transport works with request and response objects
type Transport struct {

}

func (this *Transport) BeforeRequest(app contracts.App, ctx contracts.Request) {
}

func (this *Transport) AfterRequest(app contracts.App, ctx contracts.Request) {
}
