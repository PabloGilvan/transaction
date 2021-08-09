package main

import (
	"github.com/PabloGilvan/transaction/internal/config/global"
	"github.com/PabloGilvan/transaction/internal/container"
)

func main() {
	global.ViperConfig()
	StartServer(container.Injector())
}
