# This is my Go learning repository, where I'm going through the typical use cases I have for a new language

- REST Server
- REST Client
- DB Access (Postgres)
- Redis Access
- RabbitMQ Access

# Setup Go on the machine
https://golang.org/doc/install?download=go1.14.2.windows-amd64.msi

# HTTP REST Server and Client

## Install the required package
```
go get -u github.com/gorilla/mux
```

## Build
```
cd ./grey-goose-go-srv
go build

cd ..

cd ./grey-goose-go-cli
go build
```
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
```
go get github.com/lib/pq
```

## Build and run
```
cd cli-sql-db-pg
go build
cli-sql-db-pg.exe
```

## Reference
PGSQL Connection : https://godoc.org/github.com/lib/pq
PGSQL Commands : https://www.calhoun.io/using-postgresql-with-go/

# Connecting to Redis

Redis is my preferred key value store. I use it as soon as I need a cache and/or events. This particular connector is maybe not the "cutest" but it was recommended as an official one on the Redis website.

## Run my Redis docker

```
docker run --name redis-1 -p 6379:6379 -d redis
```

## Install required package
```
go get github.com/gomodule/redigo/redis
```

## Build and run
```
cd cli-redis
go build
cli-redis.exe
```
## Reference 
Connector Redis: https://godoc.org/github.com/gomodule/redigo

# Connecting to Rabbit mq

## Launch Rabbitmq in docker

```
docker run -d --hostname rabbitmq --name rabbit-1 -p 5672:5672 rabbitmq:3.7.18 
```

## Install required package
```
go get github.com/streadway/amqp
```
## Producer and Consumer on the queue Hello 

```
go run send.go
go run receive.go
```

## Reference
https://www.rabbitmq.com/tutorials/tutorial-one-go.html
