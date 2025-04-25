# URL Shortener - Go, Redis, and Docker

## Overview

**URL Shortener** is a fast and efficient URL shortening service developed in **Go**, utilizing **Redis** for fast storage and lookup. The project provides a simple API that allows users to shorten long URLs and retrieve them via a shortened version. The service is built using **Gin** for the web framework and **Redis** as a cache for storing the URL mappings. The project also uses **Docker** for easy deployment and scalability.

The goal is to offer a lightweight, high-performance URL shortening service, making it easy to create and use shortened URLs for sharing.

## Completed Features

### Core Functionality

#### URL Shortening
*  Users can input long URLs and generate shortened URLs using the API endpoint `POST /create-short-url`.
*  Each shortened URL is mapped to a long URL and stored in Redis.
*  Shortened URLs follow a simple structure, for example: `http://localhost:9808/{shortened_id}`.

#### URL Redirection
*  Users can visit shortened URLs and will be automatically redirected to the original long URL.
*  The redirection process is fast, leveraging Redis for quick lookup.
*  The redirection is handled via the `GET /{shortUrl}` endpoint.

#### Redis Integration
*  Redis is used for storing and retrieving mappings between short and long URLs.
*  The key for Redis is `short_url:{shortened_id}`, and the value is the original long URL.
*  Redis ensures low-latency access, even under heavy traffic.

#### API Endpoints
* `POST /create-short-url`: Accepts a long URL and returns a shortened URL.
* `GET /{shortened_id}`: Redirects the user to the original long URL.
* `GET /`: Displays all shortened URLs and checks the status of the service.

#### Docker Integration
*  The project is containerized using Docker, with a `docker-compose.yml` file provided to set up both the URL shortener service and a Redis container.
*  Redis stores URL mappings, and the application is designed to be scalable with minimal configuration.

### Technical Details

### Technology Stack

*  **Language:** Go
*  **Web Framework:** Gin
*  **Database:** Redis
*  **Containerization:** Docker
*  **Hashing Algorithm:** SHA-256 for generating unique short URLs, followed by Base58 encoding for the shortened URL format.

### Redis Integration

*  Redis stores mappings of shortened URLs to their original long URLs.
*  Shortened URL IDs are generated using a SHA-256 hash of the original URL concatenated with the user's ID. The hash is then encoded in Base58 for a shorter representation.
*  The short URL is stored in Redis with a key of `short_url:{shortened_id}` and the original long URL as the value.

### Design Principles

*  The project follows a simple and clean architecture with a clear separation of concerns between the URL shortening logic, Redis interactions, and the API handling.
*  All URL mappings are stored in Redis for quick access and retrieval.
*  Designed for high scalability and performance using Docker containers for easy deployment.

### Docker Setup

*  Docker is used to containerize the Go application and the Redis instance.
*  The `docker-compose.yml` file manages the Docker containers for both the URL shortener service and Redis.
*  To run the application, simply use the command `docker-compose up` after setting up Docker on your local machine.

### Error Handling

*  If the short URL is not found in Redis, the user receives a 404 status with an appropriate error message.
*  If any error occurs during URL creation or redirection, the service returns detailed error messages with HTTP status codes.

## Contributing

Contributions are welcome! If you would like to enhance the project or fix bugs, feel free to fork the repository and submit a pull request. For larger changes, please open an issue to discuss your ideas first.
