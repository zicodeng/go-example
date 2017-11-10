# Message Queue

## Run Example Code

Start RabbitMQ Docker container.

    docker run -d -p 5672:5672 --name dev-rabbitmq --hostname dev-rabbitmq rabbitmq

Run Go program (MQ Consumer)

    go run main.go

Run Node.js program (MQ Producer)

    ./send.js
