package domain

import (
	"context"

	invoiceDto "github.com/kevintanuhardi/mvs_api/domain/invoice/dto"
	userDto "github.com/kevintanuhardi/mvs_api/domain/user/dto"
	"github.com/xuri/excelize/v2"
)

type UserDomainInterface interface {
	UserRegister(ctx context.Context, userData *userDto.RegisterUserRequest) (user *userDto.RegisterUserResponse, err error)
	Login(ctx context.Context, body *userDto.LoginRequest) (*userDto.LoginResponse, error)
}

type InvoiceDomainInterface interface {
	CreateInvoice(ctx context.Context, invoiceDetail *invoiceDto.CreateInvoiceRequest ) (file *excelize.File, filename string, err error)
	GetInvoiceXls(ctx context.Context, invoiceDetail *invoiceDto.GetInvoiceXlsRequest ) (file *excelize.File, filename string, err error)
}

type DomainService struct {
	User UserDomainInterface
	Invoice InvoiceDomainInterface
}

func NewDomain (user UserDomainInterface, invoice InvoiceDomainInterface) DomainService {
	return DomainService{
		User: user,
		Invoice: invoice,
	}
}
