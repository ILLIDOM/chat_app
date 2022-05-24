package main

import (
	"sync"

	pb "github.com/ILLIDOM/chat_app/proto"
	"github.com/google/uuid"
)

var (
	workChannel = Channel{
		name: "workChannel",
	}

	funChannel = Channel{
		name: "funChannel",
	}

	channelNameToChannel = make(map[string]Channel)
)

type User struct {
	id   string
	name string
}

type Channel struct {
	name string
	sync.RWMutex
	users    []*User
	messages chan *pb.Message
}

func addUserToChannel(channel *Channel, username string) {
	user := &User{
		id:   uuid.New().String(),
		name: username,
	}
	channel.Lock()
	defer channel.Unlock()
	channel.users = append(channel.users, user)
}

func createChannelMapping() {
	channelNameToChannel["workChannel"] = workChannel
	channelNameToChannel["funChannel"] = funChannel
}
