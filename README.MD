# Setup Go on the machine
https://golang.org/doc/install?download=go1.14.2.windows-amd64.msi

# Install the required package
go get -u github.com/gorilla/mux

# Build
go build

# Run 
grey-goose-go.exe

# Open localhost and check it works
Chrome.exe http://localhost:8080


# reference
https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da