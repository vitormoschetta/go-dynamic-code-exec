package main

import "testing"

func TestDynamicExecute(t *testing.T) {
	// Arrange
	client := &ClientMock{}
	dynamic := NewDynamic(client)
	data := "Expense,2023-12-17,0"
	prompt := "Teste"

	// Act
	movement, err := dynamic.Execute(prompt, data)

	// Assert
	if err != nil {
		t.Errorf("Erro ao executar a função dinâmica: %s", err.Error())
	}
	if movement == nil {
		t.Error("Movimento não deve ser nulo")
	}
	if movement.Title != "Expense" {
		t.Errorf("Título do movimento deve ser 'Expense', mas foi '%s'", movement.Title)
	}
	if movement.Date != "2023-12-17" {
		t.Errorf("Data do movimento deve ser '2023-12-17', mas foi '%s'", movement.Date)
	}
	if movement.Value != 0 {
		t.Errorf("Valor do movimento deve ser 0, mas foi %f", movement.Value)
	}
}
