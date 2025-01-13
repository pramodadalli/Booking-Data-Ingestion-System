package routers

import (
    "booking-data-integration-system/controllers"
     "github.com/astaxie/beego"
)

func init() {
    beego.Router("/bookings", &controllers.BookingController{}, "post:CreateBooking")
    beego.Router("/bookings", &controllers.BookingController{}, "get:GetBookings")
    beego.Router("/bookings/:id", &controllers.BookingController{}, "get:GetBookingByID")
    beego.Router("/bookings/:id", &controllers.BookingController{}, "delete:DeleteBooking")
}
