package jwt

import (
	"errors"
	"net/http"

	"github.com/goa-go/goa"
	"github.com/goa-go/goa/utils"
)

// Options is used to declare options, use it by New(jwt.Options{...}).
type Options struct {
	Secret   interface{}
	Unless   []string
	GetToken func(*goa.Context) string
	Verify   func(string, interface{}) bool
}

type token struct {
	getToken func(*goa.Context) string
	verify   func(string, interface{}) bool
}

// New returns a jwt middleware for goa.
func New(ops Options) goa.Middleware {
	if ops.Secret == nil {
		panic(errors.New("require secret"))
	}

	if str, ok := ops.Secret.(string); ok {
		ops.Secret = utils.Str2Bytes(str)
	}
	s := ops.Secret

	token := token{}
	if ops.GetToken != nil {
		token.getToken = ops.GetToken
	} else {
		token.getToken = getToken
	}

	if ops.Verify != nil {
		token.verify = ops.Verify
	} else {
		token.verify = verify
	}

	return func(c *goa.Context) {
		if !include(ops.Unless, c.Path) {
			tokenString := token.getToken(c)
			if !token.verify(tokenString, s) {
				c.Error(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			}
		}
		c.Next()
	}
}

func include(strArray []string, str string) bool {
	for _, s := range strArray {
		if s == str {
			return true
		}
	}

	return false
}
