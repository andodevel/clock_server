# clock_server readme

## Requisites

1. Postgres  
2. `[DEV Only]` GNU Make  

## Run

### Production

Run OS-specific executable binary in `${baseDir}/bin` sub-folders  

### Development

Create your own `${baseDir}/.env.local` file  with *PORT* parameter to start echo server on, otherwise *8080* is used  

```sh
PORT=3000
```  

Execute folllow command  

```sh
make start
```

### Working with VSCode

In case `go module` not work well with VSCode, we need workaround by:  

1. Disable VSCode settings `"go.useLanguageServer" = false`  
2. Keep your project under `$GOPATH`  
3. Execute `make ide` to update local `vendor` directory for every module change  
4. Add `vendor` directory to .gitignore  
