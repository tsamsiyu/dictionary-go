package middlewares

import (
  // "log"
  http "github.com/valyala/fasthttp"
)

// Transport works with request and response objects
type Transport struct {

}

func (this *Transport) BeforeHandle(ctx *http.RequestCtx) {
  // log.Println("Request started")
}

func (this *Transport) AfterHandle(ctx *http.RequestCtx) {
  // log.Println("Request stopped")
}
