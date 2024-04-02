package main

import (
	"fmt"
	"time"

	Api "micro.service.tag/api"
	Service "micro.service.tag/data"
)

func main() {

	interval := 30 * time.Second

	ticker := time.NewTicker(interval)

	for range ticker.C {
		if res, err := Api.Get(); err != nil {
			fmt.Printf("Erro ao fazer a requisição: %v\n", err)
		} else {
			fmt.Println("Corpo da resposta:", string(res))
			Service.Save(res)
		}
	}
}
