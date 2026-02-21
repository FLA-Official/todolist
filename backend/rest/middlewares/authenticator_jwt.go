package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"net/http"
	"strings"
	"todolist/utils"
)

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//parse JWT
		//parse header and payload claims
		//hmac-sha-256 algorithm -> hash (header, payload, secret)
		//parse signature part from the JWT
		//if the signature and hash is same => forward to create products
		//otherwise 401 status code with Unauthorized

		//getting Bearer {{jwt_secret}} from Authorization Header
		header := r.Header.Get("Authorization")

		// checking if header is empty or not. If it's empty then returning unauthorized status
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		//spliting "Bearer {{jwt_secret}}" to headerArr[0] = Bearer and headerArr[1] = {{jwt_secret}}
		headerArr := strings.Split(header, " ")

		if len(headerArr) != 2 {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]

		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		byteArrSecret := []byte(m.cnf.JWTSecretKey)
		byteArrMessage := []byte(message)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)

		hash := h.Sum(nil)
		newSignature := utils.Base64UrlEncode(hash)

		if newSignature != signature {
			http.Error(w, "Unauthorized, Hacker Detected", http.StatusUnauthorized)
			return
		}

		// JWT is valid, pass to next handler
		next.ServeHTTP(w, r)
	})
}
