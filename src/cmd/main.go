package main

import (
	"base/src/base"
	"base/src/classes"
	"base/src/middlewares"
)

func main() {
	base.Ignite().Attach(middlewares.NewUserMid()).
		Mount("v1", classes.NewUserClass(), classes.NewIndexClass()).
		Mount("v2", classes.NewIndexClass()).
		Launch()
}
