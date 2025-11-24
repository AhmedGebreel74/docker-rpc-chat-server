package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"sync"
	"time"
)

type ChatMessage struct {
	Author    string
	Text      string
	Timestamp time.Time
}

type ChatServer struct {
	mu      sync.Mutex
	history []ChatMessage
}

func (s *ChatServer) SendMessage(msg ChatMessage, reply *[]ChatMessage) error {
	if msg.Author == "" {
		return errors.New("author is required")
	}
	if msg.Text == "" {
		return errors.New("text is required")
	}
	if msg.Timestamp.IsZero() {
		msg.Timestamp = time.Now()
	}

	s.mu.Lock()
	s.history = append(s.history, msg)
	fmt.Printf("[%s] %s: %s\n", msg.Timestamp.Format("15:04:05"), msg.Author, msg.Text)
	h := make([]ChatMessage, len(s.history))
	copy(h, s.history)
	s.mu.Unlock()

	*reply = h
	return nil
}

func (s *ChatServer) GetHistory(_ struct{}, reply *[]ChatMessage) error {
	s.mu.Lock()
	h := make([]ChatMessage, len(s.history))
	copy(h, s.history)
	s.mu.Unlock()
	*reply = h
	return nil
}

func main() {
	port := os.Getenv("CHAT_PORT")
	if port == "" {
		port = "1234"
	}

	server := new(ChatServer)
	if err := rpc.Register(server); err != nil {
		log.Fatalf("register error: %v", err)
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}
	fmt.Printf("Chat server running on port %s...\n", port)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == "exit" {
				fmt.Println("Shutting down server...")
				listener.Close()
				return
			}
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				break
			}
			log.Printf("accept error: %v", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
