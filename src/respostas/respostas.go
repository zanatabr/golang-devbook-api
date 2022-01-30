package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON: retorna uma reposta em JSon para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// tratamento feito para os casos em que a resposta for NoContent
	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}

}

// Erro: retorna um erro no formato JSon
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
