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

## Docker
This has been published as an docker image

Use it with your compose file.

```yml
# docker-compose.yml
version: "3.8"
service:
  # ...
  proxy:
    image: faagerholm/reverse-proxy-go
    ports:
      - 10000:10000 # modify this for your needs
    environment:
      - PROXY_URL=http:example.com
      - VERBOSE_DEBUG=false
    command: go run main.go
```
connect to your reverse proxy with `localhost:1000` <- modify this to match your environment