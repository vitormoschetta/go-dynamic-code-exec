package main

type ClientMock struct {
}

func (c *ClientMock) GenerateCode(prompt string) (string, error) {
	return `package main

	import (
		"encoding/json"
		"fmt"
		"os"
		"strings"
		"strconv"
	)

	type Result struct {
		Title string 
		Date  string
		Value float64
	}
	
	func transformCSVToMovement(csvData string) ([]byte, error) {
		// Simulação de lógica de transformação
		parts := strings.Split(csvData, ",")
		if len(parts) < 3 {
			return nil, fmt.Errorf("CSV inválido: esperado pelo menos 3 partes, mas obteve %d", len(parts))
		}
	
		value, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			return nil, err
		}
	
		result := Result{
			Title: parts[0],
			Date:  parts[1],
			Value: value,
		}
	
		// Serializa o resultado em JSON
		jsonResult, err := json.Marshal(result)
		if err != nil {
			return nil, err
		}
	
		return jsonResult, nil
	}
	
	func main() {
		// Simule dados CSV de entrada
		if len(os.Args) < 2 {
			fmt.Println("Forneca os dados CSV como argumento de linha de comando.")
			return
		}
	
		csvData := os.Args[2]
	
		// Chame a função gerada com os dados CSV
		jsonResult, err := transformCSVToMovement(csvData)
		if err != nil {
			fmt.Println("Erro ao transformar CSV:", err)
			return
		}
	
		// Imprima o resultado
		fmt.Println(string(jsonResult))
	}
	`, nil
}
