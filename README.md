# Event Service

Event Service is a GoLang-based REST API designed to facilitate the creation, modification, and deletion of events for authorized users. The authentication mechanism is implemented using JWT tokens. The application utilizes PostgreSQL as its database backend and provides Docker and Docker Compose files for easy deployment.

## Features

- **User Authentication**: Secure authentication mechanism using JWT tokens.
- **Event Management**: Allows users to create, update, and delete events.
- **Event Registration**: Users can register and unregister for events.
- **Scalable Database**: Utilizes PostgreSQL as the backend database for efficient data storage.
- **Easy Deployment**: Docker and Docker Compose files provided for seamless deployment.

## Prerequisites

Before running the application, ensure you have the following installed:

- GoLang
- Docker

## Getting Started

To get started with the Event Service, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/event-service.git
    ```

2. Navigate to the project directory:

    ```bash
    cd event-service
    ```

3. Download dependencies:

    ```bash
    go mod download
    ```

4. Update the necessary environment variables in `.env` file:
   
    ```plaintext
    # Database
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_username
    DB_PASSWORD=your_password
    DB_NAME=event_service_db

    # JWT
    JWT_SECRET=your_jwt_secret

    # Server
    SERVER_PORT=8080
    ```

5. Build the project to an executable file in the `build` directory:

    ```bash
    go build -o ./build/bookingservice ./cmd/eventsservice/main.go
    ```

6. Start the Docker containers:

    ```bash
    docker-compose up
    ```

7. The API should now be accessible at `http://localhost:your_port`.

## API Endpoints

- **POST /auth/signUp**: Register a new user.
- **POST /auth/login**: Authenticate and generate JWT token.
- **POST /events**: Create a new event.
- **GET /events**: Retrieve all events.
- **GET /events/{id}**: Retrieve a specific event by ID.
- **PUT /events/{id}**: Update a specific event.
- **DELETE /events/{id}**: Delete a specific event.
- **POST /events/{id}/register**: Register for a specific event.
- **DELETE /events/{id}/register**: Unregister from a specific event.

## Configuration

You can configure the application settings by modifying the `.env` file. This includes database connection details, JWT secret key, server port, and other parameters.

