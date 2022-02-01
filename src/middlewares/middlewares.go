package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Camada que fica entre a requisição e a resposta.
//
// Comumente utilizado quando temos alguma função ou recurso
// que deve ser aplicada a todas as rotas, sem a necessidade
// de aplicar isso "rota-por-rota".

// Autenticar: verifica se o usuário que fez a requisição está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validando...") // provisoriamente
		next(w, r)
	}
}

// Logger: escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}
