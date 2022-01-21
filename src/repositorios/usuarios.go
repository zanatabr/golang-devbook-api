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
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}
