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
	GetInvoice(ctx context.Context, invoiceDetail *dto.GetInvoiceRequest ) (file *excelize.File, err error)
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

func (s *Service) GetInvoice(ctx context.Context, invoiceDetail *dto.GetInvoiceRequest ) (file *excelize.File, err error) {
	f, err := excelize.OpenFile("/Users/KevinTanuhardi/Documents/MVS/mvs_api/assets/format_invoice_mvs.xlsx")
	if err != nil {
			return nil, err
	}

	lastInvoice, err := s.invoice.GetLastInvoice(ctx)
	if err != nil {
		return nil, err
	}
	
	invoiceDate, err := time.Parse("02-01-2006", invoiceDetail.InvoiceDate)
	if err != nil {
		return nil, err
	}
	dueDate := invoiceDate.AddDate(0, 0, invoiceDetail.PaymentPeriodInDays)
	currentInvoiceNo := lastInvoice.InvoiceNo + 1

	fullInvoiceNum := []string{
		fmt.Sprintf("%08d", currentInvoiceNo),
		"MVS",
		constants.IntToRomNum[int(invoiceDate.Month())],
		fmt.Sprint(invoiceDate.Year()),
	}


	err = setExcelCell(f, constants.INVOICE_NO_AXIS, strings.Join(fullInvoiceNum, "/"))
	if err != nil {
		return nil, err
	}

	err = setExcelCell(f, constants.INVOICE_DATE_AXIS, invoiceDate)
	if err != nil {
		return nil, err
	}

	err = setExcelCell(f, constants.DUE_DATE_AXIS, dueDate)
	if err != nil {
		return nil, err
	}


	// STYLING date format
	dateStyle, err := f.NewStyle(`{"number_format": 15}`)
	if err != nil {
			fmt.Println(err)
			return nil, err
	}
	err = f.SetCellStyle("Sheet1", "D8", "D9", dateStyle)
	if err != nil {
			fmt.Println(err)
			return nil, err
	}

	err = setExcelCell(f, constants.SHIPPING_CONTACT_NAME_AXIS, invoiceDetail.ShippingContactName)
	if err != nil {
		return nil, err
	}
	err = setExcelCell(f, constants.SHIPPING_ADDRESS_AXIS, invoiceDetail.ShippingAddress)
	if err != nil {
		return nil, err
	}
	err = setExcelCell(f, constants.SHIPPING_CONTACT_PHONE_AXIS, invoiceDetail.ShippingContactPhone)
	if err != nil {
		return nil, err
	}
	// err = f.SetCellValue("Sheet1", "C20", 5)
	// if err != nil {
	// 	return nil, err
	// }
	// err = f.SetCellValue("Sheet1", "D20", 25000)
	// if err != nil {
	// 	return nil, err
	// }

	// err = f.SetCellValue("Sheet1", "C21", 5)
	// if err != nil {
	// 	return nil, err
	// }
	// err = f.SetCellValue("Sheet1", "D21", 25000)
	// if err != nil {
	// 	return nil, err
	// }

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
	s.invoice.CreateInvoice(ctx, invoiceEnt)

	// TODO: return file and download it from http
	// fileName := fmt.Sprintf(
	// 	"/Users/KevinTanuhardi/Documents/MVS/mvs_api/assets/invoice_mvs_%s_%d.xlsx", 
	// 	strings.Replace(invoiceDetail.ShippingContactName, " ", "_", -1),
	// 	currentInvoiceNo,
	// )
	// f.SaveAs(fileName)
	return f, nil;
}
