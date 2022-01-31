package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota: representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar: Insere todas as Rotas no Router, e retorna o Router com todas as rotas configuradas
//             a) Recebe um Router sem nenhuma Rota
//             b) Configura todas as Rotas
//             c) Devolve um Router preparado
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin) // faz o append da rota de login

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
