# Tasklister

Tasklister is a simple task management application built with Go and Docker. It allows users to create, read, update, and delete tasks.

## Features

- Add new tasks
- View all tasks
- Update existing tasks
- Delete tasks

## Prerequisites

- Go 1.16 or later
- Docker

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/tasklister.git
    cd tasklister
    ```

2. Build the Go application:
    ```sh
    go build -o tasklister
    ```

3. Run the application:
    ```sh
    ./tasklister
    ```

## Docker

To run the application using Docker:

1. Build the Docker image:
    ```sh
    docker build -t tasklister .
    ```

2. Run the Docker container:
    ```sh
    docker run -p 8080:8080 tasklister
    ```

## Usage

- Access the application at `http://localhost:8080`
- Use the web interface to manage your tasks

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
