package main

import (
	j "jwt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/goa-go/goa"
)

func main() {
	app := goa.New()
	app.Use(j.New(j.Options{
		Secret: "aa",
		Unless: []string{"/l"},
		Verify: func(tokenString string, secret interface{}) bool {
			return true
		},
	}))
	app.Use(func(c *goa.Context) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iat": time.Now().Unix(),
			"exp": time.Now().Unix(),
		})
		ss, _ := token.SignedString([]byte("aa"))
		c.String(ss)
	})

	app.Listen(":3000")
}
