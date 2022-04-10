package main

import (
	systemBundle "github.com/moadkey/gosvelte/system/bundle"
	"go.uber.org/fx"
	
	"github.com/moadkey/gosvelte/modules/welcome"
)

func main() {
	app := fx.New(
		systemBundle.Module,
		fx.Invoke(welcome.New),
	)
	app.Run()
}