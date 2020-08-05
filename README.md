# go-module-package-testing
 
 - Following training from https://www.youtube.com/watch?v=Z1VhG7cf83M
 
## Packages added
- https://github.com/julienschmidt/httprouter
- go.uber.org/zap

## Run

- run `localhost:8080/hello/alim`
- run `localhost:8080` will log in console:
   `{"level":"info","ts":1596650175.2169445,"caller":"go-module-package-testing/main.go:14","msg":"Success"}
`

## Commands Run

```batch
go mod init github.com/aukgit/go-module-package-testing
go get github.com/julienschmidt/httprouter
go get go.uber.org/zap
go mod tidy
```


# Shortcuts

- [Ctrl + K] Commit
- [Alt + F12] Terminals
- [Alt + 1] Projects