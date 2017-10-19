# Work with MYSQL in Go

## Installation and Run

Step 1: build a new Docker image based on mysql.

    docker run -t zicodeng/mysql-demo .

Step 2: run the Docker container based on the image we just created.

    ./run.sh

Step 3: connect to MYSQL CLI.

    ./connect.sh

Step 4: execute our GO program.

    go run main.go