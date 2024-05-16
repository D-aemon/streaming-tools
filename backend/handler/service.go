package handler

import (
	"backend/db"
	"backend/proto"
)

type StreamingToolsService struct {
	proto.UnimplementedStreamingToolsServiceServer
	db.DB
}

func (s StreamingToolsService) mustEmbedUnimplementedStreamingToolsServiceServer() {}
