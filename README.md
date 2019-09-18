# jwt
jsonwebtoken-middleware for goa.

[![Build Status](https://travis-ci.org/goa-go/jwt.svg?branch=master)](https://travis-ci.org/goa-go/jwt)
[![Codecov](https://codecov.io/gh/goa-go/jwt/branch/master/graph/badge.svg)](https://codecov.io/github/goa-go/jwt?branch=master)
[![Go Doc](https://godoc.org/github.com/goa-go/jwt?status.svg)](http://godoc.org/github.com/goa-go/jwt)
[![Go Report](https://goreportcard.com/badge/github.com/goa-go/jwt)](https://goreportcard.com/report/github.com/goa-go/jwt)

## Installation

```bash
$ go get -u github.com/goa-go/jwt
```

## Example
```go
import (
  "github.com/goa-go/goa"
  "github.com/goa-go/jwt"
)

func main(){
  app = goa.New()
  app.Use(jwt.New(jwt.Options{
    Secret: "example-secret",
  }))

  ...
}
```

## Options

Field | Type | Reqired | Description
-|-|-|-
Secret | interface{} | true | jwt secret |
Unless | []string | false | unless paths |
GetToken | func(*goa.Context) string | false | custom getToken function |
Verify | func(string, interface{}) bool | false | custom verify function |

## License

[MIT](https://github.com/goa-go/goa/blob/master/LICENSE)
