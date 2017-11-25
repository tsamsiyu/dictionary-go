package config

import (
  "dict/api/core"
  "dict/api/middlewares"
)

var Middlewares = []core.Middleware {
  new(middlewares.Transport),
}
