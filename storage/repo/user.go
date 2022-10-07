package repo
import (
	"github.com/jmoiron/sqlx"
	pb "github.com/Jamshidbek1608/temp-service/genproto"
)

type userRepo struct{
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *pb.User) (*pb.User, error) {
	userRepo := pb.User{}
	err:=r.
}

