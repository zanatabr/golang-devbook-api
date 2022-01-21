package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Usuarios: representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
	// Aqui ficarão os métodos que farão a integração com as tabelas de BD
}

// NovoRepositorioDeusuarios: cria um repositório de usuários
func NovoRepositorioDeusuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar: insere um usuario no BD
func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
