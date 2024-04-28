package crypt

import (
	"context"

	"github.com/zarix908/passwords_store/pkg/encryptor"
	pb "github.com/zarix908/passwords_store/pkg/pb/crypt/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewCryptServiceServer(encryptor encryptor.Encryptor) pb.CryptServiceServer {
	return &cryptServiceServer{
		encryptor: encryptor,
	}
}

type cryptServiceServer struct {
	encryptor encryptor.Encryptor

	pb.UnimplementedCryptServiceServer
}

func (srv *cryptServiceServer) EncryptFile(
	ctx context.Context,
	req *pb.EncryptFileRequest,
) (*pb.EncryptFileResponse, error) {
	if req.AgeRecipients == "" {
		return nil, status.Error(codes.InvalidArgument, "at least one age recipient must be provided")
	}

	encrypted, err := srv.encryptor.Encrypt(req.File, req.AgeRecipients)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to encrypt file: %v", err)
	}

	return &pb.EncryptFileResponse{File: encrypted}, nil
}
