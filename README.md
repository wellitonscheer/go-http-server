# Server send message WhatsApp

For now its just using a test message from the api documentation

- create and set the .env variables
- run `go run main.go`
- make a call `curl -X POST http://localhost:8080/send/ass -d '{"to":"5555996583242"}' -H "Content-Type: application/json"`
