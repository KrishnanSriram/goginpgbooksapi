## Golang API

## Why Golang for REST?
Golang has been around for sometime to build real-time applications for a long time. It's web presence has been limited but growing. The simplicity of the language and framework demands has made it a more popular language. Further frameworks like Gin make the adoption process more easier. Although not a measure, here's something that may catch your eyes to get started
https://medium.com/deno-the-complete-reference/go-gin-vs-springboot-webflux-hello-world-performance-comparison-e25fd8d85216

I know this is just a hello world program. But you can find a lot of stats on sizable development and comparison too. In the end choose what is right for you and for your organization. Knowing one other language will only make you better.

## Get started
Create a directory in your workspace, say "goginpg". Make sure you have GOROOT and GOPATH setup right, per Go documentation. Mod in Go allows you to build application with much needed isolation. Let's start with it. Browse into "goginpg" directory and execute the following command
```shell
go mod init goginpg
```
Open this directory in your favorite IDE. I prefer to to use Golang or IntelliJ. Create a directory structure for ease of operations.
- main.go - First program to execute
- handlers - Handle HTTP requests
- models - Models and DTOs can be defined
- initializers - All application wide initializations are done here

Above structure is simple, straight forward and self-explanatory. We can have more structures for authentication, services etc. But for now, let's keep it easy and build them as we go along 

## What we wish to accomplish
- Simple REST endpoints
- Dockerize API
- Use .env
- Logging
- Model validation
- Documentation - Swagger

## Setup process
These are the libraries we need to install to get started
```shell
    go get -u github.com/gin-gonic/gin
    go get github.com/go-playground/validator/v10
    go get github.com/joho/godotenv/cmd/godotenv
    go get github.com/lib/pq
        
```

## What next
- JWT for authentication
- Connect to EventBus for event driven development
- GraphQL support
- HTMX support
- ISTIO integration