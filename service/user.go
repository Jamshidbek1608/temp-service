package service

import (
	"context"

	pb "github.com/Jamshidbek1608/temp-service/genproto"
	l "github.com/Jamshidbek1608/temp-service/pkg/logger"
	"github.com/Jamshidbek1608/temp-service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().Create(req)
	if err != nil {
		s.logger.Error("error insert", l.Any("error insert user", err))
		return &pb.User{}, status.Error(codes.Internal, "something went wrong, please check user info")
	}

	return user, nil
}
