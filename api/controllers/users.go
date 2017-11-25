package controllers

import (
	"fmt"

	http "github.com/valyala/fasthttp"
)

func UsersList(ctx *http.RequestCtx) {
	fmt.Fprintf(ctx, "Hello!!!\n")
}
