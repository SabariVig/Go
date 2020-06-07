package main

import (
	"log"
	"golang.org/x/net/context"
	"github.com/sabarivig/go/grpc/chat"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":4545", grpc.WithInsecure())
	if err!= nil{
		log.Fatalf("Unable TO Connect To Server",err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)

	// message := chat.Message{
	// 	Body: "Hello From Client",
	// }

	response, err := c.SayHello(context.Background(), &chat.Message{Body:"lololol"})
	if err!= nil{
		log.Fatalf("Unable to call the function Hello",err)
	}
	log.Println("Response from server", response.Body)

}

