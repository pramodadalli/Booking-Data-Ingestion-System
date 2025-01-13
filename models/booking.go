package models

import (
    "time"
)

type Booking struct {
    ID          int       `json:"id" orm:"auto"`
    Customer    string    `json:"customer" orm:"size(100)"`
    BookingDate time.Time `json:"booking_date" orm:"type(datetime)"`
    Amount      float64   `json:"amount" orm:"digits(12);decimals(2)"`
    Vendor      string    `json:"vendor" orm:"size(100)"`
}
