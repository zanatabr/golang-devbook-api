package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go" // import com alias
)

// CriarToken: retorna um token assinadocom as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{} // Map que conterá as permissões no Token
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix() // expiração 6h
	permissoes["usuarioID"] = usuarioID

	// Assinatura digital do token com uma chave secreta (preparação)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	// Assinatura com chave secreta (chave ficará em var de ambiente)
	return token.SignedString([]byte(config.SecretKey))

}

// ValidarToken: verifica se o token passado na requisição é válido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	// A função jwt.Parse verifica todos os segmentos do token (header,
	// claims/permissões, assinatura)
	// O segundo parâmetro deve ser uma "função" que devolve a chave de
	// verificação utilizada para realizar o parse do token.
	// Mas não poderia ser apenas o SecretKey que já temos? Não, porque
	// de acordo com a documentação do JWT, antes de retornar a chave de
	// verificação, é necessário verificar também se o "método de
	// assinatura" é o esperado. Pois não faz sentido assinar um token
	// com um método, e fazer o parse usando outro método.
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// permissoes["usuarioId"] é do tipo "interface". Espera-se um string
		// Precisamos primeiro converter para string, para depois uint
		// Mas... o JWT armazena o uint64 como um float. Então precisaremos
		// converter o float para string, para depois converter para uint
		usuarioID, erro := strconv.ParseUint(
			fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return usuarioID, nil
	}
	return 0, errors.New("token inválido")
}

// Essa função apenas "extrai" o token do request
// Não faz a validação
func extrairToken(r *http.Request) string {

	// o Token é passado no Authorization do Request
	token := r.Header.Get("Authorization")

	// Formato: (Bearer) + (Token)
	// Bearer dskdssdksdldslkdskljds

	if len(strings.Split(token, " ")) == 2 {
		// devolve a 2a paravra
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	// O método que está sendo usado é de uma família específica?
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
