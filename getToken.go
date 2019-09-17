package jwt

import (
	"strings"

	"github.com/goa-go/goa"
)

func getToken(c *goa.Context) string {
	token := c.Header.Get("Authorization")

	if len(token) > 6 && strings.ToUpper(token[0:7]) == "BEARER " {
		return token[7:]
	}
	return token
}
