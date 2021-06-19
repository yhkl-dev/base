package main

import (
	"base/src/base"
	"base/src/classes"
	"base/src/middlewares"
)

func main() {
	base.Ignite().
		Beans(base.NewGormAdapter()).
		Attach(middlewares.NewUserMid()).
		Mount("v1", classes.NewUserClass(), classes.NewIndexClass()).
		Launch()
}
