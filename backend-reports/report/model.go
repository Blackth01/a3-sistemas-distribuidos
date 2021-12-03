package report

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/Blackth01/a3-sistemas-distribuidos/backend-reports/database"
)

type Report struct {
	product []struct {
		ID       uint64
		Name     string
		Quantity uint64
	}
	rest []MaterialRest
}

type MaterialRest struct {
	RawMaterialId uint64
	Name          string
	Quantity      uint64
}

type raw_material struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Inventory uint64 `json:"inventory"`
}

type productWithInput struct {
	ID         uint64  `json:"id"`
	Name       string  `json:"name"`
	Value      float64 `json:"value"`
	InputsId   string  `json:"inputs_id"`
	InputsName string  `json:"inputs_name"`
	InputsQty  string  `json:"inputs_qty"`
}

func GetFullReport() (*Report, error) {
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database!")
		return nil, error
	}
	defer db.Close()

	products, err := GetAllProductsWithInputs(db)
	if err != nil {
		fmt.Println("Error while getting products: ", err)
		return nil, err
	}

	rawMaterials, err := GetAllRawMaterials(db)
	if err != nil {
		fmt.Println("Error while getting products: ", err)
		return nil, err
	}

	fmt.Println("Products: ", products)
	fmt.Println("Raw Materials: ", rawMaterials)

	report := &Report{}
	var rest []*MaterialRest = []*MaterialRest{}
	for _, product := range products {
		// Just do the magic
		inputsIdsList := []uint64{}
		for _, id := range strings.Split(product.InputsId, ",") {
			parsedId, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				fmt.Println("Error converting ID")
				continue
			}
			inputsIdsList = append(inputsIdsList, parsedId)
		}

		inputsNames := []string{}
		for _, name := range strings.Split(product.InputsName, ",") {
			inputsNames = append(inputsNames, name)
		}

		inputsQtyList := []uint64{}
		for _, qty := range strings.Split(product.InputsQty, ",") {
			parsedQty, err := strconv.ParseUint(qty, 10, 64)
			if err != nil {
				fmt.Println("Error converting qty")
				continue
			}
			inputsQtyList = append(inputsQtyList, parsedQty)
		}

		quantity := 0
		// Precisa fazer isso contar a quantidade de peças pra fazer, então
		// faz rawMaterial % quantidade necessária pra peça
		// Depois pra conseguir a quantidade de peças que vai fazer
		// só pegar o valor do resto (rawMaterial % quantidade) e diminuir o total de rawMaterial
		// Aí divide pelo valor necessário pra fazer a peça, pronto, tem seu relatório
		// E fazer com que o for quebre quanto não tiver mais obra prima
		for _, material := range rawMaterials {
			for index := range inputsIdsList {
				if material.Inventory >= inputsQtyList[index] {
					for i, restMaterial := range rest {
						if restMaterial.RawMaterialId == inputsIdsList[index] {
							rest[i] = &MaterialRest{
								RawMaterialId: material.ID,
								Name:          material.Name,
								Quantity:      uint64(quantity),
							}
						}
					}
				}
			}
		}
	}

	return &Report{}, nil
}

func GetAllProductsWithInputs(db *sql.DB) ([]productWithInput, error) {
	lines, error := db.Query("select p.id, p.nome, p.valor, GROUP_CONCAT(i.id_materia_prima) materias_prima, GROUP_CONCAT(mp.nome) materia_prima_nome, GROUP_CONCAT(i.quantidade) quantidades from produto p left join insumo i on (p.id = i.id_produto) left join materia_prima mp on (i.id_materia_prima = mp.id) group by p.id order by p.valor desc")
	if error != nil {
		fmt.Println("Error while getting product: ", error)
		return nil, error
	}
	defer lines.Close()

	var products []productWithInput

	for lines.Next() {
		var prod productWithInput

		if error := lines.Scan(&prod.ID, &prod.Name, &prod.Value, &prod.InputsId, &prod.InputsName, &prod.InputsQty); error != nil {
			fmt.Println("Error while scanning product: ", error)
			return nil, error
		}

		products = append(products, prod)
	}

	return products, nil
}

func GetAllRawMaterials(db *sql.DB) ([]raw_material, error) {
	lines, error := db.Query("SELECT * FROM materia_prima")
	if error != nil {
		fmt.Println("Error while getting raw materials: ", error)
		return nil, error
	}
	defer lines.Close()

	var raw_materials []raw_material
	for lines.Next() {
		var raw raw_material

		if error := lines.Scan(&raw.ID, &raw.Name, &raw.Inventory); error != nil {
			fmt.Println("Error while scanning raw material: ", error)
			return nil, error
		}

		raw_materials = append(raw_materials, raw)
	}

	return raw_materials, nil
}
