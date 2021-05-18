# Simple reverse proxy for your frontend

## SETUP

`cp example.env $USER.env`

config `$USER.env` port and proxy url for your server.

`source $USER.env`

`docker-compose up` 

You're good to go.

> Connect to your reverse proxy with `localhost:$PORT`.

## Local 

*required local copy of golang*  
*you'll need to source $USER*  

You can also run this locally, either by running the main.go or by compiling the main.go file and run the binary.

```bash
> go run main.go
# or
> go build -o reverse-proxy main.go
> ./reverse-proxy
```