package service

import (
	"context"
	"time"
	"google.golang.org/grpc"
	"github.com/Brotiger/gpu-4-ai-worker/proto"
)

type OllamaService struct {
	WorkerAddr string
}

func NewOllamaService(workerAddr string) *OllamaService {
	return &OllamaService{WorkerAddr: workerAddr}
}

func (s *OllamaService) withClient(f func(proto.WorkerClient, context.Context) error) error {
	connCtx, connCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer connCancel()
	conn, err := grpc.DialContext(connCtx, s.WorkerAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := proto.NewWorkerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return f(client, ctx)
}

func (s *OllamaService) Generate(req *proto.GenerateRequest) (*proto.GenerateResponse, error) {
	var resp *proto.GenerateResponse
	err := s.withClient(func(client proto.WorkerClient, ctx context.Context) error {
		var err error
		resp, err = client.Generate(ctx, req)
		return err
	})
	return resp, err
}

func (s *OllamaService) Tags() (*proto.TagsResponse, error) {
	var resp *proto.TagsResponse
	err := s.withClient(func(client proto.WorkerClient, ctx context.Context) error {
		var err error
		resp, err = client.Tags(ctx, &proto.TagsRequest{})
		return err
	})
	return resp, err
}

func (s *OllamaService) Show(req *proto.ShowRequest) (*proto.ShowResponse, error) {
	var resp *proto.ShowResponse
	err := s.withClient(func(client proto.WorkerClient, ctx context.Context) error {
		var err error
		resp, err = client.Show(ctx, req)
		return err
	})
	return resp, err
}

func (s *OllamaService) Pull(req *proto.PullRequest) (*proto.PullResponse, error) {
	var resp *proto.PullResponse
	err := s.withClient(func(client proto.WorkerClient, ctx context.Context) error {
		var err error
		resp, err = client.Pull(ctx, req)
		return err
	})
	return resp, err
}

func (s *OllamaService) Create(req *proto.CreateRequest) (*proto.CreateResponse, error) {
	var resp *proto.CreateResponse
	err := s.withClient(func(client proto.WorkerClient, ctx context.Context) error {
		var err error
		resp, err = client.Create(ctx, req)
		return err
	})
	return resp, err
}

func (s *OllamaService) Delete(req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	var resp *proto.DeleteResponse
	err := s.withClient(func(client proto.WorkerClient, ctx context.Context) error {
		var err error
		resp, err = client.Delete(ctx, req)
		return err
	})
	return resp, err
} 