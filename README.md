# Voting System

## Diagrams

### Architecture Diagram

![arch](/docs/image.png)

## Tech Stack-**Backend**: Golang

- Go 1.22
- PostgreSQL
- RabbitMQ
- Nginx

## Setup

1. Install dependencies:

```bash
go mod tidy
```

2. Set up environment variables:
   Create a `.env` file from `.env.exampl` and replace the values with your secrets

```bash
cp .env.example .env
```

3. Set up docker container

```bash
docker-compose up -d
```

4. Start HTTP server

```bash
go run cmd/http/main.go
```

5. Start consumer

```bash
go run cmd/consumer/main.go
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
