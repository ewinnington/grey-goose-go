# This is my Go learning repository, where I'm going through the typical use cases I have for a new language

REST Server
REST Client
DB Access (Postgres)
_Redis Access_
_RabbitMQ Access_

# Setup Go on the machine
https://golang.org/doc/install?download=go1.14.2.windows-amd64.msi

# HTTP REST Server and Client

## Install the required package
go get -u github.com/gorilla/mux

## Build
cd ./grey-goose-go-srv
go build

cd ..

cd ./grey-goose-go-cli
go build

## Run 
The go-srv create a listening go server that answers to http GET and POST Requests on port 8080. 
The go-cli is a client that pushes a json object to the server with POST and then gets all objects back with a GET

```
./grey-goose-go-srv/grey-goose-go-srv.exe
./grey-goose-go-cli/grey-goose-go-cli.exe
```

## Open localhost and check it works
Chrome.exe http://localhost:8080


## reference
Server : https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da
Client : https://golang.org/pkg/net/http/

# Database access to PG SQL

The cli-sql-db-pg program accesses a pgsql db and runs a simple query against it the database. I will be using the pg db that I have in my blue-cobalt-whale repository (docker experiments) 

## Run my Postgres docker
```
docker run -d -p 5432:5432 --name pgdb-1 -e POSTGRES_USER=docker -e POSTGRES_PASSWORD=XdccDa85_JK -e POSTGRES_DB=docker postgres
```

## Install required package
go get github.com/lib/pq

## Reference
PGSQL Connection : https://godoc.org/github.com/lib/pq
PGSQL Commands : https://www.calhoun.io/using-postgresql-with-go/