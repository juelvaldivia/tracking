[![Deploy](https://github.com/juelvaldivia/tracking/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/juelvaldivia/tracking/actions/workflows/ci.yml)

# GPS Tracking in Go
Welcome to the GPS Tracking project in Go! This project aims to provide a robust and efficient solution for GPS device tracking.

Whether you're building a vehicle tracking system, a fitness application, or any app requiring real-time tracking, this Go project has got you covered.

## Installation

1. Clone the repository to your local machine:

   ```shell
   git clone git@github.com:juelvaldivia/tracking.git
   ```

2. Navigate to the application directory::
   ```shell
   cd tracking
   ```

# Configuration

The behavior of the application is determined by the `config.yml` file, where you can configure the following settings:

- Create the `config.yml` file you can base it on `config-example.yml` you can copy it like this:
    ```shell
    cp config-example.yml config.yml
    ```

- Api Configuration:

  ```yml
  api_port = 5678
  ```

- Database Configuration:

  ```yml
  database_driver = memory|sql
  ```

  if you choose a sql database, configure your environment accordingly
  ```yml
  sql_database:
    host: localhost
    port: 5432
    username: your_username
    password: your_password
    database: your_database
  ```

# Usage

- Run the application:
    ```bash
    go run main.go
    ```

- To stop the application, press `Ctrl+C` in the terminal.

# Development Usage

## Lint Usage
- Install [golangci-lint](https://golangci-lint.run/usage/install/) to use
```bash
golangci-lint run
```

## Tests
- Run test
 ```bash
  go test -v ./tests/...
 ```

# Acknowledgments

This application was developed by Luis Velazco, Eduardo Valdivia and Joel Valdivia.
