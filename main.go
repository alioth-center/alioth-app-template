package main

import (
	_ "alioth-app-template/app/router"
	"github.com/alioth-center/infrastructure/exit"
)

func main() {
	exit.BlockedUntilTerminate()
}
