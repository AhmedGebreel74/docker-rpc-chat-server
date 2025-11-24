package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
	"time"
)

type ChatMessage struct {
	Author    string
	Text      string
	Timestamp time.Time
}

func printHistory(history []ChatMessage) {
	fmt.Println("\n--- Chat History ---")
	for _, m := range history {
		fmt.Printf("%s: %s\n", m.Author, m.Text)
	}
	fmt.Println("--------------------")
}

func main() {
	addr := os.Getenv("CHAT_ADDR")
	if addr == "" {
		addr = "localhost:1234"
	}

	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("failed to connect to server at %s: %v", addr, err)
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		name = "Anonymous"
	}
	fmt.Printf("Welcome %s! You've joined the chat. Type a message to see the chat history.\n", name)

	var history []ChatMessage
	if err := client.Call("ChatServer.GetHistory", struct{}{}, &history); err == nil && len(history) > 0 {
		printHistory(history)
	}

	for {
		fmt.Print("Enter message (or 'exit' to quit): ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if strings.EqualFold(line, "exit") {
			fmt.Println("Bye!")
			return
		}
		if line == "" {
			continue
		}

		msg := ChatMessage{
			Author:    name,
			Text:      line,
			Timestamp: time.Now(),
		}

		history = nil
		if err := client.Call("ChatServer.SendMessage", msg, &history); err != nil {
			log.Printf("server error: %v", err)
			fmt.Println("The server might be down. Try again later.")
			return
		}

		printHistory(history)
	}
}
