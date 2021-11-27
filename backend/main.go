package main

import (
	"github.com/labstack/echo/v4"
	"github.com/Blackth01/a3-sistemas-distribuidos/raw_material"
)

func main() {
	e := echo.New()

	// Routes
	e.POST("/raw_material", raw_material.CreateRawMaterial)
	e.GET("/raw_material", raw_material.GetRawMaterials)
    e.GET("/raw_material/:id", raw_material.GetRawMaterial)
    e.PUT("/raw_material/:id", raw_material.PutRawMaterial)
    e.DELETE("/raw_material/:id", raw_material.RemoveRawMaterial)

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
