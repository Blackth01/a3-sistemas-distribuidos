package raw_material

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type raw_material struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Inventory uint64 `json:"inventory"`
}

func CreateRawMaterial(c echo.Context) error {
    raw := new(raw_material)

 	if err := c.Bind(raw); err != nil {
		return err
	}

    if err := AddRawMaterial(raw); err != nil {
    	return c.JSON(http.StatusInternalServerError, "Error while adding raw material!")
    }

	return c.JSON(http.StatusCreated, "Raw material sucessfully added!")
}


func GetRawMaterials(c echo.Context) error{
	raws, error := AllRawMaterials()

	if error != nil{
		return c.JSON(http.StatusInternalServerError, "Error while geting raw materials!")
	}

	return c.JSON(http.StatusOK, raws)
}

func GetRawMaterial(c echo.Context) error{
	ID, error := strconv.ParseUint(c.Param("id"), 10, 64)
	if error != nil {
		return c.JSON(http.StatusInternalServerError, "Error while converting id to integer!")
	}

	raw, error := OneRawMaterial(ID)

	if error != nil {
	    return c.JSON(http.StatusInternalServerError, "Error while getting raw material!")	
	}

	return c.JSON(http.StatusOK, raw)
}
