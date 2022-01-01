package usecase

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

type Service struct {
}

type ServiceManager interface {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetInvoice(ctx context.Context) (err error) {
	f, err := excelize.OpenFile("/Users/KevinTanuhardi/Documents/MVS/mvs_api/assets/format_invoice_mvs.xlsx")
	if err != nil {
			fmt.Println(err)
			return err
	}

	cell, err := f.GetCellValue("Sheet1", "D7")
	if err != nil {
		fmt.Println(err)
		return err
	}
	splittedCell := strings.Split(cell, "/")
	invoiceNum, _ := strconv.Atoi(splittedCell[0])
	invoiceNum += 1
	splittedCell[0] = fmt.Sprintf("%08d", invoiceNum)
	fmt.Println(splittedCell[0])
	err = f.SetCellValue("Sheet1", "D7", strings.Join(splittedCell, "/"))
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	err = f.SetCellValue("Sheet1", "D8", time.Now())
	if err != nil {
		return err
	}
	err = f.SetCellValue("Sheet1", "D9", time.Now().AddDate(0, 0, 2))
	if err != nil {
		return err
	}

	err = f.SetCellValue("Sheet1", "C20", 5)
	if err != nil {
		return err
	}
	err = f.SetCellValue("Sheet1", "D20", 25000)
	if err != nil {
		return err
	}

	err = f.SetCellValue("Sheet1", "C21", 5)
	if err != nil {
		return err
	}
	err = f.SetCellValue("Sheet1", "D21", 25000)
	if err != nil {
		return err
	}

	dateStyle, err := f.NewStyle(`{"number_format": 14}`)
	if err != nil {
			fmt.Println(err)
			return err
	}
	err = f.SetCellStyle("Sheet1", "D8", "D9", dateStyle)
	if err != nil {
			fmt.Println(err)
			return err
	}

	// TODO: save file with doctor name
	f.SaveAs("/Users/KevinTanuhardi/Documents/MVS/mvs_api/assets/invoice_mvs_test.xlsx")
	return nil
}
