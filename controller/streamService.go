package controller

import "go-gateway/proto"

type streamService struct{}

// StreamService .
var StreamService = new(streamService)

func (s *streamService) List(r *proto.StreamRequest, stream proto.StreamService_ListServer) error {
	return nil
}

func (s *streamService) Record(stream proto.StreamService_RecordServer) error {
	return nil
}

func (s *streamService) Route(stream proto.StreamService_RouteServer) error {
	return nil
}
