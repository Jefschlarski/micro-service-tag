package api

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func Get() ([]byte, error) {
	err := godotenv.Load("go.env")

	if err != nil {
		fmt.Println("Erro ao carregar o arquivo go.env:", err)
	}

	url := os.Getenv("API_URL")

	client := http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("Requisição GET realizada com sucesso!")
	return body, nil
}
