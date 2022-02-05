package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publicacoes: representa um repositório de publicacoes
type Publicacoes struct {
	db *sql.DB
	// Aqui ficarão os métodos que farão a integração com as tabelas de BD
}

// NovoRepositorioDePublicacoes: cria um repositório de Publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar: insere uma publicação no BD
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutotID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}
