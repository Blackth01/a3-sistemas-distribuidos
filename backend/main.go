package main

import (
	"fmt"
	"log"

	"github.com/Blackth01/a3-sistemas-distribuidos/database"
)

func main() {
	db, erro := database.Connect()

	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	fmt.Println("A conexão está aberta!")

	linhas, erro := db.Query("SELECT nome FROM materia_prima LIMIT 1")

	if erro != nil {
		log.Fatal(erro)
	}

	defer linhas.Close()

	var materiaPrimas []struct {
		Name string `json:"nome"`
	}

	for linhas.Next() {
		var cdd struct {
			Name string `json:"nome"`
		}
		if err := linhas.Scan(&cdd.Name); err != nil {
			return
		}
		materiaPrimas = append(materiaPrimas, cdd)
	}

	fmt.Println("FINISHED")
	fmt.Println(materiaPrimas)
}
