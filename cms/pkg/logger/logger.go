package logger

import (
	"log"
	"sync"

	"go.uber.org/zap"
)

var (
	LL   *zap.Logger
	L    *zap.SugaredLogger
	once sync.Once
)

func Init() {
	once.Do(func() {
		l, err := zap.NewDevelopment()
		if err != nil {
			log.Fatal(err)
		}
		LL = l
		L = l.Sugar()
	})
}
