package repo

import (
	pb "github.com/Jamshidbek1608/temp-service/genproto"
)

// UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
}
