//go:generate goagen bootstrap -d github.com/minhkhiemm/goa1/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/minhkhiemm/goa1/app"
)

func main() {
	// Create service
	service := goa.New("adder")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "operands" controller
	c := NewOperandsController(service)
	app.MountOperandsController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
