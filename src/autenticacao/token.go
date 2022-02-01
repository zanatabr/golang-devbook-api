package autenticacao

import (
	"api/src/config"
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
