package entity

import "time"

type Otp struct {
	ID      *int      `json:"id,omitempty"`
	OwnerId string    `json:"owner_id"`
	Otp     string    `json:"otp"`
	Type    string    `json:"type"`
	ExpTime time.Time `json:"exp_time"`
}

func (Otp) TableName() string {
	return "otp"
}
