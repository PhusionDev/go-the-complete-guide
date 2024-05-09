# Go Packages

## Built-in Packages

- `fmt` - provides formatting to output, and general I/O
- `math` - math functionality

## Defining your own packages

- packages must be in their own named folders
- functions are exported/available if they start with capital letter
- must be imported with full path name `example.com/mod/package`

## Third-party Packages

```bash
go get github.com/Pallinder/go-randomdata
go get -u github.com/gin-gonic/gin
go get github.com/mattn/go-sqlite3
go get -u golang.org/x/crypto
go get -u github.com/golang-jwt/jwt/v5
```
