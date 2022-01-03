package invoice

import (
	"github.com/kevintanuhardi/mvs_api/config"
	"github.com/kevintanuhardi/mvs_api/domain/invoice/repository"
	"github.com/kevintanuhardi/mvs_api/domain/invoice/usecase"
)

func NewInvoice (cfg config.Config, invoice repository.Repository) usecase.ServiceManager {
	return usecase.NewService(invoice)
}