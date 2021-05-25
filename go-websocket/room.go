package main

import (
	"github.com/lolmourne/go-groupchat/client"
	"github.com/lolmourne/go-websocket/resource/chat"
	"github.com/lolmourne/go-websocket/resource/user"
)

type RoomManager struct {
	hubs    map[int64]*Hub
	cli     client.GroupchatClientItf
	userRsc user.IResource
	chatRsc chat.IResource
}

type RoomManagerItf interface {
	JoinRoom(roomID int64) *Hub
}

func NewRoomManager(cli client.GroupchatClientItf, chatRsc chat.IResource, userRsc user.IResource) RoomManagerItf {
	return &RoomManager{
		hubs:    make(map[int64]*Hub),
		cli:     cli,
		chatRsc: chatRsc,
		userRsc: userRsc,
	}
}

func (r *RoomManager) JoinRoom(roomID int64) *Hub {
	room := r.cli.GetGroupchatRoom(roomID)
	if room == nil {
		return nil
	}

	hub, ok := r.hubs[roomID]
	if !ok {
		hub = NewHub(roomID, sub, rdb, r.chatRsc, r.userRsc)
		r.hubs[roomID] = hub
		go r.hubs[roomID].run()
		return hub
	}
	return hub
}
