package controller

import "go-gateway/proto"

type streamService struct{}

// StreamService .
var StreamService = new(streamService)

func (s *streamService) List(req *proto.StreamRequest, stream proto.StreamService_ListServer) error {
	for n := 0; n <= 6; n++ {
		err := stream.Send(&proto.StreamResponse{
			Pt: &proto.StreamPoint{
				Name:  req.Pt.Name,
				Value: req.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *streamService) Record(stream proto.StreamService_RecordServer) error {
	return nil
}

func (s *streamService) Route(stream proto.StreamService_RouteServer) error {
	return nil
}
