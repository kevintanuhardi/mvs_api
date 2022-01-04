package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/kevintanuhardi/mvs_api/domain/invoice/dto"
	"github.com/kevintanuhardi/mvs_api/domain/invoice/entity"
	"github.com/kevintanuhardi/mvs_api/domain/invoice/repository"
	"github.com/kevintanuhardi/mvs_api/internal/constants"
	"github.com/xuri/excelize/v2"
)

type Service struct {
	invoice repository.Repository
}

type ServiceManager interface {
	CreateInvoice(ctx context.Context, invoiceDetail *dto.CreateInvoiceRequest ) (file *excelize.File, filename string, err error)
	GetInvoiceXls(ctx context.Context, invoiceDetail *dto.GetInvoiceXlsRequest ) (file *excelize.File, filename string, err error)
}

func NewService(invoice repository.Repository) *Service {
	return &Service{invoice}
}

func setExcelCell(f *excelize.File, axis string ,content interface{}) error {
	err := f.SetCellValue("Sheet1", axis, content)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CreateInvoice(ctx context.Context, invoiceDetail *dto.CreateInvoiceRequest ) (invoice *dto.CreateInvoiceResponse, err error) {
	lastInvoice, err := s.invoice.GetLastInvoice(ctx)
	if err != nil {
		return nil, err
	}
	
	invoiceDate, err := time.Parse("02-01-2006", invoiceDetail.InvoiceDate)
	if err != nil {
		return nil, err
	}
	currentInvoiceNo := lastInvoice.InvoiceNo + 1

	invoiceEnt := &entity.Invoice{
		InvoiceNo:            currentInvoiceNo,
		InvoiceDate:          invoiceDate,
		PaymentPeriodInDays:  invoiceDetail.PaymentPeriodInDays,
		ShippingContactName:  invoiceDetail.ShippingContactName,
		ShippingContactPhone: invoiceDetail.ShippingContactPhone,
		ShippingAddress:      invoiceDetail.ShippingAddress,
		ShippingCity:         invoiceDetail.ShippingCity,
		ShippingPostalCode:   invoiceDetail.ShippingPostalCode,
		BillingContactName:   invoiceDetail.BillingContactName,
		BillingContactPhone:  invoiceDetail.BillingContactPhone,
		BillingAddress:       invoiceDetail.BillingAddress,
		BillingCity:          invoiceDetail.BillingCity,
		BillingPostalCode:    invoiceDetail.BillingPostalCode,
		Status:               "unpaid",
		CreatedAt:            time.Time{},
		UpdatedAt:            time.Time{},
	}
	invoiceEnt, err = s.invoice.CreateInvoice(ctx, invoiceEnt)
	if err != nil {
		return nil, err
	}

	invoice = &dto.CreateInvoiceResponse{
		InvoiceNo:            invoiceEnt.InvoiceNo,
		InvoiceDate:          invoiceEnt.InvoiceDate.Format("02-01-2006"),
		PaymentPeriodInDays:  invoiceEnt.PaymentPeriodInDays,
		ShippingContactName:  invoiceEnt.ShippingContactName,
		ShippingContactPhone: invoiceEnt.ShippingContactPhone,
		ShippingAddress:      invoiceEnt.ShippingAddress,
		ShippingCity:         invoiceEnt.ShippingCity,
		ShippingPostalCode:   invoiceEnt.ShippingPostalCode,
		BillingContactName:   invoiceEnt.BillingContactName,
		BillingContactPhone:  invoiceEnt.BillingContactPhone,
		BillingAddress:       invoiceEnt.BillingAddress,
		BillingCity:          invoiceEnt.BillingCity,
		BillingPostalCode:    invoiceEnt.BillingPostalCode,
		Status:               invoiceEnt.Status,
	}

	return invoice, nil

}

func (s *Service) GetInvoiceXls(ctx context.Context, invoiceDetail *dto.GetInvoiceXlsRequest ) (file *excelize.File, filename string, err error) {
	f, err := excelize.OpenFile("/Users/KevinTanuhardi/Documents/MVS/mvs_api/assets/format_invoice_mvs.xlsx")
	if err != nil {
			return nil, "", err
	}

	invoice, err := s.invoice.FindInvoiceByInvoiceId(ctx, invoiceDetail.InvoiceId)
	if err != nil {
		return nil, "", err
	}

	dueDate := invoice.InvoiceDate.AddDate(0, 0, invoice.PaymentPeriodInDays)

	fullInvoiceNum := []string{
		fmt.Sprintf("%08d", invoice.InvoiceNo),
		"MVS",
		constants.IntToRomNum[int(invoice.InvoiceDate.Month())],
		fmt.Sprint(invoice.InvoiceDate.Year()),
	}

	err = setExcelCell(f, constants.INVOICE_NO_AXIS, strings.Join(fullInvoiceNum, "/"))
	if err != nil {
		return nil, "", err
	}

	err = setExcelCell(f, constants.INVOICE_DATE_AXIS, invoice.InvoiceDate)
	if err != nil {
		return nil, "", err
	}

	err = setExcelCell(f, constants.DUE_DATE_AXIS, dueDate)
	if err != nil {
		return nil, "", err
	}


	// STYLING date format
	dateStyle, err := f.NewStyle(`{"number_format": 15}`)
	if err != nil {
			fmt.Println(err)
			return nil, "", err
	}
	err = f.SetCellStyle("Sheet1", "D8", "D9", dateStyle)
	if err != nil {
			fmt.Println(err)
			return nil, "", err
	}

	err = setExcelCell(f, constants.SHIPPING_CONTACT_NAME_AXIS, invoice.ShippingContactName)
	if err != nil {
		return nil, "", err
	}
	err = setExcelCell(f, constants.SHIPPING_ADDRESS_AXIS, invoice.ShippingAddress)
	if err != nil {
		return nil, "", err
	}
	err = setExcelCell(f, constants.SHIPPING_CONTACT_PHONE_AXIS, invoice.ShippingContactPhone)
	if err != nil {
		return nil, "", err
	}
	// err = f.SetCellValue("Sheet1", "C20", 5)
	// if err != nil {
	// 	return nil, "", err
	// }
	// err = f.SetCellValue("Sheet1", "D20", 25000)
	// if err != nil {
	// 	return nil, "", err
	// }

	// err = f.SetCellValue("Sheet1", "C21", 5)
	// if err != nil {
	// 	return nil, "", err
	// }
	// err = f.SetCellValue("Sheet1", "D21", 25000)
	// if err != nil {
	// 	return nil, "", err
	// }

	filename = fmt.Sprintf(
		"invoice_mvs_%s_%d.xlsx", 
		strings.Replace(invoice.ShippingContactName, " ", "_", -1),
		invoice.InvoiceNo,
	)
	// f.SaveAs(fileName)
	return f, filename, nil;
}