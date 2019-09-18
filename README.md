# jwt
jsonwebtoken middleware for goa.

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
- | - | - | -
Secret | interface{} | true | jwt secret
Unless | []string | false | unless paths
GetToken | func(*goa.Context) string | false | custom getToken function
Verify | func(string, interface{}) bool | false | custom verify function

## License

[MIT](https://github.com/goa-go/goa/blob/master/LICENSE)
