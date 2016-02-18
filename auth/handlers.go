package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func ValidationHandler(res http.ResponseWriter, req *http.Request) {
	// authHeader := req.Header.Get("Authorization")
	//
	// token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
	//       // Don't forget to validate the alg is what you expect:
	//       if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
	//           return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	//       }
	//       return myLookupKey(token.Header["kid"]), nil
	//   })
	//
	//   if err == nil && token.Valid {
	//       	res.Write([]byte("Token is valid"))
	//   } else {
	//       res.Write([]byte("Invalid Token"))
	//   }
}

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	token.Claims["foo"] = "bar"
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Sign and get the complete encoded token as a string
	if jwt, err := token.SignedString("everything_is_awesome!"); err == nil {
		res.Header().Set("Authorization", "Bearer "+jwt)
		res.Write([]byte("Token generated: " + jwt))
	} else {
		res.Write([]byte("Failed to generate token: " + err.Error()))
	}
}
