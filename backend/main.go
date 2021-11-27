package main

import (
	"github.com/labstack/echo/v4"
	"github.com/Blackth01/a3-sistemas-distribuidos/raw_material"
	"github.com/Blackth01/a3-sistemas-distribuidos/product"
)

func main() {
	e := echo.New()

	// Raw material routes
	e.POST("/raw_material", raw_material.CreateRawMaterial)
	e.GET("/raw_material", raw_material.GetRawMaterials)
    e.GET("/raw_material/:id", raw_material.GetRawMaterial)
    e.PUT("/raw_material/:id", raw_material.PutRawMaterial)
    e.DELETE("/raw_material/:id", raw_material.RemoveRawMaterial)

    // Product routes
	e.POST("/product", product.CreateProduct)
	e.GET("/product", product.GetProducts)
    e.GET("/product/:id", product.GetProduct)
    e.PUT("/product/:id", product.PutProduct)
    e.DELETE("/product/:id", product.RemoveProduct)

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
