package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
)

func (a *Auth) saveTokenInClientCookie(token *oauth2.Token, w http.ResponseWriter, r *http.Request) error {

	// Genera un token JWT con la información del *oauth2.Token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		a.token: token,
	})

	// Firma el token con una clave secreta
	tokenString, err := jwtToken.SignedString([]byte(a.secret_key_token_signing))
	if err != nil {
		return fmt.Errorf("error al generar el token: %v", err)
	}

	// Duración del token
	// tokenDuration := time.Until(token.Expiry)
	tokenDuration := token.Expiry.Sub(time.Now())

	SetCookie(a.token, tokenString, a.domain, a.https, tokenDuration, w)

	return nil
}

func (a Auth) getTokenFromClientCookie(r *http.Request) (*oauth2.Token, error) {

	cookie, err := GetCookie(a.token, r)
	if err != nil {
		return nil, err
	}

	// Obtén el valor del token JWT de la cookie
	tokenString := cookie.Value

	// Analiza y verifica el token JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Devuelve la misma clave secreta utilizada para firmar el token
		return []byte(a.secret_key_token_signing), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error al analizar el token: %v", err)
	}

	// Verifica si el token es válido
	if token.Valid {
		// Accede a las reclamaciones del token según tus necesidades
		claims := token.Claims.(jwt.MapClaims)
		// Construye un *oauth2.Token a partir de las reclamaciones del token
		accessToken := claims["AccessToken"].(string)
		tokenType := claims["TokenType"].(string)
		refreshToken := claims["RefreshToken"].(string)
		expiry := claims["Expiry"].(string) // Asumiendo que la reclamación Expiry está en formato string

		expiryTime, err := time.Parse(time.RFC3339, expiry)
		if err != nil {
			return nil, fmt.Errorf("error al analizar la fecha de expiración: %v", err)
		}

		token := &oauth2.Token{
			AccessToken:  accessToken,
			TokenType:    tokenType,
			RefreshToken: refreshToken,
			Expiry:       expiryTime,
		}

		return token, nil
	}

	return nil, fmt.Errorf("token no válido")

}
