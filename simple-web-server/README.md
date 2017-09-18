# Simple Web Server

The code creates a simple web server that responds to HTTP requests.

## Usage

Configure `ADDR` environment variable.

    export ADDR=localhost:3000

Start the server.

    cd path/to/server.go
    go run server.go

Test the server by sending a HTTP request.

    GET request:
        curl -i http://localhost:3000/hello

    POST request:
        curl -i -X POST http://localhost:3000/hello

## Reference

https://drstearns.github.io/tutorials/goweb/

https://drstearns.github.io/tutorials/cors/
