# Docker RPC chat (Go)

A lightweight, containerized RPC-based chat system implemented in Go.\
The server runs inside Docker and exposes a simple RPC interface, while
clients run locally to send and retrieve chat messages.

------------------------------------------------------------------------

## 📦 Docker Hub Image

**Image:** https://hub.docker.com/r/ahmedgebreel74/rpc-chat-server

------------------------------------------------------------------------

## 🧠 System Overview

### **Server**

-   Go-based RPC server
-   In‑memory, mutex‑protected message history
-   Exposes RPC methods:
    -   `SendMessage(ChatMessage)` --- append and broadcast message
    -   `GetHistory()` --- return current chat log
-   Runs on TCP port **1234**

### **Client**

-   Runs locally
-   Sends messages to the server and retrieves synchronized history
-   Displays chat log in a simple console interface

------------------------------------------------------------------------

## 🚀 Running the System

### **Start the Server (Docker)**

``` bash
docker pull ahmedgebreel74/rpc-chat-server:latest
docker run --rm -p 1234:1234 ahmedgebreel74/rpc-chat-server:latest
```

### **Run a Client**

``` bash
go run client.go
```

------------------------------------------------------------------------

## 🐳 Dockerfile (Server)

``` dockerfile
FROM golang:1.22-alpine
WORKDIR /app
COPY server.go .
RUN go build -o server server.go
ENV CHAT_PORT=1234
EXPOSE 1234
CMD ["./server"]
```

------------------------------------------------------------------------

## 🔧 Configuration

### Server

    CHAT_PORT   (default: 1234)

### Client

    CHAT_ADDR   (default: localhost:1234)

------------------------------------------------------------------------

## 🔒 Concurrency & Safety

-   Thread-safe shared message storage using `sync.Mutex`
-   Timestamped messages for consistent ordering

------------------------------------------------------------------------

## 📌 Limitations

-   No data persistence\
-   No authentication\
-   Clients pull history manually (no real-time push)

------------------------------------------------------------------------

## 🚀 Future Enhancements

-   Add persistence (Redis, SQLite)
-   Authentication & user tracking
-   Web UI with WebSockets
-   Custom RPC protocol / performance improvements

------------------------------------------------------------------------

## 👤 Author

**Ahmed Gebreel**\
Docker Hub: https://hub.docker.com/r/ahmedgebreel74\
GitHub: https://github.com/ahmedgebreel74
