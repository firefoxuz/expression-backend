### Supported Operators:
- Addition: `+`
- Subtraction: `-`
- Multiplication: `*`
- Division: `/`

### Supported Operand Range:
- Positive integers: 0-9

## Installation

Follow these steps to set up the Expression Calculator Service and Expression Agent:

### Expression Calculator Service

1. Clone the Expression Calculator Service project repository:

    ```bash
    git clone https://github.com/firefoxuz/expression-backend.git
    ```

2. Navigate to the project directory:

    ```bash
    cd expression-backend
    ```

3. Copy the example environment configuration file:

    ```bash
    cp .env.json.example .env.json
    ```

4. Start the Docker containers using Docker Compose:

    ```bash
    docker-compose up -d
    ```

5. Apply database migrations:

    ```bash
    docker run -v ./db/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://expression_user:expression_password@localhost:5432/expression_db?sslmode=disable up
    ```

### Expression Agent

6. Clone the Expression Agent project repository:

    ```bash
    git clone https://github.com/firefoxuz/expression-agent.git
    ```

7. Navigate to the project directory:

    ```bash
    cd expression-agent
    ```

8. Copy the example environment configuration file:

    ```bash
    cp .env.json.example .env.json
    ```

9. Start the Docker containers using Docker Compose:

    ```bash
    docker-compose up -d
    ```

Now, both the Expression Calculator Service and Expression Agent should be up and running. You can access the Expression Calculator Service at [http://127.0.0.1:8082/](http://127.0.0.1:8082/) and the Expression Agent at their respective API endpoints.

## Microservice Architecture

Below is the architecture of our microservices:

![Microservice Architecture](https://i.imgur.com/GPWQPvn.png)

## Web Interface Example

Below is an example of how to interact with our web:

![Web Interface Example](https://i.imgur.com/XDPcoyP.png)

![Web Interface Example](https://i.imgur.com/3Tr33S0.png)