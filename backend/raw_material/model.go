package raw_material

import (
	"fmt"

	"github.com/Blackth01/a3-sistemas-distribuidos/database"
)


func AddRawMaterial(raw *raw_material) error{
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database!")
		return error
	}
	defer db.Close()

	statement, error := db.Prepare("INSERT INTO materia_prima (nome, estoque) VALUES (?, ?)")
	if error != nil {
		fmt.Println("Error preparing statement!: ",error)
		return error
	}
	defer statement.Close()

    insert, error := statement.Exec(raw.Name, raw.Inventory)
	if error != nil {
		fmt.Println("Error while executing statement!: ",error)
		return error
	}

	_ = insert

	return nil
}

func AllRawMaterials() ([]raw_material, error){
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database!")
		return nil, error
	}
	defer db.Close()

	lines, error := db.Query("SELECT * FROM materia_prima")
	if error != nil {
		fmt.Println("Error while getting raw materials: ",error)
		return nil, error
	}
	defer lines.Close()

	var raw_materials []raw_material
	for lines.Next() {
		var raw raw_material

		if error := lines.Scan(&raw.ID, &raw.Name, &raw.Inventory); error != nil {
			fmt.Println("Error while scanning raw material: ",error)
			return nil, error
		}

		raw_materials = append(raw_materials, raw)
	}

	return raw_materials, nil
}

func OneRawMaterial(id uint64) (*raw_material, error){
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database!")
		return nil, error
	}
	defer db.Close()

	line, error := db.Query("SELECT * FROM materia_prima WHERE id = ?", id)
	if error != nil {
		fmt.Println("Error while getting raw material: ",error)
		return nil, error
	}
	defer line.Close()

	var raw raw_material
	if line.Next() {
		if error := line.Scan(&raw.ID, &raw.Name, &raw.Inventory); error != nil {
			fmt.Println("Error while scanning raw material: ",error)
			return nil, error
		}
	}

	return &raw, nil
}