package main

import (
	"context"

	_ "github.com/joho/godotenv"
	"github.com/pubgo/funk/log"
	"github.com/pubgo/lava-examples/internal/bootstrap"
)

func main() {
	log.SetEnableChecker(func(ctx context.Context, lvl log.Level, nameOrMessage string, fields log.Map) bool {
		//if nameOrMessage == "eval type value" {
		//	return false
		//}
		//
		//if nameOrMessage == "grpc-server" {
		//	return false
		//}

		return true
	})

	bootstrap.Main()
}
