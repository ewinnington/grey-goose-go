# Setup Go on the machine
https://golang.org/doc/install?download=go1.14.2.windows-amd64.msi

# Install the required package
go get -u github.com/gorilla/mux

# Build
cd ./grey-goose-go-srv
go build

cd ..

cd ./grey-goose-go-cli
go build

# Run 
The go-srv create a listening go server that answers to http GET and POST Requests on port 8080. 
The go-cli is a client that pushes a json object to the server with POST and then gets all objects back with a GET

```
./grey-goose-go-srv/grey-goose-go-srv.exe
./grey-goose-go-cli/grey-goose-go-cli.exe
```

# Open localhost and check it works
Chrome.exe http://localhost:8080


# reference
Server : https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da
Client : https://golang.org/pkg/net/http/
