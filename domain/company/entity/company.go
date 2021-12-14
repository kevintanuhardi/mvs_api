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

type CompanyObject struct {
	Code        string 
	Name        string 
	Address     string 
	Country     string 
	Province    string 
	City        string 
	District    string 
	Village     string 
	PostalCode  string 
	PhoneNumber string 
	FaxNumber   string 
	NPWP        string 
	SPPKP       string 
}

func (Company) TableName() string {
	return "company"
}
