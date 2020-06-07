package chat

import (
	"log"

	"golang.org/x/net/context"
)

// Server Is Struct 
type Server struct {
	UnimplementedChatServiceServer
}

//SayHello Just Says Hello When Called 
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Recived Message %s:", message.Body)
	return &Message{Body: "Hello From Server!!!"}, nil
}
