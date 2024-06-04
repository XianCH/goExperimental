package service

import (
	"context"
	"strings"
	"time"

	"github.com/x14n/goExperimental/go_advanced/rpc_protobuf/pub_sub/pb"
	"github.com/x14n/goExperimental/go_advanced/rpc_protobuf/pub_sub/pubsub"
)

type PubsubService struct {
	pub *pubsub.Publisher
	pb.UnimplementedPubsubServiceServer
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(10, time.Millisecond*100),
	}
}

func (p *PubsubService) Publish(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	p.pub.Publishe(req.GetMessage())
	return &pb.Response{}, nil
}

func (p *PubsubService) Subscribe(req *pb.Request, stream pb.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, req.GetMessage()) {
				return true
			}
		}
		return false
	})
	for v := range ch {
		if err := stream.Send(&pb.Response{Message: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}
