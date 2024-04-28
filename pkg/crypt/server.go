package crypt

import (
	"context"

	"github.com/zarix908/passwords_store/pkg/encryption"
	pb "github.com/zarix908/passwords_store/pkg/pb/crypt/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewCryptServiceServer(encryptor encryption.Encryptor) pb.CryptServiceServer {
	return &cryptServiceServer{
		encryptor: encryptor,
	}
}

type cryptServiceServer struct {
	encryptor encryption.Encryptor

	pb.UnimplementedCryptServiceServer
}

func (srv *cryptServiceServer) Encrypt(
	ctx context.Context,
	req *pb.EncryptRequest,
) (*pb.EncryptResponse, error) {
	if req.Recipients == "" {
		return nil, status.Error(codes.InvalidArgument, "at least one recipient must be provided")
	}

	encrypted, err := srv.encryptor.Encrypt(req.Data, req.Recipients)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to encrypt file: %v", err)
	}

	return &pb.EncryptResponse{Data: encrypted}, nil
}
