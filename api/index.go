package api

import (
	"os"
	"log"
	http "github.com/valyala/fasthttp"
	"dict/api/config"
)

func Run() {
	http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), func (ctx *http.RequestCtx) {
		log.Println("Server started at port", os.Getenv("SERVER_PORT"))
		for _, m := range config.Middlewares {
			m.BeforeHandle(ctx)
		}
		Router().Handler(ctx)
		for _, m := range config.Middlewares {
			m.AfterHandle(ctx)
		}
	})
}
