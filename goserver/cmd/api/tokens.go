package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/akouvach/assessments/models"

	"github.com/pascaldekloe/jwt"
	"golang.org/x/crypto/bcrypt"
)

var validUser = models.Usuario{
	UsuarioId: 1,
	Email:     "akouvach@hotmail.com",
	Pass:      "$2a$12$DdX42i/pKDYnIzyFML.iKuG0hM93k2Y0URI/zVa4NnD1PeTnd6LhK",
}

// pass= password

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"pass"`
}

func (app *application) Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	fmt.Println(r.Body)

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		app.errorJSON(w, errors.New("Creds - unauthorized"))
		return
	}

	fmt.Println("Creds:", creds)

	//Acá es donde debo ir a la base de datos a buscar el uduario
	hashedPassword := validUser.Pass

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil {
		app.errorJSON(w, errors.New("Compare - unauthorized"))
		return
	}

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(validUser.UsuarioId)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "mydomain.com"
	claims.Audiences = []string{"mydomain.com"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(app.config.jwt.secret))
	if err != nil {
		app.errorJSON(w, errors.New("error signing"))
		return
	}

	app.writeJSON(w, http.StatusOK, string(jwtBytes), "reponse")

}

// https://go.dev/play/p/s8KlqJIOWej

// package main

// import (
//     "crypto/hmac"
//     "crypto/sha256"
//     "encoding/hex"
//     "fmt"
// )

// func main() {

//     secret := "mysecret"
//     data := "data"
//     fmt.Printf("Secret: %s Data: %s\n", secret, data)

//     // Create a new HMAC by defining the hash type and the key (as byte array)
//     h := hmac.New(sha256.New, []byte(secret))

//     // Write Data to it
//     h.Write([]byte(data))

//     // Get result and encode as hexadecimal string
//     sha := hex.EncodeToString(h.Sum(nil))

//     fmt.Println("Result: " + sha)
// }

// https://go.dev/play/p/uKMMCzJWGsW

// package main

// import (
// 	"fmt"
// 	"golang.org/x/crypto/bcrypt"
// )

// func main() {
// 	password := "password"

// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

// 	fmt.Println(string(hashedPassword))
// }
