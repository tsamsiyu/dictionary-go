package config

import (
  "dict/api/core"
  "dict/api/initializers"
)

var Initializers = []core.Initializer {
  new(initializers.LoggerInitializer),
}
