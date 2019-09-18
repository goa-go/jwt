package jwt

import (
	"net/http"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/goa-go/goa"
	"github.com/stretchr/testify/assert"
)

func testDefaultGetToken(t *testing.T, authHeader string, token string) {
	c := &goa.Context{}
	c.Header = http.Header{
		"Authorization": []string{token},
	}
	assert.Equal(t, token, getToken(c))
}

func TestDefaultGetToken(t *testing.T) {
	testDefaultGetToken(t, "Bearer token", "token")
	testDefaultGetToken(t, "bearer token2", "token2")
	testDefaultGetToken(t, "BEARER token3", "token3")
}

func TestDefaultVerify(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
	})
	tokenString, _ := token.SignedString([]byte("test-secret"))
	assert.True(t, verify(tokenString, []byte("test-secret")))
}

func TestDefaultVerifyFailed(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"nbf": time.Now().AddDate(0, 0, 1).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("test-secret"))
	assert.False(t, verify(tokenString, []byte("test-secret")))
}

func TestNoSecret(t *testing.T) {
	assert.Panics(t, func() { New(Options{}) })
}

func TestUnless(t *testing.T) {
	c := &goa.Context{
		Path: "/login",
	}

	assert.NotPanics(t, func() {
		New(Options{
			Secret: "test-secret",
			Unless: []string{"/login"},
		})(c)
	})
}

func TestJWT(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
	})
	tokenString, _ := token.SignedString([]byte("test-secret"))

	c := &goa.Context{}
	c.Header = http.Header{
		"Authorization": []string{"Bearer " + tokenString},
	}

	assert.NotPanics(t, func() {
		New(Options{
			Secret: "test-secret",
		})(c)
	})
}

func TestJWTFailed(t *testing.T) {
	c := &goa.Context{}

	assert.Panics(t, func() {
		New(Options{
			Secret: "test-secret",
		})(c)
	})
}

func TestJWTCustomHandler(t *testing.T) {
	customGetToken := false
	customVerify := false
	c := &goa.Context{}
	New(Options{
		Secret: "test-secret",
		GetToken: func(c *goa.Context) string {
			customGetToken = true
			return ""
		},
		Verify: func(tokenString string, secret interface{}) bool {
			customVerify = true
			return true
		},
	})(c)

	assert.True(t, customGetToken)
	assert.True(t, customVerify)
}
