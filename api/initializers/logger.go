package initializers

import (
  "dict/core/contracts"
  // "github.com/sirupsen/logrus"
)

type Logger struct {

}

func (logger *Logger) Prepare(app contracts.App) {

}

func (logger *Logger) Run(app contracts.App) {
  // api.Logger = *logrus.New()
}
