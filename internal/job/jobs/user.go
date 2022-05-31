package job

import (
	"context"

	"github.com/GodYao1995/Goooooo/internal/domain"
)

type UserESJob struct {
	repo domain.UserRepositoryFace
	espo domain.UserEsRepositoryFace
}

func NewUserESJob(repo domain.UserRepositoryFace, espo domain.UserEsRepositoryFace) *UserESJob {
	return &UserESJob{repo: repo, espo: espo}
}

// Just for testing purposes
func (us UserESJob) UsersToES() error {
	users, err := us.repo.GetAllUsers(context.Background())
	if err != nil {
		return err
	}
	objs := make([]*domain.UserEsPO, 0)
	for _, user := range users {
		deleteAt := ""
		if user.DeleteAt.Valid {
			deleteAt = user.DeleteAt.Time.Format("2006-01-02 15:04:05")
		}
		objs = append(objs, &domain.UserEsPO{
			UserName:  user.UserName,
			NickName:  user.NickName.String,
			Age:       uint8(user.Age.Int16),
			Gender:    uint8(user.Gender.Int16),
			Email:     user.Email.String,
			Phone:     user.Phone.String,
			State:     uint8(user.State.Int16),
			Ip:        uint64(user.Ip.Int64),
			LastLogin: user.LastLogin.String,
			UpdateAt:  user.UpdateAt,
			CreateAt:  user.CreateAt,
			DeleteAt:  deleteAt,
		})
	}
	return us.espo.CreateUserDocuments(context.Background(), objs)
}
