package entity

type Company struct {
	ID          *int   `json:"id,omitempty"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Country     string `json:"country"`
	Province    string `json:"province"`
	City        string `json:"city"`
	District    string `json:"district"`
	Village     string `json:"village"`
	PostalCode  string `json:"postal_code"`
	PhoneNumber string `json:"phone_number"`
	FaxNumber   string `json:"fax_number"`
	NPWP        string `json:"npwp"`
	SPPKP       string `json:"sppkp"`
}

func (Company) TableName() string {
	return "company"
}
