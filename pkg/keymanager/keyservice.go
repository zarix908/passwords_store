package keymanager

import (
	"context"
	pb "github.com/zarix908/passwords_store/pkg/pb/keymanager/v1"
)

func NewKeyServiceServer() pb.KeyServiceServer {
	return &keyServiceServer{}
}

type keyServiceServer struct{
	pb.UnimplementedKeyServiceServer
}

func (k *keyServiceServer) GenerateKey(context.Context, *pb.GenerateKeyRequest) (*pb.GenerateKeyResponse, error) {
	panic("unimplemented")
}
