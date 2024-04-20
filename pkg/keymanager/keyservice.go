package keymanager

import (
	"bytes"
	"context"
	"encoding/pem"

	"github.com/zarix908/passwords_store/pkg/key"
	pb "github.com/zarix908/passwords_store/pkg/pb/keymanager/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const PassHeimdallPEMType = "PASSHEIMDALKEY"

func NewKeyServiceServer(
	keyGenerator key.KeyGenerator,
) pb.KeyServiceServer {
	return &keyServiceServer{
		keyGenerator: keyGenerator,
	}
}

type keyServiceServer struct {
	pb.UnimplementedKeyServiceServer

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

	buf := bytes.NewBuffer(nil)

	if err := pem.Encode(buf, &pem.Block{
		Type:  PassHeimdallPEMType,
		Bytes: generatedKey,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to encode to pem format: %v", err)
	}

	return &pb.GenerateKeyResponse{Key: string(buf.Bytes())}, nil
}
