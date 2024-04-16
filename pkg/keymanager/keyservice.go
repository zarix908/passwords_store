package keymanager

import (
	"context"

	"github.com/zarix908/passwords_store/pkg/key"
	pb "github.com/zarix908/passwords_store/pkg/pb/keymanager/v1"
	"github.com/zarix908/passwords_store/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewKeyServiceServer(
	fileStoreFactory func(filePath string) storage.Storage,
	keyGenerator key.KeyGenerator,
) pb.KeyServiceServer {
	return &keyServiceServer{
		newStore:     fileStoreFactory,
		keyGenerator: keyGenerator,
	}
}

type keyServiceServer struct {
	pb.UnimplementedKeyServiceServer

	newStore     func(filePath string) storage.Storage
	keyGenerator key.KeyGenerator
}

func (srv *keyServiceServer) GenerateKey(
	ctx context.Context,
	req *pb.GenerateKeyRequest,
) (*pb.GenerateKeyResponse, error) {
	generatedKey, err := srv.keyGenerator.GenerateEncrypted(req.Password)
	if err != nil {
		return nil, status.Newf(codes.Internal, "failed to generate new key: %s", err.Error()).Err()
	}

	store := srv.newStore(req.FilePath)
	if err := store.AddKey(generatedKey); err != nil {
		return nil, status.Newf(codes.Internal, "failed add key: %s", err.Error()).Err()
	}

	return &pb.GenerateKeyResponse{FilePath: req.FilePath}, nil
}
