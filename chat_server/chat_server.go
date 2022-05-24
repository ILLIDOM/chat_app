package chatserver

import "context"

type chaServer struct {
}

func (s *chatServer) JoinChannel(ctx context.Context, channel *pb.Channel) (*pb.Message, error) {

}
