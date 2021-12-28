package user

import (
	"github.com/kevintanuhardi/mvs_api/config"
	"github.com/kevintanuhardi/mvs_api/domain/user/repository"
	"github.com/kevintanuhardi/mvs_api/domain/user/usecase"
)

func NewUser (cfg config.Config, users repository.Repository) usecase.ServiceManager {
	return usecase.NewService(users)
}