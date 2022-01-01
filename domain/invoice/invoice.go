package invoice

import (
	"github.com/kevintanuhardi/mvs_api/config"
	"github.com/kevintanuhardi/mvs_api/domain/invoice/usecase"
)

func NewInvoice (cfg config.Config) usecase.ServiceManager {
	return usecase.NewService()
}