package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario: representa um Usuário que usa a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Preparar: chama os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar vazio")
	}
	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar vazio")
	}
	if usuario.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar vazio")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar vazia")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace((usuario.Nome))
	usuario.Nick = strings.TrimSpace((usuario.Nick))
	usuario.Email = strings.TrimSpace((usuario.Email))
}
