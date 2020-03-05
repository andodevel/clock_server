# Clock Server

Server side of clock in/out  
[![CircleCI](https://circleci.com/gh/andodevel/clock_server.svg?style=svg)](https://circleci.com/gh/andodevel/clock_server) [![Coverage Status](https://coveralls.io/repos/github/andodevel/clock_server/badge.svg?branch=master)](https://coveralls.io/github/andodevel/clock_server?branch=master)  

## Requisites

1. Postgres  
2. `[DEV Only]` GNU Make  

## Run

### Production

Run OS-specific executable binary in `${baseDir}/bin` sub-folders  

### Development

Set env PORT to run app on desired port, otherwise *38080* is used  

```sh
export PORT=80
```  

In cloud environment, PORT is set by the provider.  

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
  
*Update:* Visual Code go language server works pretty well now. Should switch to it  