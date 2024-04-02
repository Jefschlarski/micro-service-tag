package data

import (
	"encoding/json"
	"fmt"
)

type Slip struct {
	ID     int    `json:"id"`
	Advice string `json:"advice"`
}

type SlipResponse struct {
	Slip Slip `json:"slip"`
}

func Save(data []byte) {
	var slip SlipResponse
	err := json.Unmarshal([]byte(data), &slip)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}
	fmt.Println("Salvando os dados:", slip)
}
