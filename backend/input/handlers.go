package input

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type input struct {
	Product    uint64 `json:"product"`
	Material  uint64 `json:"material"`
	Quantity uint64 `json:"quantity"`
}

func CreateInput(c echo.Context) error {
    inp := new(input)

 	if err := c.Bind(inp); err != nil {
		return err
	}

    if err := AddInput(inp); err != nil {
    	return c.JSON(http.StatusInternalServerError, "Error while adding input relation!")
    }

	return c.JSON(http.StatusCreated, "Input relation sucessfully added!")
}

/*
func GetInputs(c echo.Context) error{
	raws, error := AllInputs()

	if error != nil{
		return c.JSON(http.StatusInternalServerError, "Error while geting raw materials!")
	}

	return c.JSON(http.StatusOK, raws)
}

func GetInput(c echo.Context) error{
	ID, error := strconv.ParseUint(c.Param("id"), 10, 64)
	if error != nil {
		return c.JSON(http.StatusInternalServerError, "Error while converting id to integer!")
	}

	raw, error := OneInput(ID)

	if error != nil {
	    return c.JSON(http.StatusInternalServerError, "Error while getting raw material!")	
	}

	return c.JSON(http.StatusOK, raw)
}
*/
func PutInput(c echo.Context) error{
    inp := new(input)

 	if err := c.Bind(inp); err != nil {
		return err
	}

    if err := UpdateInput(inp); err != nil {
    	return c.JSON(http.StatusInternalServerError, "Error while updating input relation!")
    }

    return c.JSON(http.StatusOK, "Input relation sucessfully updated!")
}

func RemoveInput(c echo.Context) error{
    inp := new(input)

 	if err := c.Bind(inp); err != nil {
		return err
	}

    if err := DeleteInput(inp); err != nil {
    	return c.JSON(http.StatusInternalServerError, "Error while removing input relation!")
    }

    return c.JSON(http.StatusOK, "Input relation sucessfully deleted!")
}
