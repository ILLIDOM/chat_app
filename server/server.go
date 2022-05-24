package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/ILLIDOM/chat_app/proto"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50500, "Server listener port")
)

type chatServer struct {
	pb.UnimplementedChatServer
}

func (s *chatServer) JoinChannel(channel *pb.Channel, stream pb.Message) error {
	channel_to_join := channelNameToChannel[channel.Name]

	//add user to channel
	addUserToChannel(&channel_to_join, channel.User)

	for {
		select {
		case <- stream.Context().Done():
			return nil
		case msg := channel_to_join:
			stream.Send(msg)
		}
	}

}

func (s *chatServer) SendMessage(stream pb.Chat_SendMessageServer) error {
	message, err := stream.Recv()

	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	ack := pb.MessageAck{Status: "SENT"}
	stream.SendAndClose(&ack)


	go func() {
		curr_channel = channelNameToChannel[message.Channel.Name]
		curr_channel.
	}

}

func newChatServer() *chatServer {
	s := &chatServer{}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen on port: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterChatServer(grpcServer, newChatServer())
	grpcServer.Serve(lis)
	createChannelMapping()
}
