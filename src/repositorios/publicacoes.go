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

// // Buscar: traz todos os usuários que atendem um filtro de nome ou nick
// func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {

// 	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //  %nomeOuNick%

// 	linhas, erro := repositorio.db.Query(
// 		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? OR nick LIKE ?",
// 		nomeOuNick, nomeOuNick,
// 	)

// 	if erro != nil {
// 		return nil, erro
// 	}

// 	defer linhas.Close()

// 	var usuarios []modelos.Usuario

// 	for linhas.Next() {
// 		var usuario modelos.Usuario

// 		if erro = linhas.Scan(
// 			&usuario.ID,
// 			&usuario.Nome,
// 			&usuario.Nick,
// 			&usuario.Email,
// 			&usuario.CriadoEm,
// 		); erro != nil {
// 			return nil, erro
// 		}

// 		usuarios = append(usuarios, usuario)
// 	}

// 	return usuarios, nil

// }

// // BuscarPorID: traz um usuário do BD
// func (repositorio Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
// 	linhas, erro := repositorio.db.Query(
// 		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?",
// 		ID,
// 	)
// 	if erro != nil {
// 		return modelos.Usuario{}, erro
// 	}
// 	defer linhas.Close()

// 	var usuario modelos.Usuario

// 	if linhas.Next() {
// 		if erro = linhas.Scan(
// 			&usuario.ID,
// 			&usuario.Nome,
// 			&usuario.Nick,
// 			&usuario.Email,
// 			&usuario.CriadoEm,
// 		); erro != nil {
// 			return modelos.Usuario{}, erro
// 		}
// 	}

// 	return usuario, nil
// }

// // Atualizar: altera as informações de um usuário no BD
// func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
// 	statement, erro := repositorio.db.Prepare(
// 		"UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?",
// 	)
// 	if erro != nil {
// 		return erro
// 	}
// 	defer statement.Close()

// 	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
// 		return erro
// 	}

// 	return nil
// }

// // Deletar: exclui as informações de um usuário no BD
// func (repositorio Usuarios) Deletar(ID uint64) error {
// 	statement, erro := repositorio.db.Prepare("DELETE FROM usuarios WHERE id = ?")
// 	if erro != nil {
// 		return erro
// 	}
// 	defer statement.Close()

// 	if _, erro = statement.Exec(ID); erro != nil {
// 		return erro
// 	}

// 	return nil

// }

// // BuscarPorEmail: busca um usuário por e-mail e retorna o seu id e senha com hash
// func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
// 	linha, erro := repositorio.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)
// 	if erro != nil {
// 		return modelos.Usuario{}, erro
// 	}
// 	defer linha.Close()

// 	var usuario modelos.Usuario

// 	if linha.Next() {
// 		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
// 			return modelos.Usuario{}, erro
// 		}
// 	}

// 	return usuario, nil

// }

// // Seguir: permite que um usuário siga outro
// func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {
// 	statement, erro := repositorio.db.Prepare(
// 		"insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)",
// 	)
// 	if erro != nil {
// 		return erro
// 	}
// 	defer statement.Close()

// 	if _, erro := statement.Exec(usuarioID, seguidorID); erro != nil {
// 		return erro
// 	}

// 	return nil

// }

// // PararDeSeguir: permite que um usuário deixe de seguir outro
// func (repositorio Usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
// 	statement, erro := repositorio.db.Prepare(
// 		"DELETE FROM seguidores WHERE usuario_id = ? AND seguidor_id = ?",
// 	)
// 	if erro != nil {
// 		return erro
// 	}
// 	defer statement.Close()

// 	if _, erro := statement.Exec(usuarioID, seguidorID); erro != nil {
// 		return erro
// 	}

// 	return nil

// }

// // BuscarSeguidores: traz todos os seguidores de um usuário
// func (repositorio Usuarios) BuscarSeguidores(usuarioId uint64) ([]modelos.Usuario, error) {

// 	linhas, erro := repositorio.db.Query(`
// 		SELECT u.id, u.nome, u.nick, u.email, u.criadoEm
// 		FROM usuarios u
// 		INNER JOIN seguidores s
// 		ON u.id = seguidor_id
// 		WHERE s.usuario_id =  ?`, usuarioId,
// 	)

// 	if erro != nil {
// 		return nil, erro
// 	}

// 	defer linhas.Close()

// 	var seguidores []modelos.Usuario

// 	for linhas.Next() {
// 		var usuario modelos.Usuario

// 		if erro = linhas.Scan(
// 			&usuario.ID,
// 			&usuario.Nome,
// 			&usuario.Nick,
// 			&usuario.Email,
// 			&usuario.CriadoEm,
// 		); erro != nil {
// 			return nil, erro
// 		}

// 		seguidores = append(seguidores, usuario)
// 	}

// 	return seguidores, nil

// }

// // BuscarSeguindo: traz todos os usuários que um usuário segue
// func (repositorio Usuarios) BuscarSeguindo(usuarioId uint64) ([]modelos.Usuario, error) {

// 	linhas, erro := repositorio.db.Query(`
// 		SELECT u.id, u.nome, u.nick, u.email, u.criadoEm
// 		FROM usuarios u
// 		INNER JOIN seguidores s
// 		ON u.id = s.usuario_id
// 		WHERE s.seguidor_id =  ?`, usuarioId,
// 	)

// 	if erro != nil {
// 		return nil, erro
// 	}

// 	defer linhas.Close()

// 	var usuários []modelos.Usuario

// 	for linhas.Next() {
// 		var usuario modelos.Usuario

// 		if erro = linhas.Scan(
// 			&usuario.ID,
// 			&usuario.Nome,
// 			&usuario.Nick,
// 			&usuario.Email,
// 			&usuario.CriadoEm,
// 		); erro != nil {
// 			return nil, erro
// 		}

// 		usuários = append(usuários, usuario)
// 	}

// 	return usuários, nil

// }

// // BuscarSenha: devolve a senha de um usuário pelo ID
// func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
// 	linha, erro := repositorio.db.Query("SELECT senha FROM usuarios WHERE id = ?", usuarioID)
// 	if erro != nil {
// 		return "", erro
// 	}
// 	defer linha.Close()

// 	var usuario modelos.Usuario

// 	if linha.Next() {
// 		if erro = linha.Scan(&usuario.Senha); erro != nil {
// 			return "", erro
// 		}
// 	}

// 	return usuario.Senha, nil

// }

// // Atualizar: altera a senha de um usuário
// func (repositorio Usuarios) AtualizarSenha(ID uint64, senha string) error {
// 	statement, erro := repositorio.db.Prepare(
// 		"UPDATE usuarios SET senha = ? WHERE id = ?",
// 	)
// 	if erro != nil {
// 		return erro
// 	}
// 	defer statement.Close()

// 	if _, erro = statement.Exec(senha, ID); erro != nil {
// 		return erro
// 	}

// 	return nil
// }
