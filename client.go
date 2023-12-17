package main

import "errors"

type IClient interface {
	GenerateCode(prompt string) (string, error)
}

type Client struct {
}

func (c *Client) GenerateCode(prompt string) (string, error) {
	// Substitua esta chamada simulada pelo código real para interagir com o GPT
	// Aqui, estou usando um prompt fixo, mas você deve adaptar isso para se comunicar com o GPT
	// através da API ou de outro método de integração.
	return "", errors.New("not implemented")
}
