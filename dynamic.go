package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Dynamic struct {
	client IClient
}

func NewDynamic(client IClient) *Dynamic {
	return &Dynamic{
		client: client,
	}
}

// TODO: alterar o parametro data para um path do arquivo e seu formato
func (d *Dynamic) Execute(prompt, data string) (*Movement, error) {
	generatedCode, err := d.client.GenerateCode(prompt)
	if err != nil {
		return nil, err
	}

	fmt.Println("Código Gerado pelo GPT:")
	fmt.Println(generatedCode)

	result, err := executeDynamicFunction(generatedCode, data)
	if err != nil {
		return nil, err
	}

	fmt.Println("Resultado da Transformação:")
	fmt.Println(string(result))

	var movement Movement
	if err := json.Unmarshal(result, &movement); err != nil {
		log.Fatal(err)
	}

	return &movement, nil
}

// Função para executar a função dinâmica gerada
func executeDynamicFunction(code string, csvData string) ([]byte, error) {
	// Crie um arquivo temporário para o código gerado
	tmpFile, err := os.CreateTemp("", "generated_code*.go")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	// Escreva o código gerado no arquivo temporário
	if _, err := tmpFile.WriteString(code); err != nil {
		return nil, err
	}
	tmpFile.Close()

	// Compile e execute o código gerado
	cmd := exec.Command("go", "run", tmpFile.Name(), "-csv", csvData)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("erro ao executar o código gerado: %v\nSaída de erro detalhada: %s", err, output)
	}

	return output, nil
}
