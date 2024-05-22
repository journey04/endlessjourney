package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
	"os"
	"github.com/gin-gonic/gin"

)

type BookingDetails struct {
	Drop_of_Location string `json:"drop_of_location"`
	No_of_Days       string `json:"no_of_days"`
	PickUpDay        string `json:"pick_up_day"`
	PickUpDate       string `json:"pick_up_date"`
	PickUpTiming     string `json:"pick_up_timing"`
	PickUpAddress    string `json:"pick_up_address"`
	DropOfDay        string `json:"drop_of_day"`
	DropOfDate       string `json:"drop_of_date"`
	DropOfTiming     string `json:"drop_of_timing"`
	DropOfAddress    string `json:"drop_of_address"`
	ConfirmationNo   string `json:"car_confirmation_no"`
	BillingAmount    string `json:"billing_amount"`
	CarDetails       string `json:"car_details"`
	DriverName       string `json:"driver_name"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	BookingNo        string `json:"bookingNo"`
  Link             string `json:"link"`
}

type Pickup struct {
	PickUpDay       string `json:"pick_up_day"`
	PickUpDate      string `json:"pick_up_date"`
	DepartsTiming   string `json:"departs_timing"`
	DepartsLocation string `json:"departs_location"`
	ArrivalTiming   string `json:"arrival_timing"`
	ArrivalLocation string `json:"arrival_location"`
	FlightDuration  string `json:"flight_duration"`
	AirlineName     string `json:"airlineName"`
}

type Dropof struct {
	DropOfDay       string `json:"pick_up_day"`
	DropOfDate      string `json:"pick_up_date"`
	DepartsTiming   string `json:"departs_timing"`
	DepartsLocation string `json:"departs_location"`
	ArrivalTiming   string `json:"arrival_timing"`
	ArrivalLocation string `json:"arrival_location"`
	FlightDuration  string `json:"flight_duration"`
	AirlineName     string `json:"airlineName"`
}

type FlightBookingDetails struct {
	Pickupinfo         Pickup   `json:"pickup_info"`
	Dropofinfo         Dropof   `json:"dropof_info"`
	BillingAmount      string   `json:"billing_amount"`
	TripType           string   `json:"trip_type"`
	AirportName        string   `json:"airport_name"`
	PassengerName      []string `json:"passenger_name"`
	Email              string   `json:"email"`
	Phone              string   `json:"phone"`
	BookingNo          string   `json:"bookingNo"`
	PaymentPaidBy      string   `json:"payment_paid"`
	PetName            string   `json:"pet_name"`
	BookingDate        string   `json:"booking_date"`
	TripPickupLocation string   `json:"tripPickupLocation"`
	TripDropofLocation string   `json:"tripDropofLocation"`
	Name               string   `json:"name"`
  Link               string   `json:"link"`
}

type HotelBookingDetails struct {
	BillingAmount      string   `json:"billing_amount"`
	CheckinDate        string   `json:"checkin_date"`
	CheckinTime        string   `json:"checkin_time"`
	CheckoutDate       string   `json:"checkout_date"`
	CheckoutTime       string   `json:"checkout_time"`
	RoomType           string   `json:"room_type"`
	HotelName          string   `json:"hotel_name"`
	HotelRatting       string   `json:"hotel_ratting"`
	HotelEmail         string   `json:"hotel_email"`
	HotelPhone         string   `json:"hotel_phone"`
	HotelAddress       string   `json:"hotel_address"`
	RoomsNo            string   `json:"room_no"`
	AdultsNo           string   `json:"adults_no"`
	ChildsNo           string   `json:"childs_no"`
	GuestName          []string `json:"passenger_name"`
	Email              string   `json:"email"`
	Phone              string   `json:"phone"`
	BookingNo          string   `json:"bookingNo"`
	PaymentPaidBy      string   `json:"payment_paid"`
	PetName            string   `json:"pet_name"`
	BookingDate        string   `json:"booking_date"`
	Name               string   `json:"name"`
  CancelPolicy       string   `json:"cancel_policy"`
  Link               string    `json:"link"`
}

type EmailRequest struct {
	To             string         `json:"to"`
	Subject        string         `json:"subject"`
	BookingDetails BookingDetails `json:"booking_details"`
}

type FlightEmailRequest struct {
	To                   string               `json:"to"`
	Subject              string               `json:"subject"`
	FlightBookingDetails FlightBookingDetails `json:"flight_booking_details"`
}
type HotelEmailRequest struct {
	To                  string              `json:"to"`
	Subject             string              `json:"subject"`
	HotelBookingDetails HotelBookingDetails `json:"hotel_booking_details"`
}

type Verifier struct {
	UserName       string `json:"username"`
	VerifyPassword string `json:"password"`
}

const carEmailHTML = `
<!DOCTYPE html>
<html>
<head>
    <style>
        .header {
            background: #243c54;
            padding: 10px;
            text-align: center;
       }
        .container {
            width: 650px;
            margin: 0 auto;
            padding: 20px;
            border: 1px solid #d1d2d3;
            border-radius: 5px;
        }
        .title {
            font-size: 24px;
            font-weight: bold;
            color: #076ba7;
            margin-bottom: 10px;
        }
        .content {
            font-size: 14px;
            line-height: 1.5;
            color: #666;
        }
        .footer {
            background: #f5f6f7;
            padding: 10px;
            text-align: center;
            margin-top: 20px;
            border-radius: 0 0 5px 5px;
        }
    </style>
</head>
<body>
    <div class="header">
        <img src="https://traveloment.com/wp-content/uploads/2023/08/trav_logo-removebg-preview-768x213.png" alt="Traveloment" width="150">
    </div>
    <div class="container">
        <div class="title">Booking Acknowledgement</div>
        <div class="content">
            <p>Dear Customer,</p>
            <p>Your booking acknowledgement for Traveloment Booking: {{.BookingNo}} has been received. This email does not confirm successful ticketing. Confirmation of your booking will be sent in a separate email.</p>
            <table cellspacing="0" cellpadding="0" width="650" border="0" align="center" style="font-family:Arial;font-size:14px;color:#666666">
                <tbody>
                    <tr>
                        <td height="20">&nbsp;</td>
                    </tr>
                    <tr>
                        <td style="padding-bottom:5px">
                            <table cellspacing="0" cellpadding="0" width="100%" border="0">
                                <tr>
                                    <td width="60%" style="color:#076ba7;font-size:24px;font-family:Arial">
                                        Car Details
                                    </td>
                                    <td width="30%" style="color:#666666">
                                        <strong>Status:</strong>
                                        <a href="https://traveloment.com/">
                                            Check now
                                        </a>
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>
                    <tr>
                       <td style="border:1px solid #d4d4d4;border-bottom:1px solid #9f9f9f">
    <table cellspacing="0" cellpadding="0" width="100%" border="0" style="font-family:Arial;font-size:12px;color:#666666">
        <tr>
            <td style="background-color:#ffffff;padding:15px">
                <table cellspacing="0" cellpadding="0" width="100%" border="0" style="font-family:Arial;font-size:12px;color:#666666">
                    <tr>
                        <td style="color:#056ba8;font-size:20px;font-weight:bold;padding-right:25px" valign="top" width="26%" height="50">
                            <div>Special</div>
                        </td>
                        <td style="font-size:16px;padding-right:25px" valign="top" width="44%">
                            Location:<br>
                            <strong>{{.Drop_of_Location}}</strong>
                        </td>
                        <td style="font-size:16px" valign="top" width="30%">
                            Days:<br>
                            <strong>{{.No_of_Days}}</strong>
                        </td>
                    </tr>
                    <tr>
                        
<td style="" valign="top">
                                                    <table cellspacing="0" cellpadding="0" width="" border="0" style="font-family:Arial;font-size:12px;color:#666666">
                                                       
                                                       
                                                       
                                                        <tr>
                                                            <td style="padding-top: 30px" align="center">
                                                                <img style="border:none; width:80%" src="https://i.pinimg.com/originals/b6/6e/14/b66e14f44dd38a6fac2aff3207345058.gif" >
                                                            </td>
                                                        </tr>
                                                    </table>
                                                </td>
                        <td valign="top" style="padding:12px 25px 0 0">
                            <table cellspacing="0" cellpadding="0" width="100%" border="0" style="font-family:Arial;font-size:12px;color:#666666">
                                <tr>
                                    <td valign="top" style="padding-bottom:10px">
                                        Car Pick-Up
                                    </td>
                                </tr>
                                <tr>
                                    <td valign="top" style="padding-bottom:15px">
                                        <strong style="font-size:20px" class="notranslate">
                                        {{.PickUpDay}}, {{.PickUpDate}} - {{.PickUpTiming}}
                                        </strong>
                                    </td>
                                </tr>
                                <tr>
                                    <td valign="top" style="padding-bottom:20px;line-height:17px">
                                        {{.PickUpAddress}}
                                    </td>
                                </tr>
                                <tr>
                                    <td valign="top" style="padding-bottom:10px">
                                        Car Drop-Off
                                    </td>
                                </tr>
                                <tr>
                                    <td valign="top" style="padding-bottom:15px">
                                        <strong style="font-size:20px" class="notranslate">
                                        {{.DropOfDay}}, {{.DropOfDate}} - {{.DropOfTiming}}
                                        </strong>
                                    </td>
                                </tr>
                                <tr>
                                    <td valign="top" style="line-height:17px">
                                    {{.DropOfAddress}}
                                    </td>
                                </tr>
                            </table>
                        </td>
                        <td valign="top">
                            <table cellspacing="0" cellpadding="0" width="100%" border="0" style="font-family:Arial;font-size:12px;color:#666666">
                                <tr>
                                    <td style="background-color:#f6f6f6;padding:15px;line-height:16px">
                                        Car Confirmation: <br>
                                        <strong>{{.ConfirmationNo}}</strong>
                                        <br>
                                        <br>
                                        <br>
                                        <span>Total Charge: : <strong><span>{{.BillingAmount}}</span></strong></span>
                                        <br>
                                        <br>
                                        <a href={{.Link}} style="display: inline-block; width: 100%; padding: 0.5rem; background-color: #243c54; color: #f5dd42; text-align: center; text-decoration: none;">
                    Pay Now
                </a>
                                    </td>
                                </tr>
                                <tr>
                                    <td height="8">&nbsp;</td>
                                </tr>
                            </table>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</td>

                                </tr>
                                <tr>
                                    <td style="padding:15px;background-color:#ffffff">
                                        <table cellspacing="0" cellpadding="0" width="100%" border="0" style="font-family:Arial;font-size:12px;color:#666666">
                                            <tr>
                                                <td width="26%" style="padding-right:25px" valign="top" height="30">
                                                    Type
                                                </td>
                                                <td width="44%" valign="top">
                                                    Driver
                                                </td>
                                               
                                            </tr>
                                            <tr>
                                                <td valign="top" style="padding-right:25px;line-height:17px">
                                                    <span style="font-size:16px;font-weight:bold">
                                                        {{.CarDetails}}
                                                    </span>
                                                </td>
                                                <td style="font-size:16px;font-weight:bold;text-transform:capitalize" valign="top">
                                                    {{.DriverName}}
                                                </td>
                                                
                                            </tr>
                                        </table>
                                    </td>
                                </tr>
                               
                            </table>
                        </td>
                    </tr>
                </tbody>
            </table>
            <p>Your Details:</p>
            <table cellspacing="0" cellpadding="0" width="100%" border="0" style="font-family:Arial;font-size:12px">
                <tbody>
                    <tr>
                        <td valign="top" style="padding-bottom:10px">Name:</td>
                        <td valign="top" style="padding-bottom:10px">{{.Name}}</td>
                    </tr>
                    <tr>
                        <td valign="top" style="padding-bottom:10px">Email:</td>
                        <td valign="top" style="padding-bottom:10px">{{.Email}}</td>
                    </tr>
                    <tr>
                        <td valign="top" style="padding-bottom:10px">Phone:</td>
                        <td valign="top" style="padding-bottom:10px">{{.Phone}}</td>
                    </tr>
                    <!-- Add more rows for billing details as needed -->
                </tbody>
            </table>
        </div>
    </div>
    <div class="footer">
        <p>For any changes with your flight, date, route or names, call us at +1(800) 986-6901 or click <a href="https://traveloment.com/" style="color:#005083;font-size:12px;text-decoration:none;font-weight:bold" rel="noreferrer noreferrer noreferrer" target="_blank">here to chat with us</a>.</p>
        <p>Copyright © 2024 Traveloment. All rights reserved.</p>
    </div>
</body>
</html>
`

const flightEmailHtml = `
<html
  dir="ltr"
  xmlns="http://www.w3.org/1999/xhtml"
  xmlns:o="urn:schemas-microsoft-com:office:office"
  lang="und"
  style="font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;font-family:&amp;apos;trebuchet ms&amp;apos;, &amp;apos;lucida grande&amp;apos;, &amp;apos;lucida sans unicode&amp;apos;, &amp;apos;lucida sans&amp;apos;, tahoma, sans-serif"
>
  <head>
    <meta charset="UTF-8" />
    <meta content="width=device-width, initial-scale=1" name="viewport" />
    <meta name="x-apple-disable-message-reformatting" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta content="telephone=no" name="format-detection" />
    <title>Booking confirmation</title>
    <!--[if (mso 16)
      ]><style type="text/css">
        a {
          text-decoration: none;
        }
      </style>
    <![endif]-->
    <!--[if gte mso 9
      ]><style>
        sup {
          font-size: 100% !important;
        }
      </style><!
    [endif]-->
    <!--[if gte mso 9
      ]><xml>
        <o:OfficeDocumentSettings>
          <o:AllowPNG></o:AllowPNG> <o:PixelsPerInch>96</o:PixelsPerInch>
        </o:OfficeDocumentSettings>
      </xml>
    <![endif]-->
    
    <base href="#" />
    <style type="text/css">
      #outlook a {
        padding: 0;
      }
      .es-button {
        mso-style-priority: 100 !important;
        text-decoration: none !important;
      }
      a[x-apple-data-detectors] {
        color: inherit !important;
        text-decoration: none !important;
        font-size: inherit !important;
        font-family: inherit !important;
        font-weight: inherit !important;
        line-height: inherit !important;
      }
      .es-desk-hidden {
        display: none;
        float: left;
        overflow: hidden;
        width: 0;
        max-height: 0;
        line-height: 0;
        mso-hide: all;
      }
      @media only screen and (max-width: 600px) {
        p,
        ul li,
        ol li,
        a {
          line-height: 150% !important;
        }
        h1,
        h2,
        h3,
        h1 a,
        h2 a,
        h3 a {
          line-height: 120%;
        }
        h1 {
          font-size: 40px !important;
          text-align: center;
        }
        h2 {
          font-size: 28px !important;
          text-align: center;
        }
        h3 {
          font-size: 20px !important;
          text-align: left;
        }
        .es-header-body h1 a,
        .es-content-body h1 a,
        .es-footer-body h1 a {
          font-size: 40px !important;
          text-align: center;
        }
        .es-header-body h2 a,
        .es-content-body h2 a,
        .es-footer-body h2 a {
          font-size: 28px !important;
          text-align: center;
        }
        .es-header-body h3 a,
        .es-content-body h3 a,
        .es-footer-body h3 a {
          font-size: 20px !important;
          text-align: left;
        }
        .es-menu td a {
          font-size: 14px !important;
        }
        .es-header-body p,
        .es-header-body ul li,
        .es-header-body ol li,
        .es-header-body a {
          font-size: 14px !important;
        }
        .es-content-body p,
        .es-content-body ul li,
        .es-content-body ol li,
        .es-content-body a {
          font-size: 14px !important;
        }
        .es-footer-body p,
        .es-footer-body ul li,
        .es-footer-body ol li,
        .es-footer-body a {
          font-size: 14px !important;
        }
        .es-infoblock p,
        .es-infoblock ul li,
        .es-infoblock ol li,
        .es-infoblock a {
          font-size: 12px !important;
        }
        *[class="gmail-fix"] {
          display: none !important;
        }
        .es-m-txt-c,
        .es-m-txt-c h1,
        .es-m-txt-c h2,
        .es-m-txt-c h3 {
          text-align: center !important;
        }
        .es-m-txt-r,
        .es-m-txt-r h1,
        .es-m-txt-r h2,
        .es-m-txt-r h3 {
          text-align: right !important;
        }
        .es-m-txt-l,
        .es-m-txt-l h1,
        .es-m-txt-l h2,
        .es-m-txt-l h3 {
          text-align: left !important;
        }
        .es-m-txt-r img,
        .es-m-txt-c img,
        .es-m-txt-l img {
          display: inline !important;
        }
        .es-button-border {
          display: inline-block !important;
        }
        a.es-button,
        button.es-button {
          font-size: 18px !important;
          display: inline-block !important;
        }
        .es-adaptive table,
        .es-left,
        .es-right {
          width: 100% !important;
        }
        .es-content table,
        .es-header table,
        .es-footer table,
        .es-content,
        .es-footer,
        .es-header {
          width: 100% !important;
          max-width: 600px !important;
        }
        .es-adapt-td {
          display: block !important;
          width: 100% !important;
        }
        .adapt-img {
          width: 100% !important;
          height: auto !important;
        }
        .es-m-p0 {
          padding: 0 !important;
        }
        .es-m-p0r {
          padding-right: 0 !important;
        }
        .es-m-p0l {
          padding-left: 0 !important;
        }
        .es-m-p0t {
          padding-top: 0 !important;
        }
        .es-m-p0b {
          padding-bottom: 0 !important;
        }
        .es-m-p20b {
          padding-bottom: 20px !important;
        }
        .es-mobile-hidden,
        .es-hidden {
          display: none !important;
        }
        tr.es-desk-hidden,
        td.es-desk-hidden,
        table.es-desk-hidden {
          width: auto !important;
          overflow: visible !important;
          float: none !important;
          max-height: inherit !important;
          line-height: inherit !important;
        }
        tr.es-desk-hidden {
          display: table-row !important;
        }
        table.es-desk-hidden {
          display: table !important;
        }
        td.es-desk-menu-hidden {
          display: table-cell !important;
        }
        .es-menu td {
          width: 1% !important;
        }
        table.es-table-not-adapt,
        .esd-block-html table {
          width: auto !important;
        }
        table.es-social {
          display: inline-block !important;
        }
        table.es-social td {
          display: inline-block !important;
        }
        .es-m-p5 {
          padding: 5px !important;
        }
        .es-m-p5t {
          padding-top: 5px !important;
        }
        .es-m-p5b {
          padding-bottom: 5px !important;
        }
        .es-m-p5r {
          padding-right: 5px !important;
        }
        .es-m-p5l {
          padding-left: 5px !important;
        }
        .es-m-p10 {
          padding: 10px !important;
        }
        .es-m-p10t {
          padding-top: 10px !important;
        }
        .es-m-p10b {
          padding-bottom: 10px !important;
        }
        .es-m-p10r {
          padding-right: 10px !important;
        }
        .es-m-p10l {
          padding-left: 10px !important;
        }
        .es-m-p15 {
          padding: 15px !important;
        }
        .es-m-p15t {
          padding-top: 15px !important;
        }
        .es-m-p15b {
          padding-bottom: 15px !important;
        }
        .es-m-p15r {
          padding-right: 15px !important;
        }
        .es-m-p15l {
          padding-left: 15px !important;
        }
        .es-m-p20 {
          padding: 20px !important;
        }
        .es-m-p20t {
          padding-top: 20px !important;
        }
        .es-m-p20r {
          padding-right: 20px !important;
        }
        .es-m-p20l {
          padding-left: 20px !important;
        }
        .es-m-p25 {
          padding: 25px !important;
        }
        .es-m-p25t {
          padding-top: 25px !important;
        }
        .es-m-p25b {
          padding-bottom: 25px !important;
        }
        .es-m-p25r {
          padding-right: 25px !important;
        }
        .es-m-p25l {
          padding-left: 25px !important;
        }
        .es-m-p30 {
          padding: 30px !important;
        }
        .es-m-p30t {
          padding-top: 30px !important;
        }
        .es-m-p30b {
          padding-bottom: 30px !important;
        }
        .es-m-p30r {
          padding-right: 30px !important;
        }
        .es-m-p30l {
          padding-left: 30px !important;
        }
        .es-m-p35 {
          padding: 35px !important;
        }
        .es-m-p35t {
          padding-top: 35px !important;
        }
        .es-m-p35b {
          padding-bottom: 35px !important;
        }
        .es-m-p35r {
          padding-right: 35px !important;
        }
        .es-m-p35l {
          padding-left: 35px !important;
        }
        .es-m-p40 {
          padding: 40px !important;
        }
        .es-m-p40t {
          padding-top: 40px !important;
        }
        .es-m-p40b {
          padding-bottom: 40px !important;
        }
        .es-m-p40r {
          padding-right: 40px !important;
        }
        .es-m-p40l {
          padding-left: 40px !important;
        }
        .es-desk-hidden {
          display: table-row !important;
          width: auto !important;
          overflow: visible !important;
          max-height: inherit !important;
        }
      }
      @media screen and (max-width: 384px) {
        .mail-message-content {
          width: 414px !important;
        }
      }
      * {
        scrollbar-width: thin;
        scrollbar-color: #888 #f6f6f6;
      } /* Chrome, Edge, Safari */
      ::-webkit-scrollbar {
        width: 10px;
        height: 10px;
      }
      ::-webkit-scrollbar-track {
        background: #f6f6f6;
      }
      ::-webkit-scrollbar-thumb {
        background: #888;
        border-radius: 6px;
        border: 2px solid #f6f6f6;
      }
      ::-webkit-scrollbar-thumb:hover {
        box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.3);
      }
      textarea::-webkit-scrollbar-track {
        margin: 15px;
      }
      #outlook a {
        padding: 0;
      }
      .es-button {
        mso-style-priority: 100 !important;
        text-decoration: none !important;
      }
      a[x-apple-data-detectors] {
        color: inherit !important;
        text-decoration: none !important;
        font-size: inherit !important;
        font-family: inherit !important;
        font-weight: inherit !important;
        line-height: inherit !important;
      }
      .es-desk-hidden {
        display: none;
        float: left;
        overflow: hidden;
        width: 0;
        max-height: 0;
        line-height: 0;
        mso-hide: all;
      }
      @media only screen and (max-width: 600px) {
        p,
        ul li,
        ol li,
        a {
          line-height: 150% !important;
        }
        h1,
        h2,
        h3,
        h1 a,
        h2 a,
        h3 a {
          line-height: 120%;
        }
        h1 {
          font-size: 40px !important;
          text-align: center;
        }
        h2 {
          font-size: 28px !important;
          text-align: center;
        }
        h3 {
          font-size: 20px !important;
          text-align: left;
        }
        .es-header-body h1 a,
        .es-content-body h1 a,
        .es-footer-body h1 a {
          font-size: 40px !important;
          text-align: center;
        }
        .es-header-body h2 a,
        .es-content-body h2 a,
        .es-footer-body h2 a {
          font-size: 28px !important;
          text-align: center;
        }
        .es-header-body h3 a,
        .es-content-body h3 a,
        .es-footer-body h3 a {
          font-size: 20px !important;
          text-align: left;
        }
        .es-menu td a {
          font-size: 14px !important;
        }
        .es-header-body p,
        .es-header-body ul li,
        .es-header-body ol li,
        .es-header-body a {
          font-size: 14px !important;
        }
        .es-content-body p,
        .es-content-body ul li,
        .es-content-body ol li,
        .es-content-body a {
          font-size: 14px !important;
        }
        .es-footer-body p,
        .es-footer-body ul li,
        .es-footer-body ol li,
        .es-footer-body a {
          font-size: 14px !important;
        }
        .es-infoblock p,
        .es-infoblock ul li,
        .es-infoblock ol li,
        .es-infoblock a {
          font-size: 12px !important;
        }
        *[class="gmail-fix"] {
          display: none !important;
        }
        .es-m-txt-c,
        .es-m-txt-c h1,
        .es-m-txt-c h2,
        .es-m-txt-c h3 {
          text-align: center !important;
        }
        .es-m-txt-r,
        .es-m-txt-r h1,
        .es-m-txt-r h2,
        .es-m-txt-r h3 {
          text-align: right !important;
        }
        .es-m-txt-l,
        .es-m-txt-l h1,
        .es-m-txt-l h2,
        .es-m-txt-l h3 {
          text-align: left !important;
        }
        .es-m-txt-r img,
        .es-m-txt-c img,
        .es-m-txt-l img {
          display: inline !important;
        }
        .es-button-border {
          display: inline-block !important;
        }
        a.es-button,
        button.es-button {
          font-size: 18px !important;
          display: inline-block !important;
        }
        .es-adaptive table,
        .es-left,
        .es-right {
          width: 100% !important;
        }
        .es-content table,
        .es-header table,
        .es-footer table,
        .es-content,
        .es-footer,
        .es-header {
          width: 100% !important;
          max-width: 600px !important;
        }
        .es-adapt-td {
          display: block !important;
          width: 100% !important;
        }
        .adapt-img {
          width: 100% !important;
          height: auto !important;
        }
        .es-m-p0 {
          padding: 0 !important;
        }
        .es-m-p0r {
          padding-right: 0 !important;
        }
        .es-m-p0l {
          padding-left: 0 !important;
        }
        .es-m-p0t {
          padding-top: 0 !important;
        }
        .es-m-p0b {
          padding-bottom: 0 !important;
        }
        .es-m-p20b {
          padding-bottom: 20px !important;
        }
        .es-mobile-hidden,
        .es-hidden {
          display: none !important;
        }
        tr.es-desk-hidden,
        td.es-desk-hidden,
        table.es-desk-hidden {
          width: auto !important;
          overflow: visible !important;
          float: none !important;
          max-height: inherit !important;
          line-height: inherit !important;
        }
        tr.es-desk-hidden {
          display: table-row !important;
        }
        table.es-desk-hidden {
          display: table !important;
        }
        td.es-desk-menu-hidden {
          display: table-cell !important;
        }
        .es-menu td {
          width: 1% !important;
        }
        table.es-table-not-adapt,
        .esd-block-html table {
          width: auto !important;
        }
        table.es-social {
          display: inline-block !important;
        }
        table.es-social td {
          display: inline-block !important;
        }
        .es-m-p5 {
          padding: 5px !important;
        }
        .es-m-p5t {
          padding-top: 5px !important;
        }
        .es-m-p5b {
          padding-bottom: 5px !important;
        }
        .es-m-p5r {
          padding-right: 5px !important;
        }
        .es-m-p5l {
          padding-left: 5px !important;
        }
        .es-m-p10 {
          padding: 10px !important;
        }
        .es-m-p10t {
          padding-top: 10px !important;
        }
        .es-m-p10b {
          padding-bottom: 10px !important;
        }
        .es-m-p10r {
          padding-right: 10px !important;
        }
        .es-m-p10l {
          padding-left: 10px !important;
        }
        .es-m-p15 {
          padding: 15px !important;
        }
        .es-m-p15t {
          padding-top: 15px !important;
        }
        .es-m-p15b {
          padding-bottom: 15px !important;
        }
        .es-m-p15r {
          padding-right: 15px !important;
        }
        .es-m-p15l {
          padding-left: 15px !important;
        }
        .es-m-p20 {
          padding: 20px !important;
        }
        .es-m-p20t {
          padding-top: 20px !important;
        }
        .es-m-p20r {
          padding-right: 20px !important;
        }
        .es-m-p20l {
          padding-left: 20px !important;
        }
        .es-m-p25 {
          padding: 25px !important;
        }
        .es-m-p25t {
          padding-top: 25px !important;
        }
        .es-m-p25b {
          padding-bottom: 25px !important;
        }
        .es-m-p25r {
          padding-right: 25px !important;
        }
        .es-m-p25l {
          padding-left: 25px !important;
        }
        .es-m-p30 {
          padding: 30px !important;
        }
        .es-m-p30t {
          padding-top: 30px !important;
        }
        .es-m-p30b {
          padding-bottom: 30px !important;
        }
        .es-m-p30r {
          padding-right: 30px !important;
        }
        .es-m-p30l {
          padding-left: 30px !important;
        }
        .es-m-p35 {
          padding: 35px !important;
        }
        .es-m-p35t {
          padding-top: 35px !important;
        }
        .es-m-p35b {
          padding-bottom: 35px !important;
        }
        .es-m-p35r {
          padding-right: 35px !important;
        }
        .es-m-p35l {
          padding-left: 35px !important;
        }
        .es-m-p40 {
          padding: 40px !important;
        }
        .es-m-p40t {
          padding-top: 40px !important;
        }
        .es-m-p40b {
          padding-bottom: 40px !important;
        }
        .es-m-p40r {
          padding-right: 40px !important;
        }
        .es-m-p40l {
          padding-left: 40px !important;
        }
        .es-desk-hidden {
          display: table-row !important;
          width: auto !important;
          overflow: visible !important;
          max-height: inherit !important;
        }
      }
      @media screen and (max-width: 384px) {
        .mail-message-content {
          width: 414px !important;
        }
      }
    </style>
  </head>
  <body
    style="
      width: 100%;
      font-family: 'trebuchet ms', 'lucida grande', 'lucida sans unicode',
        'lucida sans', tahoma, sans-serif;
      -webkit-text-size-adjust: 100%;
      -ms-text-size-adjust: 100%;
      padding: 0;
      margin: 0;
    "
    data-new-gr-c-s-check-loaded="14.1167.0"
    data-gr-ext-installed
    data-new-gr-c-s-loaded="14.1167.0"
  >
    <div
      dir="ltr"
      class="es-wrapper-color"
      lang="und"
      style="background-color: #41b9ef"
    >
      <!--[if gte mso 9
        ]><v:background xmlns:v="urn:schemas-microsoft-com:vml" fill="t">
          <v:fill type="tile" color="#41B9EF"></v:fill> </v:background
      ><![endif]-->
      <table
        class="es-wrapper"
        width="100%"
        cellspacing="0"
        cellpadding="0"
        style="
          mso-table-lspace: 0pt;
          mso-table-rspace: 0pt;
          border-collapse: collapse;
          border-spacing: 0px;
          padding: 0;
          margin: 0;
          width: 100%;
          height: 100%;
          background-repeat: repeat;
          background-position: center top;
          background-color: #41b9ef;
        "
        role="none"
      >
        <tr>
          <td valign="top" style="padding: 0; margin: 0">
            <table
              cellpadding="0"
              cellspacing="0"
              class="es-header"
              align="center"
              role="none"
              style="
                mso-table-lspace: 0pt;
                mso-table-rspace: 0pt;
                border-collapse: collapse;
                border-spacing: 0px;
                table-layout: fixed !important;
                width: 100%;
                background-color: transparent;
                background-repeat: repeat;
                background-position: center top;
              "
            >
              <tr>
                <td align="center" style="padding: 0; margin: 0">
                  <table
                    class="es-header-body"
                    align="center"
                    cellpadding="0"
                    cellspacing="0"
                    style="
                      mso-table-lspace: 0pt;
                      mso-table-rspace: 0pt;
                      border-collapse: collapse;
                      border-spacing: 0px;
                      background-color: #1d9ad1;
                      border-radius: 10px 10px 0px 0px;
                      width: 560px;
                    "
                    role="none"
                  >
                    <tr>
                      <td align="left" style="padding: 20px; margin: 0">
                        <!--[if mso]><table style="width:520px" cellpadding="0" cellspacing="0"><tr>
<td style="width:250px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          align="left"
                          class="es-left"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: left;
                          "
                        >
                          <tr>
                            <td
                              class="es-m-p0r es-m-p20b"
                              valign="top"
                              align="center"
                              style="padding: 0; margin: 0; width: 250px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="left"
                                    class="es-m-txt-c"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      font-size: 0px;
                                    "
                                  >
                                    <a
                                      target="_blank"
                                      href="https://traveloment.com/"
                                      style="
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        text-decoration: none;
                                        color: #ffffff;
                                        font-size: 14px;
                                      "
                                      ><img
                                        src="https://traveloment.com/wp-content/uploads/2023/08/trav_logo-removebg-preview-768x213.png"
                                        alt="Logo"
                                        style="
                                          display: block;
                                          border: 0;
                                          outline: none;
                                          text-decoration: none;
                                          -ms-interpolation-mode: bicubic;
                                        "
                                        title="Logo"
                                        width="200"
                                        height="54"
                                    /></a>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                        <!--[if mso]></td><td style="width:20px"></td><td style="width:250px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-right"
                          align="right"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: right;
                          "
                        >
                          <tr>
                            <td
                              align="left"
                              style="padding: 0; margin: 0; width: 250px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="right"
                                    class="es-m-txt-c"
                                    style="padding: 0; margin: 0"
                                  >
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 21px;
                                        color: #ffffff;
                                        font-size: 14px;
                                      "
                                    >
                                      {{.BookingDate}}
                                    </p>
                                    <h3
                                      style="
                                        margin: 0;
                                        line-height: 24px;
                                        mso-line-height-rule: exactly;
                                        font-family: arial, 'helvetica neue',
                                          helvetica, sans-serif;
                                        font-size: 20px;
                                        font-style: normal;
                                        font-weight: bold;
                                        color: #ffffff;
                                      "
                                    >
                                      Booking confirmation
                                    </h3>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                        <!--[if mso]></td></tr></table><![endif]-->
                      </td>
                    </tr>
                    <tr>
                      <td
                        align="left"
                        style="
                          padding: 0;
                          margin: 0;
                          padding-top: 5px;
                          padding-bottom: 5px;
                        "
                      >
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          width="100%"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                          "
                        >
                          <tr>
                            <td
                              align="center"
                              valign="top"
                              style="padding: 0; margin: 0; width: 560px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="center"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      font-size: 0px;
                                    "
                                  >
                                    <img
                                      class="adapt-img"
                                      src="https://fifyhaq.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group_37_ikx.png"
                                      alt
                                      style="
                                        display: block;
                                        border: 0;
                                        outline: none;
                                        text-decoration: none;
                                        -ms-interpolation-mode: bicubic;
                                      "
                                      width="560"
                                      height="4"
                                    />
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
                  </table>
                </td>
              </tr>
            </table>
            <table
              cellpadding="0"
              cellspacing="0"
              class="es-content"
              align="center"
              role="none"
              style="
                mso-table-lspace: 0pt;
                mso-table-rspace: 0pt;
                border-collapse: collapse;
                border-spacing: 0px;
                table-layout: fixed !important;
                width: 100%;
              "
            >
              <tr>
                <td align="center" style="padding: 0; margin: 0">
                  <table
                    bgcolor="#ffffff"
                    class="es-content-body"
                    align="center"
                    cellpadding="0"
                    cellspacing="0"
                    style="
                      mso-table-lspace: 0pt;
                      mso-table-rspace: 0pt;
                      border-collapse: collapse;
                      border-spacing: 0px;
                      background-color: #1d9ad1;
                      border-radius: 0 0 10px 10px;
                      width: 560px;
                    "
                    role="none"
                  >
                    <tr>
                      <td
                        align="left"
                        style="
                          padding: 0;
                          margin: 0;
                          padding-top: 20px;
                          padding-left: 20px;
                          padding-bottom: 30px;
                        "
                      >
                        <!--[if mso]><table style="width:540px" cellpadding="0" cellspacing="0"><tr>
<td style="width:312px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          align="left"
                          class="es-left"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: left;
                          "
                        >
                          <tr>
                            <td
                              class="es-m-p0r es-m-p20b"
                              valign="top"
                              align="center"
                              style="padding: 0; margin: 0; width: 312px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="left"
                                    class="es-m-p20r es-m-txt-l"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-bottom: 10px;
                                      padding-top: 20px;
                                    "
                                  >
                                    <h1
                                      style="
                                        margin: 0;
                                        line-height: 48px;
                                        mso-line-height-rule: exactly;
                                        font-family: arial, 'helvetica neue',
                                          helvetica, sans-serif;
                                        font-size: 40px;
                                        font-style: normal;
                                        font-weight: bold;
                                        color: #ffffff;
                                      "
                                    >
                                      Hi, your booking is confirmed!
                                    </h1>
                                  </td>
                                </tr>
                                <tr>
                                  <td
                                    align="left"
                                    class="es-m-p20r"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-top: 10px;
                                      padding-bottom: 10px;
                                      padding-right: 10px;
                                    "
                                  >
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 24px;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                    >
                                      Please find your booking code below.
                                      <br />We look forward to seeing you soon!
                                    </p>
                                  </td>
                                </tr>
                                <tr>
                                  <td
                                    align="left"
                                    class="es-m-p20r"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-top: 10px;
                                      padding-bottom: 10px;
                                      padding-right: 10px;
                                    "
                                  >
                                    <h3
                                      style="
                                        margin: 0;
                                        line-height: 24px;
                                        mso-line-height-rule: exactly;
                                        font-family: arial, 'helvetica neue',
                                          helvetica, sans-serif;
                                        font-size: 20px;
                                        font-style: normal;
                                        font-weight: bold;
                                        color: #f76b0a;
                                      "
                                    >
                                      <span
                                        style="
                                          background-color: #ffffff;
                                          border-radius: 10px;
                                        "
                                        >&nbsp;Booking No. {{.BookingNo}}&nbsp;</span
                                      >
                                    </h3>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                        <!--[if mso]></td>
<td style="width:20px"></td><td style="width:208px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-right"
                          align="right"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: right;
                          "
                        >
                          <tr>
                            <td
                              align="left"
                              style="padding: 0; margin: 0; width: 208px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="center"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      font-size: 0px;
                                    "
                                  >
                                    <a
                                      target="_blank"
                                      href="https://traveloment.com/"
                                      style="
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        text-decoration: none;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                      ><img
                                        class="adapt-img"
                                        src="https://fifyhaq.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/5883100.png"
                                        alt
                                        style="
                                          display: block;
                                          border: 0;
                                          outline: none;
                                          text-decoration: none;
                                          -ms-interpolation-mode: bicubic;
                                        "
                                        width="208"
                                        height="309"
                                    /></a>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                        <!--[if mso]></td></tr></table><![endif]-->
                      </td>
                    </tr>
                  </table>
                </td>
              </tr>
            </table>
            <table
              cellpadding="0"
              cellspacing="0"
              class="es-content"
              align="center"
              role="none"
              style="
                mso-table-lspace: 0pt;
                mso-table-rspace: 0pt;
                border-collapse: collapse;
                border-spacing: 0px;
                table-layout: fixed !important;
                width: 100%;
              "
            >
              <tr>
                <td align="center" style="padding: 0; margin: 0">
                  <table
                    class="es-content-body"
                    align="center"
                    cellpadding="0"
                    cellspacing="0"
                    style="
                      mso-table-lspace: 0pt;
                      mso-table-rspace: 0pt;
                      border-collapse: collapse;
                      border-spacing: 0px;
                      background-color: transparent;
                      width: 560px;
                    "
                    role="none"
                    bgcolor="#ffffff"
                  >
                    <tr>
                      <td
                        align="left"
                        style="
                          margin: 0;
                          padding-bottom: 20px;
                          padding-left: 20px;
                          padding-right: 20px;
                          padding-top: 30px;
                        "
                      >
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          width="100%"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                          "
                        >
                          <tr>
                            <td
                              align="center"
                              valign="top"
                              style="padding: 0; margin: 0; width: 520px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="left"
                                    style="padding: 0; margin: 0"
                                  >
                                    <h2
                                      style="
                                        margin: 0;
                                        line-height: 34px;
                                        mso-line-height-rule: exactly;
                                        font-family: arial, 'helvetica neue',
                                          helvetica, sans-serif;
                                        font-size: 28px;
                                        font-style: normal;
                                        font-weight: bold;
                                        color: #ffffff;
                                      "
                                    >
                                      Booking from
                                      <span
                                        style="
                                          background-color: #ffffff;
                                          color: #f76b0a;
                                          border-radius: 1rem;
                                          padding-left: 1rem;
                                          padding-right: 1rem;
                                        "
                                        >{{.AirportName}}</span
                                      >
                                    </h2>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
                  </table>
                </td>
              </tr>
            </table>
            <table
              cellpadding="0"
              cellspacing="0"
              class="es-content"
              align="center"
              role="none"
              style="
                mso-table-lspace: 0pt;
                mso-table-rspace: 0pt;
                border-collapse: collapse;
                border-spacing: 0px;
                table-layout: fixed !important;
                width: 100%;
              "
            >
              <tr>
                <td align="center" style="padding: 0; margin: 0">
                  <table
                    class="es-content-body"
                    align="center"
                    cellpadding="0"
                    cellspacing="0"
                    style="
                      mso-table-lspace: 0pt;
                      mso-table-rspace: 0pt;
                      border-collapse: collapse;
                      border-spacing: 0px;
                      background-color: transparent;
                      width: 560px;
                    "
                    role="none"
                    bgcolor="#ffffff"
                  >
                    <tr>
                      <td
                        align="left"
                        style="
                          margin: 0;
                          padding-bottom: 10px;
                          padding-top: 20px;
                          padding-left: 20px;
                          padding-right: 20px;
                        "
                      >
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          width="100%"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                          "
                        >
                          <tr>
                            <td
                              align="center"
                              valign="top"
                              style="padding: 0; margin: 0; width: 520px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="left"
                                    style="padding: 0; margin: 0"
                                  >
                                    <h3
                                      style="
                                        margin: 0;
                                        line-height: 24px;
                                        mso-line-height-rule: exactly;
                                        font-family: arial, 'helvetica neue',
                                          helvetica, sans-serif;
                                        font-size: 20px;
                                        font-style: normal;
                                        font-weight: bold;
                                        color: #ffffff;
                                      "
                                    >
                                      {{.Pickupinfo.PickUpDay}}, {{.Pickupinfo.PickUpDate}}
                                    </h3>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
                    <tr>
                      <td
                        align="left"
                        bgcolor="#ffffff"
                        style="
                          padding: 20px;
                          margin: 0;
                          background-color: #ffffff;
                          border-radius: 10px;
                        "
                      >
                        <!--[if mso]><table style="width:520px" cellpadding="0" cellspacing="0"><tr><td style="width:160px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-left"
                          align="left"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: left;
                          "
                        >
                          <tr>
                            <td
                              class="es-m-p0r es-m-p20b"
                              align="center"
                              style="padding: 0; margin: 0; width: 150px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="center"
                                    style="padding: 0; margin: 0"
                                  >
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 21px;
                                        color: #3e7d8e;
                                        font-size: 14px;
                                      "
                                    >
                                      by {{.Pickupinfo.AirlineName}}
                                    </p>
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 21px;
                                        color: #3e7d8e;
                                        font-size: 14px;
                                      "
                                    >
                                      Departs at
                                    </p>
                                  </td>
                                </tr>
                                <tr>
                                  <td
                                    align="center"
                                    style="padding: 0; margin: 0"
                                  >
                                    <h2
                                      style="
                                        margin: 0;
                                        line-height: 42px;
                                        mso-line-height-rule: exactly;
                                        font-family: arial, 'helvetica neue',
                                          helvetica, sans-serif;
                                        font-size: 28px;
                                        font-style: normal;
                                        font-weight: bold;
                                        color: #144d62;
                                      "
                                    >
                                      <strong>{{.Pickupinfo.DepartsTiming}}
                                        </strong>
                                    </h2>
                                  </td>
                                </tr>
                                <tr>
                                  <td
                                    align="center"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-bottom: 5px;
                                    "
                                  >
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 24px;
                                        color: #3e7d8e;
                                        font-size: 16px;
                                      "
                                    >
                                      <strong>{{.Pickupinfo.DepartsLocation}}</strong>
                                    </p>
                                  </td>
                                </tr>
                              </table>
                            </td>
                            <td
                              class="es-hidden"
                              style="padding: 0; margin: 0; width: 10px"
                            ></td>
                          </tr>
                        </table>
                        <!--[if mso]></td>
<td style="width:40px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-left"
                          align="left"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: left;
                          "
                        >
                          <tr class="es-mobile-hidden">
                            <td
                              class="es-m-p20b"
                              align="center"
                              style="padding: 0; margin: 0; width: 30px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="center"
                                    class="es-m-p0"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-bottom: 35px;
                                      padding-top: 40px;
                                      font-size: 0px;
                                    "
                                  >
                                    <a
                                      target="_blank"
                                      href="https://traveloment.com/"
                                      style="
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        text-decoration: none;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                      ><img
                                        src="https://fifyhaq.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group.png"
                                        alt
                                        style="
                                          display: block;
                                          border: 0;
                                          outline: none;
                                          text-decoration: none;
                                          -ms-interpolation-mode: bicubic;
                                        "
                                        width="25"
                                        height="16"
                                    /></a>
                                  </td>
                                </tr>
                              </table>
                            </td>
                            <td
                              class="es-hidden"
                              style="padding: 0; margin: 0; width: 10px"
                            ></td>
                          </tr>
                        </table>
                        <!--[if mso]></td><td style="width:130px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-left"
                          align="left"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: left;
                          "
                        >
                          <tr>
                            <td
                              class="es-m-p20b"
                              align="center"
                              style="padding: 0; margin: 0; width: 120px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="center"
                                    class="es-m-p0t"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-top: 40px;
                                    "
                                  >
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 24px;
                                        color: #3e7d8e;
                                        font-size: 16px;
                                      "
                                    >
                                      <strong>{{.Pickupinfo.FlightDuration}}</strong>
                                    </p>
                                  </td>
                                </tr>
                              </table>
                            </td>
                            <td
                              class="es-hidden"
                              style="padding: 0; margin: 0; width: 10px"
                            ></td>
                          </tr>
                        </table>
                        <!--[if mso]></td>
<td style="width:30px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-left"
                          align="left"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: left;
                          "
                        >
                          <tr class="es-mobile-hidden">
                            <td
                              class="es-m-p20b"
                              align="center"
                              style="padding: 0; margin: 0; width: 30px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="center"
                                    class="es-m-p0t es-m-p0b"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-bottom: 35px;
                                      padding-top: 40px;
                                      font-size: 0px;
                                    "
                                  >
                                    <a
                                      target="_blank"
                                      href="https://traveloment.com/"
                                      style="
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        text-decoration: none;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                      ><img
                                        src="https://fifyhaq.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group1.png"
                                        alt
                                        style="
                                          display: block;
                                          border: 0;
                                          outline: none;
                                          text-decoration: none;
                                          -ms-interpolation-mode: bicubic;
                                        "
                                        width="25"
                                        height="18"
                                    /></a>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                        <!--[if mso]></td><td style="width:10px"></td><td style="width:150px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-right"
                          align="right"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: right;
                          "
                        >
                          <tr>
                            <td
                              class="es-m-p0r"
                              align="center"
                              style="padding: 0; margin: 0; width: 150px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="center"
                                    style="padding: 0; margin: 0"
                                  >
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 21px;
                                        color: #3e7d8e;
                                        font-size: 14px;
                                      "
                                    >
                                      by {{.Pickupinfo.AirlineName}}
                                    </p>
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 21px;
                                        color: #3e7d8e;
                                        font-size: 14px;
                                      "
                                    >
                                      Arrives at<br />
                                    </p>
                                  </td>
                                </tr>
                                <tr>
                                  <td
                                    align="center"
                                    style="padding: 0; margin: 0"
                                  >
                                    <h2
                                      style="
                                        margin: 0;
                                        line-height: 42px;
                                        mso-line-height-rule: exactly;
                                        font-family: arial, 'helvetica neue',
                                          helvetica, sans-serif;
                                        font-size: 28px;
                                        font-style: normal;
                                        font-weight: bold;
                                        color: #144d62;
                                      "
                                    >
                                      <strong
                                        >{{.Pickupinfo.ArrivalTiming}}</strong
                                      >
                                    </h2>
                                  </td>
                                </tr>
                                <tr>
                                  <td
                                    align="center"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-bottom: 5px;
                                    "
                                  >
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 24px;
                                        color: #3e7d8e;
                                        font-size: 16px;
                                      "
                                    >
                                      <strong>{{.Pickupinfo.ArrivalLocation}}</strong>
                                    </p>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                        <!--[if mso]></td></tr></table><![endif]-->
                      </td>
                    </tr>
                  </table>
                </td>
              </tr>
            </table>
            <table
              cellpadding="0"
              cellspacing="0"
              class="es-content"
              align="center"
              role="none"
              style="
                mso-table-lspace: 0pt;
                mso-table-rspace: 0pt;
                border-collapse: collapse;
                border-spacing: 0px;
                table-layout: fixed !important;
                width: 100%;
              "
            >
              <tr>
                <td align="center" style="padding: 0; margin: 0">
                  <table
                    class="es-content-body"
                    align="center"
                    cellpadding="0"
                    cellspacing="0"
                    style="
                      mso-table-lspace: 0pt;
                      mso-table-rspace: 0pt;
                      border-collapse: collapse;
                      border-spacing: 0px;
                      background-color: transparent;
                      width: 560px;
                    "
                    role="none"
                    bgcolor="#ffffff"
                  >
                    <tr>
                      <td
                        align="left"
                        style="
                          margin: 0;
                          padding-bottom: 10px;
                          padding-left: 20px;
                          padding-right: 20px;
                          padding-top: 25px;
                        "
                      >
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          width="100%"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                          "
                        >
                          <tr>
                            <td
                              align="center"
                              valign="top"
                              style="padding: 0; margin: 0; width: 520px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="left"
                                    style="padding: 0; margin: 0"
                                  >
                                    <h3
                                      style="
                                        margin: 0;
                                        line-height: 24px;
                                        mso-line-height-rule: exactly;
                                        font-family: arial, 'helvetica neue',
                                          helvetica, sans-serif;
                                        font-size: 20px;
                                        font-style: normal;
                                        font-weight: bold;
                                        color: #ffffff;
                                      "
                                    >
                                    {{.Dropofinfo.DropOfDay}}, {{.Dropofinfo.DropOfDate}}
                                    </h3>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
                    <tr>
                      <td
                        align="left"
                        bgcolor="#ffffff"
                        style="
                          padding: 20px;
                          margin: 0;
                          background-color: #ffffff;
                          border-radius: 10px;
                        "
                      >
                        <!--[if mso]><table style="width:520px" cellpadding="0" cellspacing="0"><tr>
<td style="width:160px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-left"
                          align="left"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: left;
                          "
                        >
                          <tr>
                            <td
                              class="es-m-p0r es-m-p20b"
                              align="center"
                              style="padding: 0; margin: 0; width: 150px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="center"
                                    style="padding: 0; margin: 0"
                                  >
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 21px;
                                        color: #3e7d8e;
                                        font-size: 14px;
                                      "
                                    >
                                      by {{.Dropofinfo.AirlineName}}
                                    </p>
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 21px;
                                        color: #3e7d8e;
                                        font-size: 14px;
                                      "
                                    >
                                      Departs at
                                    </p>
                                  </td>
                                </tr>
                                <tr>
                                  <td
                                    align="center"
                                    style="padding: 0; margin: 0"
                                  >
                                    <h2
                                      style="
                                        margin: 0;
                                        line-height: 42px;
                                        mso-line-height-rule: exactly;
                                        font-family: arial, 'helvetica neue',
                                          helvetica, sans-serif;
                                        font-size: 28px;
                                        font-style: normal;
                                        font-weight: bold;
                                        color: #144d62;
                                      "
                                    >
                                      <strong
                                        >{{.Dropofinfo.DepartsTiming}}</strong
                                      >
                                    </h2>
                                  </td>
                                </tr>
                                <tr>
                                  <td
                                    align="center"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-bottom: 5px;
                                    "
                                  >
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 24px;
                                        color: #3e7d8e;
                                        font-size: 16px;
                                      "
                                    >
                                      <strong>{{.Dropofinfo.DepartsLocation}}</strong>
                                    </p>
                                  </td>
                                </tr>
                              </table>
                            </td>
                            <td
                              class="es-hidden"
                              style="padding: 0; margin: 0; width: 10px"
                            ></td>
                          </tr>
                        </table>
                        <!--[if mso]></td><td style="width:40px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-left"
                          align="left"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: left;
                          "
                        >
                          <tr class="es-mobile-hidden">
                            <td
                              class="es-m-p20b"
                              align="center"
                              style="padding: 0; margin: 0; width: 30px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="center"
                                    class="es-m-p0"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-bottom: 35px;
                                      padding-top: 40px;
                                      font-size: 0px;
                                    "
                                  >
                                    <a
                                      target="_blank"
                                      href="https://traveloment.com/"
                                      style="
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        text-decoration: none;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                      ><img
                                        src="https://fifyhaq.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group.png"
                                        alt
                                        style="
                                          display: block;
                                          border: 0;
                                          outline: none;
                                          text-decoration: none;
                                          -ms-interpolation-mode: bicubic;
                                        "
                                        width="25"
                                        height="16"
                                    /></a>
                                  </td>
                                </tr>
                              </table>
                            </td>
                            <td
                              class="es-hidden"
                              style="padding: 0; margin: 0; width: 10px"
                            ></td>
                          </tr>
                        </table>
                        <!--[if mso]></td><td style="width:130px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-left"
                          align="left"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: left;
                          "
                        >
                          <tr>
                            <td
                              class="es-m-p20b"
                              align="center"
                              style="padding: 0; margin: 0; width: 120px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="center"
                                    class="es-m-p0t"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-top: 40px;
                                    "
                                  >
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 24px;
                                        color: #3e7d8e;
                                        font-size: 16px;
                                      "
                                    >
                                      <strong>{{.Dropofinfo.FlightDuration}}</strong>
                                    </p>
                                  </td>
                                </tr>
                              </table>
                            </td>
                            <td
                              class="es-hidden"
                              style="padding: 0; margin: 0; width: 10px"
                            ></td>
                          </tr>
                        </table>
                        <!--[if mso]></td>
<td style="width:30px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-left"
                          align="left"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: left;
                          "
                        >
                          <tr class="es-mobile-hidden">
                            <td
                              class="es-m-p20b"
                              align="center"
                              style="padding: 0; margin: 0; width: 30px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="center"
                                    class="es-m-p0t es-m-p0b"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-bottom: 35px;
                                      padding-top: 40px;
                                      font-size: 0px;
                                    "
                                  >
                                    <a
                                      target="_blank"
                                      href="https://traveloment.com/"
                                      style="
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        text-decoration: none;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                      ><img
                                        src="https://fifyhaq.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group1.png"
                                        alt
                                        style="
                                          display: block;
                                          border: 0;
                                          outline: none;
                                          text-decoration: none;
                                          -ms-interpolation-mode: bicubic;
                                        "
                                        width="25"
                                        height="18"
                                    /></a>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                        <!--[if mso]></td><td style="width:10px"></td><td style="width:150px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-right"
                          align="right"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: right;
                          "
                        >
                          <tr>
                            <td
                              class="es-m-p0r"
                              align="center"
                              style="padding: 0; margin: 0; width: 150px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="center"
                                    style="padding: 0; margin: 0"
                                  >
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 21px;
                                        color: #3e7d8e;
                                        font-size: 14px;
                                      "
                                    >
                                      by {{.Dropofinfo.AirlineName}}
                                    </p>
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 21px;
                                        color: #3e7d8e;
                                        font-size: 14px;
                                      "
                                    >
                                      Arrives at<br />
                                    </p>
                                  </td>
                                </tr>
                                <tr>
                                  <td
                                    align="center"
                                    style="padding: 0; margin: 0"
                                  >
                                    <h2
                                      style="
                                        margin: 0;
                                        line-height: 42px;
                                        mso-line-height-rule: exactly;
                                        font-family: arial, 'helvetica neue',
                                          helvetica, sans-serif;
                                        font-size: 28px;
                                        font-style: normal;
                                        font-weight: bold;
                                        color: #144d62;
                                      "
                                    >
                                      <strong
                                        >{{.Dropofinfo.ArrivalTiming}}</strong
                                      >
                                    </h2>
                                  </td>
                                </tr>
                                <tr>
                                  <td
                                    align="center"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-bottom: 5px;
                                    "
                                  >
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 24px;
                                        color: #3e7d8e;
                                        font-size: 16px;
                                      "
                                    >
                                      <strong>{{.Dropofinfo.ArrivalLocation}}</strong>
                                    </p>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                        <!--[if mso]></td></tr></table><![endif]-->
                      </td>
                    </tr>
                  </table>
                </td>
              </tr>
            </table>
            <table
              cellpadding="0"
              cellspacing="0"
              class="es-content"
              align="center"
              role="none"
              style="
                mso-table-lspace: 0pt;
                mso-table-rspace: 0pt;
                border-collapse: collapse;
                border-spacing: 0px;
                table-layout: fixed !important;
                width: 100%;
              "
            >
              <tr>
                <td align="center" style="padding: 0; margin: 0">
                  <table
                    class="es-content-body"
                    align="center"
                    cellpadding="0"
                    cellspacing="0"
                    style="
                      mso-table-lspace: 0pt;
                      mso-table-rspace: 0pt;
                      border-collapse: collapse;
                      border-spacing: 0px;
                      background-color: transparent;
                      width: 560px;
                    "
                    role="none"
                    bgcolor="#ffffff"
                  >
                    <tr>
                      <td
                        align="left"
                        style="
                          padding: 0;
                          margin: 0;
                          padding-top: 20px;
                          padding-left: 20px;
                          padding-right: 20px;
                        "
                      >
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          width="100%"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                          "
                        >
                          <tr>
                            <td
                              align="center"
                              valign="top"
                              style="padding: 0; margin: 0; width: 520px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="center"
                                    style="padding: 0; margin: 0"
                                  >
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 24px;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                    >
                                      Remember to arrive at least&nbsp;<br />
                                    </p>
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 24px;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                    >
                                      <strong
                                        ><span style="color: #f76b0a"
                                          ><span
                                            style="
                                              background-color: #ffffff;
                                              border-radius: 10px;
                                            "
                                          >
                                            &nbsp;2 hours before
                                            departure&nbsp;</span
                                          ></span
                                        ></strong
                                      >&nbsp;time.
                                    </p>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
                    <tr>
                      <td
                        align="left"
                        style="
                          padding: 0;
                          margin: 0;
                      
                          padding-left: 20px;
                          padding-right: 20px;
                        "
                      >
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          width="100%"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                          "
                        >
                          <tr>
                            <td
                              align="center"
                              valign="top"
                              style="padding: 0; margin: 0; width: 520px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                             
                                <tr>
                                  <td
                                    align="center"
                                    height="40"
                                    style="padding: 0; margin: 0"
                                  ></td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
                  </table>
                </td>
              </tr>
            </table>
            <table
              cellpadding="0"
              cellspacing="0"
              class="es-content"
              align="center"
              role="none"
              style="
                mso-table-lspace: 0pt;
                mso-table-rspace: 0pt;
                border-collapse: collapse;
                border-spacing: 0px;
                table-layout: fixed !important;
                width: 100%;
              "
            >
              <tr>
                <td align="center" style="padding: 0; margin: 0">
                  <table
                    bgcolor="#ffffff"
                    class="es-content-body"
                    align="center"
                    cellpadding="0"
                    cellspacing="0"
                    role="none"
                    style="
                      mso-table-lspace: 0pt;
                      mso-table-rspace: 0pt;
                      border-collapse: collapse;
                      border-spacing: 0px;
                      background-color: #1d9ad1;
                      width: 560px;
                    "
                  >
                    <tr>
                      <td
                        align="left"
                        style="
                          padding: 0;
                          margin: 0;
                          padding-left: 20px;
                          padding-right: 20px;
                          padding-top: 30px;
                        "
                      >
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          width="100%"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                          "
                        >
                          <tr>
                            <td
                              align="center"
                              valign="top"
                              style="padding: 0; margin: 0; width: 520px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="left"
                                    style="padding: 0; margin: 0"
                                  >
                                    <h2
                                      style="
                                        margin: 0;
                                        line-height: 34px;
                                        mso-line-height-rule: exactly;
                                        font-family: arial, 'helvetica neue',
                                          helvetica, sans-serif;
                                        font-size: 28px;
                                        font-style: normal;
                                        font-weight: bold;
                                        color: #ffffff;
                                      "
                                    >
                                      Booking details
                                    </h2>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
                    <tr>
                      <td
                        align="left"
                        style="
                          margin: 0;
                          padding-bottom: 10px;
                          padding-top: 20px;
                          padding-left: 20px;
                          padding-right: 20px;
                        "
                      >
                        <!--[if mso]><table style="width:520px" cellpadding="0" cellspacing="0"><tr>
<td style="width:253px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-left"
                          align="left"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: left;
                          "
                        >
                          <tr>
                            <td
                              class="es-m-p20b"
                              align="left"
                              style="padding: 0; margin: 0; width: 253px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                bgcolor="#41b9ef"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: separate;
                                  border-spacing: 0px;
                                  background-color: #41b9ef;
                                  border-radius: 10px;
                                "
                                role="presentation"
                              >
                                <tr>
                                  <td
                                    align="center"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      font-size: 0px;
                                    "
                                  >
                                    <a
                                      target="_blank"
                                      href="https://traveloment.com/"
                                      style="
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        text-decoration: none;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                      ><img
                                        class="adapt-img"
                                        src="https://fifyhaq.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/18.png"
                                        alt
                                        style="
                                          display: block;
                                          border: 0;
                                          outline: none;
                                          text-decoration: none;
                                          -ms-interpolation-mode: bicubic;
                                          border-radius: 10px;
                                        "
                                        width="253"
                                        height="190"
                                    /></a>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                        <!--[if mso]></td><td style="width:15px"></td><td style="width:252px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-right"
                          align="right"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: right;
                          "
                        >
                          <tr>
                            <td
                              align="left"
                              style="padding: 0; margin: 0; width: 252px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="left"
                                    class="es-m-p10t"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-bottom: 10px;
                                      padding-top: 30px;
                                    "
                                  >
                                    <h3
                                      style="
                                        margin: 0;
                                        line-height: 24px;
                                        mso-line-height-rule: exactly;
                                        font-family: arial, 'helvetica neue',
                                          helvetica, sans-serif;
                                        font-size: 20px;
                                        font-style: normal;
                                        font-weight: bold;
                                        color: #ffffff;
                                      "
                                    >
                                      {{.TripPickupLocation}}&nbsp;— {{.TripDropofLocation}}
                                    </h3>
                                  </td>
                                </tr>
                                <tr>
                                  <td
                                    align="left"
                                    class="es-m-p20r"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-top: 10px;
                                      padding-bottom: 10px;
                                      padding-right: 10px;
                                    "
                                  >
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 19px;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                    >
                                      {{.TripType}}
                                    </p>
                                  </td>
                                </tr>
                                <tr>
                                  <td
                                    align="left"
                                    class="es-m-p20r"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-top: 10px;
                                      padding-bottom: 10px;
                                      padding-right: 10px;
                                    "
                                  >
                                    <h3
                                      style="
                                        margin: 0;
                                        line-height: 24px;
                                        mso-line-height-rule: exactly;
                                        font-family: arial, 'helvetica neue',
                                          helvetica, sans-serif;
                                        font-size: 20px;
                                        font-style: normal;
                                        font-weight: bold;
                                        color: #f76b0a;
                                      "
                                    >
                                      <span
                                        style="
                                          background-color: #ffffff;
                                          border-radius: 10px;
                                        "
                                        >&nbsp;Booking No. {{.BookingNo}}&nbsp;</span
                                      >
                                    </h3>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                        <!--[if mso]></td></tr></table><![endif]-->
                      </td>
                    </tr>
                    <tr>
                      <td
                        align="left"
                        style="
                          padding: 0;
                          margin: 0;
                          padding-top: 20px;
                          padding-left: 20px;
                          padding-right: 20px;
                        "
                      >
                        <!--[if mso]><table style="width:520px" cellpadding="0" cellspacing="0"><tr><td style="width:60px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-left es-table-not-adapt"
                          align="left"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: left;
                          "
                        >
                          <tr>
                            <td
                              class="es-m-p20b"
                              align="left"
                              style="
                                padding: 0;
                                margin: 0;
                                border-radius: 10px;
                                overflow: hidden;
                                width: 60px;
                              "
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                bgcolor="#41b9ef"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: separate;
                                  border-spacing: 0px;
                                  border-left: 2px solid #6abde1;
                                  border-right: 2px solid #108ec5;
                                  border-top: 2px solid #6abde1;
                                  border-bottom: 2px solid #108ec5;
                                  background-color: #41b9ef;
                                  border-radius: 10px;
                                "
                                role="presentation"
                              >
                                <tr>
                                  <td
                                    align="center"
                                    style="
                                      padding: 10px;
                                      margin: 0;
                                      font-size: 0px;
                                    "
                                  >
                                    <a
                                      target="_blank"
                                      href="https://traveloment.com/"
                                      style="
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        text-decoration: none;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                      ><img
                                        src="https://fifyhaq.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/userverificationinterfacesymbol.png"
                                        alt
                                        style="
                                          display: block;
                                          border: 0;
                                          outline: none;
                                          text-decoration: none;
                                          -ms-interpolation-mode: bicubic;
                                          border-radius: 10px;
                                        "
                                        width="36"
                                        height="36"
                                    /></a>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                        <!--[if mso]></td><td style="width:15px"></td><td style="width:445px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-right"
                          align="right"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: right;
                          "
                        >
                          <tr>
                            <td
                              align="left"
                              style="padding: 0; margin: 0; width: 445px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="left"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-bottom: 10px;
                                    "
                                  >
                                    <h3
                                      style="
                                        margin: 0;
                                        line-height: 24px;
                                        mso-line-height-rule: exactly;
                                        font-family: arial, 'helvetica neue',
                                          helvetica, sans-serif;
                                        font-size: 20px;
                                        font-style: normal;
                                        font-weight: bold;
                                        color: #ffffff;
                                      "
                                    >
                                      Passengers
                                    </h3>
                                  </td>
                                </tr>
                                <tr>
                                  <td style="padding: 0; margin: 0">
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 24px;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                    ></p>
                                    <table
                                      width="100%"
                                      border="0"
                                      cellpadding="0"
                                      cellspacing="0"
                                      role="presentation"
                                      style="
                                        mso-table-lspace: 0pt;
                                        mso-table-rspace: 0pt;
                                        border-collapse: collapse;
                                        border-spacing: 0px;
                                      "
                                    >
                                      <tr>
                                        <td
                                          align="left"
                                          width="50%"
                                          valign="middle"
                                         
                                        >
                                        {{range .PassengerName}}
                                          <p
                                            style="
                                              margin: 0;
                                              -webkit-text-size-adjust: none;
                                              -ms-text-size-adjust: none;
                                              mso-line-height-rule: exactly;
                                              font-family: 'trebuchet ms',
                                                'lucida grande',
                                                'lucida sans unicode',
                                                'lucida sans', tahoma,
                                                sans-serif;
                                              line-height: 21px;
                                              color: #ffffff;
                                              font-size: 14px;
                                            "
                                          >
                                          
                                          
                                          {{.}}
                                          
                                          
                                          </p>
                                          {{end}}
                                        </td>
                                     
                                      </tr>
                                    </table>
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 24px;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                    ></p>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                        <!--[if mso]></td></tr></table><![endif]-->
                      </td>
                    </tr>
                    <tr>
                      <td
                        align="left"
                        style="
                          padding: 0;
                          margin: 0;
                          padding-top: 20px;
                          padding-left: 20px;
                          padding-right: 20px;
                        "
                      >
                        <!--[if mso]><table style="width:520px" cellpadding="0" cellspacing="0"><tr><td style="width:60px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-left es-table-not-adapt"
                          align="left"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: left;
                          "
                        >
                          <tr>
                            <td
                              class="es-m-p20b"
                              align="left"
                              style="
                                padding: 0;
                                margin: 0;
                                border-radius: 10px;
                                overflow: hidden;
                                width: 60px;
                              "
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                bgcolor="#41b9ef"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: separate;
                                  border-spacing: 0px;
                                  border-left: 2px solid #6abde1;
                                  border-right: 2px solid #108ec5;
                                  border-top: 2px solid #6abde1;
                                  border-bottom: 2px solid #108ec5;
                                  background-color: #41b9ef;
                                  border-radius: 10px;
                                "
                                role="presentation"
                              >
                                <tr>
                                  <td
                                    align="center"
                                    style="
                                      padding: 10px;
                                      margin: 0;
                                      font-size: 0px;
                                    "
                                  >
                                    <a
                                      target="_blank"
                                      href="https://traveloment.com/"
                                      style="
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        text-decoration: none;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                      ><img
                                        src="https://fifyhaq.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/userverificationinterfacesymbol_2.png"
                                        alt
                                        style="
                                          display: block;
                                          border: 0;
                                          outline: none;
                                          text-decoration: none;
                                          -ms-interpolation-mode: bicubic;
                                          border-radius: 10px;
                                        "
                                        width="36"
                                        height="36"
                                    /></a>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                        <!--[if mso]></td><td style="width:15px"></td><td style="width:445px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-right"
                          align="right"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: right;
                          "
                        >
                          <tr>
                            <td
                              align="left"
                              style="padding: 0; margin: 0; width: 445px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="left"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-bottom: 10px;
                                    "
                                  >
                                    <h3
                                      style="
                                        margin: 0;
                                        line-height: 24px;
                                        mso-line-height-rule: exactly;
                                        font-family: arial, 'helvetica neue',
                                          helvetica, sans-serif;
                                        font-size: 20px;
                                        font-style: normal;
                                        font-weight: bold;
                                        color: #ffffff;
                                      "
                                    >
                                      Contact
                                    </h3>
                                  </td>
                                </tr>
                                <tr>
                                  <td style="padding: 0; margin: 0">
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 24px;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                    ></p>
                                    <table
                                      width="100%"
                                      border="0"
                                      cellpadding="0"
                                      cellspacing="0"
                                      role="presentation"
                                      style="
                                        mso-table-lspace: 0pt;
                                        mso-table-rspace: 0pt;
                                        border-collapse: collapse;
                                        border-spacing: 0px;
                                      "
                                    >
                                      <tr>
                                        <td
                                          align="left"
                                          width="50%"
                                          valign="middle"
                                          style="
                                            padding: 0;
                                            margin: 0;
                                            padding-top: 5px;
                                            padding-bottom: 5px;
                                          "
                                        >
                                          <p
                                            style="
                                              margin: 0;
                                              -webkit-text-size-adjust: none;
                                              -ms-text-size-adjust: none;
                                              mso-line-height-rule: exactly;
                                              font-family: 'trebuchet ms',
                                                'lucida grande',
                                                'lucida sans unicode',
                                                'lucida sans', tahoma,
                                                sans-serif;
                                              line-height: 21px;
                                              color: #ffffff;
                                              font-size: 14px;
                                            "
                                          >
                                            {{.Name}}<br /><a
                                              target="_blank"
                                              style="
                                                -webkit-text-size-adjust: none;
                                                -ms-text-size-adjust: none;
                                                mso-line-height-rule: exactly;
                                                text-decoration: none;
                                                color: #ffffff;
                                                font-size: 14px;
                                              "
                                              href="tel:+(000)123-456"
                                              >{{.Phone}}</a
                                            ><br />
                                            <a
                                              target="_blank"
                                              href="mailto:{{.Email}}"
                                              style="
                                                -webkit-text-size-adjust: none;
                                                -ms-text-size-adjust: none;
                                                mso-line-height-rule: exactly;
                                                text-decoration: none;
                                                color: #ffffff;
                                                font-size: 14px;
                                              "
                                              >{{.Email}}</a
                                            >
                                          </p>
                                        </td>
                                        
                                      </tr>
                                    </table>
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 24px;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                    ></p>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                        <!--[if mso]></td></tr></table><![endif]-->
                      </td>
                    </tr>
                    <tr>
                      <td
                        align="left"
                        style="
                          margin: 0;
                          padding-top: 20px;
                          padding-left: 20px;
                          padding-right: 20px;
                          padding-bottom: 30px;
                        "
                      >
                        <!--[if mso]><table style="width:520px" cellpadding="0" cellspacing="0"><tr><td style="width:60px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-left es-table-not-adapt"
                          align="left"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: left;
                          "
                        >
                          <tr>
                            <td
                              class="es-m-p20b"
                              align="left"
                              style="
                                padding: 0;
                                margin: 0;
                                border-radius: 10px;
                                overflow: hidden;
                                width: 60px;
                              "
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                bgcolor="#41b9ef"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: separate;
                                  border-spacing: 0px;
                                  border-left: 2px solid #6abde1;
                                  border-right: 2px solid #108ec5;
                                  border-top: 2px solid #6abde1;
                                  border-bottom: 2px solid #108ec5;
                                  background-color: #41b9ef;
                                  border-radius: 10px;
                                "
                                role="presentation"
                              >
                                <tr>
                                  <td
                                    align="center"
                                    style="
                                      padding: 10px;
                                      margin: 0;
                                      font-size: 0px;
                                    "
                                  >
                                    <a
                                      target="_blank"
                                      href="https://traveloment.com/"
                                      style="
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        text-decoration: none;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                      ><img
                                        src="https://fifyhaq.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/userverificationinterfacesymbol_1.png"
                                        alt
                                        style="
                                          display: block;
                                          border: 0;
                                          outline: none;
                                          text-decoration: none;
                                          -ms-interpolation-mode: bicubic;
                                          border-radius: 10px;
                                        "
                                        width="36"
                                        height="36"
                                    /></a>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                        <!--[if mso]></td><td style="width:15px"></td><td style="width:445px" valign="top"><![endif]-->
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          class="es-right"
                          align="right"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                            float: right;
                          "
                        >
                          <tr>
                            <td
                              align="left"
                              style="padding: 0; margin: 0; width: 445px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="left"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-bottom: 10px;
                                    "
                                  >
                                    <h3
                                      style="
                                        margin: 0;
                                        line-height: 24px;
                                        mso-line-height-rule: exactly;
                                        font-family: arial, 'helvetica neue',
                                          helvetica, sans-serif;
                                        font-size: 20px;
                                        font-style: normal;
                                        font-weight: bold;
                                        color: #ffffff;
                                      "
                                    >
                                      Paid by payment card
                                    </h3>
                                  </td>
                                </tr>
                                <tr>
                                  <td style="padding: 0; margin: 0">
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 24px;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                    ></p>
                                    <table
                                      width="100%"
                                      border="0"
                                      cellpadding="0"
                                      cellspacing="0"
                                      role="presentation"
                                      style="
                                        mso-table-lspace: 0pt;
                                        mso-table-rspace: 0pt;
                                        border-collapse: collapse;
                                        border-spacing: 0px;
                                      "
                                    >
                                      <tr>
                                        <td
                                          align="left"
                                          width="50%"
                                          valign="middle"
                                          style="
                                            padding: 0;
                                            margin: 0;
                                            padding-top: 5px;
                                            padding-bottom: 5px;
                                          "
                                        >
                                          <p
                                            style="
                                              margin: 0;
                                              -webkit-text-size-adjust: none;
                                              -ms-text-size-adjust: none;
                                              mso-line-height-rule: exactly;
                                              font-family: 'trebuchet ms',
                                                'lucida grande',
                                                'lucida sans unicode',
                                                'lucida sans', tahoma,
                                                sans-serif;
                                              line-height: 21px;
                                              color: #ffffff;
                                              font-size: 14px;
                                            "
                                          >
                                            {{.BillingAmount}}
                                          </p>
                                          <p
                                            style="
                                              margin: 0;
                                              -webkit-text-size-adjust: none;
                                              -ms-text-size-adjust: none;
                                              mso-line-height-rule: exactly;
                                              font-family: 'trebuchet ms',
                                                'lucida grande',
                                                'lucida sans unicode',
                                                'lucida sans', tahoma,
                                                sans-serif;
                                              line-height: 21px;
                                              color: #ffffff;
                                              font-size: 14px;
                                            "
                                          >
                                            {{.Name}}
                                          </p>
                                        </td>
                                        <td
                                          align="right"
                                          width="50%"
                                          valign="middle"
                                          style="
                                            padding: 0;
                                            margin: 0;
                                           
                                          "
                                        >
                                        <span
                                        class="es-button-border"
                                        style="
                                          border-style: solid;
                                          border-color: #e29f33 #2cb543 #2cb543
                                            #e29f33;
                                          background: #f76b0a;
                                          border-width: 2px 0px 0px 2px;
                                          display: inline-block;
                                          border-radius: 30px;
                                          width: auto;
                                        "
                                        ><a
                                          href={{.Link}}
                                          class="es-button"
                                          target="_blank"
                                          style="
                                            mso-style-priority: 100 !important;
                                            text-decoration: none;
                                            -webkit-text-size-adjust: none;
                                            -ms-text-size-adjust: none;
                                            mso-line-height-rule: exactly;
                                            color: #ffffff;
                                            font-size: 20px;
                                            padding: 15px 30px 15px 30px;
                                            display: inline-block;
                                            background: #f76b0a;
                                            border-radius: 30px;
                                            font-family: arial, 'helvetica neue',
                                              helvetica, sans-serif;
                                            font-weight: normal;
                                            font-style: normal;
                                            line-height: 24px;
                                            width: auto;
                                            text-align: center;
                                            mso-padding-alt: 0;
                                            mso-border-alt: 10px solid #f76b0a;
                                          "
                                          >Pay Now</a
                                        >
                                      </span>
                                        </td>
                                      </tr>
                                    </table>
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 24px;
                                        color: #ffffff;
                                        font-size: 16px;
                                      "
                                    ></p>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                        <!--[if mso]></td></tr></table><![endif]-->
                      </td>
                    </tr>
                    <tr>
                      <td
                        align="left"
                        style="
                          padding: 0;
                          margin: 0;
                          padding-left: 20px;
                          padding-right: 20px;
                          padding-bottom: 30px;
                        "
                      >
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          width="100%"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                          "
                        >
                          <tr>
                            <td
                              align="center"
                              valign="top"
                              style="padding: 0; margin: 0; width: 520px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="center"
                                    class="es-m-txt-c"
                                    style="padding: 0; margin: 0"
                                  >
                                  
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
                  </table>
                </td>
              </tr>
            </table>
            <table
              cellpadding="0"
              cellspacing="0"
              class="es-footer"
              align="center"
              role="none"
              style="
                mso-table-lspace: 0pt;
                mso-table-rspace: 0pt;
                border-collapse: collapse;
                border-spacing: 0px;
                table-layout: fixed !important;
                width: 100%;
                background-color: transparent;
                background-repeat: repeat;
                background-position: center top;
              "
            >
              <tr>
                <td align="center" style="padding: 0; margin: 0">
                  <table
                    class="es-footer-body"
                    align="center"
                    cellpadding="0"
                    cellspacing="0"
                    bgcolor="#FFFFFF"
                    style="
                      mso-table-lspace: 0pt;
                      mso-table-rspace: 0pt;
                      border-collapse: collapse;
                      border-spacing: 0px;
                      background-color: #1d9ad1;
                      border-radius: 0 0 10px 10px;
                      width: 560px;
                    "
                    role="none"
                  >
                    <tr>
                      <td
                        align="left"
                        style="
                          padding: 0;
                          margin: 0;
                          padding-top: 5px;
                          padding-bottom: 5px;
                        "
                      >
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          width="100%"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                          "
                        >
                          <tr>
                            <td
                              align="center"
                              valign="top"
                              style="padding: 0; margin: 0; width: 560px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="center"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      font-size: 0px;
                                    "
                                  >
                                    <img
                                      class="adapt-img"
                                      src="https://fifyhaq.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group_37_ikx.png"
                                      alt
                                      style="
                                        display: block;
                                        border: 0;
                                        outline: none;
                                        text-decoration: none;
                                        -ms-interpolation-mode: bicubic;
                                      "
                                      width="560"
                                      height="4"
                                    />
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
                    <tr>
                      <td
                        align="left"
                        style="
                          margin: 0;
                          padding-left: 20px;
                          padding-right: 20px;
                          padding-top: 30px;
                          padding-bottom: 30px;
                        "
                      >
                        <table
                          cellpadding="0"
                          cellspacing="0"
                          width="100%"
                          role="none"
                          style="
                            mso-table-lspace: 0pt;
                            mso-table-rspace: 0pt;
                            border-collapse: collapse;
                            border-spacing: 0px;
                          "
                        >
                          <tr>
                            <td
                              align="left"
                              style="padding: 0; margin: 0; width: 520px"
                            >
                              <table
                                cellpadding="0"
                                cellspacing="0"
                                width="100%"
                                role="presentation"
                                style="
                                  mso-table-lspace: 0pt;
                                  mso-table-rspace: 0pt;
                                  border-collapse: collapse;
                                  border-spacing: 0px;
                                "
                              >
                                <tr>
                                  <td
                                    align="center"
                                    style="
                                      padding: 0;
                                      margin: 0;
                                      padding-top: 10px;
                                      padding-bottom: 10px;
                                    "
                                  >
                                    <p
                                      style="
                                        margin: 0;
                                        -webkit-text-size-adjust: none;
                                        -ms-text-size-adjust: none;
                                        mso-line-height-rule: exactly;
                                        font-family: 'trebuchet ms',
                                          'lucida grande', 'lucida sans unicode',
                                          'lucida sans', tahoma, sans-serif;
                                        line-height: 18px;
                                        color: #ffffff;
                                        font-size: 12px;
                                      "
                                    >
                                      You are receiving this email because you
                                      have visited our site or asked us about
                                      the regular newsletter. Make sure our
                                      messages get to your Inbox (and not your
                                      bulk or junk folders).<br /><strong
                                        ><a
                                          target="_blank"
                                          href="https://traveloment.com/"
                                          style="
                                            -webkit-text-size-adjust: none;
                                            -ms-text-size-adjust: none;
                                            mso-line-height-rule: exactly;
                                            text-decoration: none;
                                            color: #ffffff;
                                            font-size: 12px;
                                          "
                                          >Privacy police</a
                                        >
                                        |
                                        <a
                                          target="_blank"
                                          style="
                                            -webkit-text-size-adjust: none;
                                            -ms-text-size-adjust: none;
                                            mso-line-height-rule: exactly;
                                            text-decoration: none;
                                            color: #ffffff;
                                            font-size: 12px;
                                          "
                                          href=""
                                          >Unsubscribe</a
                                        ></strong
                                      >
                                    </p>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
                  </table>
                </td>
              </tr>
            </table>
          </td>
        </tr>
      </table>
    </div>
  </body>
</html>
`

const hotelEmailHtml = `

<html dir="ltr" xmlns="http://www.w3.org/1999/xhtml" xmlns:o="urn:schemas-microsoft-com:office:office" lang="und" style="font-family:&apos;trebuchet ms&apos;, &apos;lucida grande&apos;, &apos;lucida sans unicode&apos;, &apos;lucida sans&apos;, tahoma, sans-serif">

<head>
    <meta charset="UTF-8">
    <meta content="width=device-width, initial-scale=1" name="viewport">
    <meta name="x-apple-disable-message-reformatting">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta content="telephone=no" name="format-detection">
    <title>Booking confirmation</title>
    <!--[if (mso 16)]>
    <style type="text/css">
    a {text-decoration: none;}
    </style>
    <![endif]-->
    <!--[if gte mso 9]><style>sup { font-size: 100% !important; }</style><![endif]-->
    <!--[if gte mso 9]>
<xml>
    <o:OfficeDocumentSettings>
    <o:AllowPNG></o:AllowPNG>
    <o:PixelsPerInch>96</o:PixelsPerInch>
    </o:OfficeDocumentSettings>
</xml>
<![endif]-->
    <link rel="shortcut icon" type="image/png" href="https://stripo.email/assets/img/favicon.png">
    <style type="text/css">
        #outlook a {
            padding: 0;
        }

        .es-button {
            mso-style-priority: 100 !important;
            text-decoration: none !important;
        }

        a[x-apple-data-detectors] {
            color: inherit !important;
            text-decoration: none !important;
            font-size: inherit !important;
            font-family: inherit !important;
            font-weight: inherit !important;
            line-height: inherit !important;
        }

        .es-desk-hidden {
            display: none;
            float: left;
            overflow: hidden;
            width: 0;
            max-height: 0;
            line-height: 0;
            mso-hide: all;
        }

        @media only screen and (max-width:600px) {

            p,
            ul li,
            ol li,
            a {
                line-height: 150% !important
            }

            h1,
            h2,
            h3,
            h1 a,
            h2 a,
            h3 a {
                line-height: 120%
            }

            h1 {
                font-size: 40px !important;
                text-align: center
            }

            h2 {
                font-size: 28px !important;
                text-align: center
            }

            h3 {
                font-size: 20px !important;
                text-align: left
            }

            .es-header-body h1 a,
            .es-content-body h1 a,
            .es-footer-body h1 a {
                font-size: 40px !important;
                text-align: center
            }

            .es-header-body h2 a,
            .es-content-body h2 a,
            .es-footer-body h2 a {
                font-size: 28px !important;
                text-align: center
            }

            .es-header-body h3 a,
            .es-content-body h3 a,
            .es-footer-body h3 a {
                font-size: 20px !important;
                text-align: left
            }

            .es-menu td a {
                font-size: 14px !important
            }

            .es-header-body p,
            .es-header-body ul li,
            .es-header-body ol li,
            .es-header-body a {
                font-size: 14px !important
            }

            .es-content-body p,
            .es-content-body ul li,
            .es-content-body ol li,
            .es-content-body a {
                font-size: 14px !important
            }

            .es-footer-body p,
            .es-footer-body ul li,
            .es-footer-body ol li,
            .es-footer-body a {
                font-size: 14px !important
            }

            .es-infoblock p,
            .es-infoblock ul li,
            .es-infoblock ol li,
            .es-infoblock a {
                font-size: 12px !important
            }

            *[class="gmail-fix"] {
                display: none !important
            }

            .es-m-txt-c,
            .es-m-txt-c h1,
            .es-m-txt-c h2,
            .es-m-txt-c h3 {
                text-align: center !important
            }

            .es-m-txt-r,
            .es-m-txt-r h1,
            .es-m-txt-r h2,
            .es-m-txt-r h3 {
                text-align: right !important
            }

            .es-m-txt-l,
            .es-m-txt-l h1,
            .es-m-txt-l h2,
            .es-m-txt-l h3 {
                text-align: left !important
            }

            .es-m-txt-r img,
            .es-m-txt-c img,
            .es-m-txt-l img {
                display: inline !important
            }

            .es-button-border {
                display: inline-block !important
            }

            a.es-button,
            button.es-button {
                font-size: 18px !important;
                display: inline-block !important
            }

            .es-adaptive table,
            .es-left,
            .es-right {
                width: 100% !important
            }

            .es-content table,
            .es-header table,
            .es-footer table,
            .es-content,
            .es-footer,
            .es-header {
                width: 100% !important;
                max-width: 600px !important
            }

            .es-adapt-td {
                display: block !important;
                width: 100% !important
            }

            .adapt-img {
                width: 100% !important;
                height: auto !important
            }

            .es-m-p0 {
                padding: 0 !important
            }

            .es-m-p0r {
                padding-right: 0 !important
            }

            .es-m-p0l {
                padding-left: 0 !important
            }

            .es-m-p0t {
                padding-top: 0 !important
            }

            .es-m-p0b {
                padding-bottom: 0 !important
            }

            .es-m-p20b {
                padding-bottom: 20px !important
            }

            .es-mobile-hidden,
            .es-hidden {
                display: none !important
            }

            tr.es-desk-hidden,
            td.es-desk-hidden,
            table.es-desk-hidden {
                width: auto !important;
                overflow: visible !important;
                float: none !important;
                max-height: inherit !important;
                line-height: inherit !important
            }

            tr.es-desk-hidden {
                display: table-row !important
            }

            table.es-desk-hidden {
                display: table !important
            }

            td.es-desk-menu-hidden {
                display: table-cell !important
            }

            .es-menu td {
                width: 1% !important
            }

            table.es-table-not-adapt,
            .esd-block-html table {
                width: auto !important
            }

            table.es-social {
                display: inline-block !important
            }

            table.es-social td {
                display: inline-block !important
            }

            .es-m-p5 {
                padding: 5px !important
            }

            .es-m-p5t {
                padding-top: 5px !important
            }

            .es-m-p5b {
                padding-bottom: 5px !important
            }

            .es-m-p5r {
                padding-right: 5px !important
            }

            .es-m-p5l {
                padding-left: 5px !important
            }

            .es-m-p10 {
                padding: 10px !important
            }

            .es-m-p10t {
                padding-top: 10px !important
            }

            .es-m-p10b {
                padding-bottom: 10px !important
            }

            .es-m-p10r {
                padding-right: 10px !important
            }

            .es-m-p10l {
                padding-left: 10px !important
            }

            .es-m-p15 {
                padding: 15px !important
            }

            .es-m-p15t {
                padding-top: 15px !important
            }

            .es-m-p15b {
                padding-bottom: 15px !important
            }

            .es-m-p15r {
                padding-right: 15px !important
            }

            .es-m-p15l {
                padding-left: 15px !important
            }

            .es-m-p20 {
                padding: 20px !important
            }

            .es-m-p20t {
                padding-top: 20px !important
            }

            .es-m-p20r {
                padding-right: 20px !important
            }

            .es-m-p20l {
                padding-left: 20px !important
            }

            .es-m-p25 {
                padding: 25px !important
            }

            .es-m-p25t {
                padding-top: 25px !important
            }

            .es-m-p25b {
                padding-bottom: 25px !important
            }

            .es-m-p25r {
                padding-right: 25px !important
            }

            .es-m-p25l {
                padding-left: 25px !important
            }

            .es-m-p30 {
                padding: 30px !important
            }

            .es-m-p30t {
                padding-top: 30px !important
            }

            .es-m-p30b {
                padding-bottom: 30px !important
            }

            .es-m-p30r {
                padding-right: 30px !important
            }

            .es-m-p30l {
                padding-left: 30px !important
            }

            .es-m-p35 {
                padding: 35px !important
            }

            .es-m-p35t {
                padding-top: 35px !important
            }

            .es-m-p35b {
                padding-bottom: 35px !important
            }

            .es-m-p35r {
                padding-right: 35px !important
            }

            .es-m-p35l {
                padding-left: 35px !important
            }

            .es-m-p40 {
                padding: 40px !important
            }

            .es-m-p40t {
                padding-top: 40px !important
            }

            .es-m-p40b {
                padding-bottom: 40px !important
            }

            .es-m-p40r {
                padding-right: 40px !important
            }

            .es-m-p40l {
                padding-left: 40px !important
            }

            .es-desk-hidden {
                display: table-row !important;
                width: auto !important;
                overflow: visible !important;
                max-height: inherit !important
            }
        }

        @media screen and (max-width:384px) {
            .mail-message-content {
                width: 414px !important
            }
        }
    </style>
    <style>
        * {
            scrollbar-width: thin;
            scrollbar-color: #888 #f6f6f6;
        }

        /* Chrome, Edge, Safari */
        ::-webkit-scrollbar {
            width: 10px;
            height: 10px;
        }

        ::-webkit-scrollbar-track {
            background: #f6f6f6;
        }

        ::-webkit-scrollbar-thumb {
            background: #888;
            border-radius: 6px;
            border: 2px solid #f6f6f6;
        }

        ::-webkit-scrollbar-thumb:hover {
            box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.3);
        }

        textarea::-webkit-scrollbar-track {
            margin: 15px;
        }
    </style>
    <base href="#">
</head>

<body style="width:100%;font-family:&apos;trebuchet ms&apos;, &apos;lucida grande&apos;, &apos;lucida sans unicode&apos;, &apos;lucida sans&apos;, tahoma, sans-serif;-webkit-text-size-adjust:100%;-ms-text-size-adjust:100%;padding:0;Margin:0" data-new-gr-c-s-check-loaded="14.1167.0" data-gr-ext-installed>
    <div dir="ltr" class="es-wrapper-color" lang="und" style="background-color:#41B9EF">
        <!--[if gte mso 9]>
			<v:background xmlns:v="urn:schemas-microsoft-com:vml" fill="t">
				<v:fill type="tile" color="#41B9EF"></v:fill>
			</v:background>
		<![endif]-->
        <table class="es-wrapper" width="100%" cellspacing="0" cellpadding="0" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;padding:0;Margin:0;width:100%;background-repeat:repeat;background-position:center top;background-color:#41B9EF" role="none">
            <tbody>
                <tr>
                    <td valign="top" style="padding:0;Margin:0" class="esd-text">
                        <table cellpadding="0" cellspacing="0" class="es-header cke_show_border" align="center" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;table-layout:fixed !important;width:100%;background-color:transparent;background-repeat:repeat;background-position:center top">
                            <tbody>
                                <tr>
                                    <td align="center" style="padding:0;Margin:0">
                                        <table class="es-header-body cke_show_border" align="center" cellpadding="0" cellspacing="0" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;background-color:#1D9AD1;border-radius:10px 10px 0px 0px;width:560px" role="none">
                                            <tbody>
                                                <tr>
                                                    <td align="left" style="padding:20px;Margin:0">
                                                        <!--[if mso]><table style="width:520px" cellpadding="0" cellspacing="0"><tr><td style="width:250px" valign="top"><![endif]-->
                                                        <table cellpadding="0" cellspacing="0" align="left" class="es-left cke_show_border" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;float:left">
                                                            <tbody>
                                                                <tr>
                                                                    <td class="es-m-p0r es-m-p20b" valign="top" align="center" style="padding:0;Margin:0;width:250px">
                                                                        <table cellpadding="0" cellspacing="0" width="100%" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="left" class="es-m-txt-c" style="padding:0;Margin:0;font-size:0px"><a target="_blank" href="https://traveloment.com/" style="-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;text-decoration:none;color:#FFFFFF;font-size:14px"><img src="https://traveloment.com/wp-content/uploads/2023/08/trav_logo-removebg-preview-768x213.png" title="Logo" width="205px" alt></a></td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                        <!--[if mso]></td><td style="width:20px"></td><td style="width:250px" valign="top"><![endif]-->
                                                        <table cellpadding="0" cellspacing="0" class="es-right cke_show_border" align="right" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;float:right">
                                                            <tbody>
                                                                <tr>
                                                                    <td align="left" style="padding:0;Margin:0;width:250px">
                                                                        <table cellpadding="0" cellspacing="0" width="100%" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="right" class="es-m-txt-c esd-text" style="padding:0;Margin:0">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:21px;color:#FFFFFF;font-size:14px">{{.BookingDate}}</p>
                                                                                        <h3 style="Margin:0;line-height:24px;mso-line-height-rule:exactly;font-family:arial, 'helvetica neue', helvetica, sans-serif;font-size:20px;font-style:normal;font-weight:bold;color:#FFFFFF">Booking confirmation</h3>
                                                                                    </td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                        <!--[if mso]></td></tr></table><![endif]-->
                                                    </td>
                                                </tr>
                                                <tr>
                                                    <td align="left" style="padding:0;Margin:0;padding-top:5px;padding-bottom:5px">
                                                        <table cellpadding="0" cellspacing="0" width="100%" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                                                            <tbody>
                                                                <tr>
                                                                    <td align="center" valign="top" style="padding:0;Margin:0;width:560px">
                                                                        <table cellpadding="0" cellspacing="0" width="100%" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="center" style="padding:0;Margin:0;font-size:0px"><img class="adapt-img" src="https://tlr.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group_37_ikx.png" alt style="display:block;border:0;outline:none;text-decoration:none;-ms-interpolation-mode:bicubic" width="560"></td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                        <table cellpadding="0" cellspacing="0" class="es-content cke_show_border" align="center" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;table-layout:fixed !important;width:100%">
                            <tbody>
                                <tr>
                                    <td align="center" style="padding:0;Margin:0">
                                        <table bgcolor="#ffffff" class="es-content-body cke_show_border" align="center" cellpadding="0" cellspacing="0" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;background-color:#1D9AD1;border-radius:0 0 10px 10px;width:560px" role="none">
                                            <tbody>
                                                <tr>
                                                    <td align="left" style="padding:0;Margin:0;padding-top:20px;padding-left:20px;padding-bottom:30px">
                                                        <!--[if mso]><table style="width:540px" cellpadding="0" cellspacing="0"><tr><td style="width:312px" valign="top"><![endif]-->
                                                        <table cellpadding="0" cellspacing="0" align="left" class="es-left cke_show_border" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;float:left">
                                                            <tbody>
                                                                <tr>
                                                                    <td class="es-m-p0r es-m-p20b" valign="top" align="center" style="padding:0;Margin:0;width:312px">
                                                                        <table cellpadding="0" cellspacing="0" width="100%" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="left" class="es-m-p20r es-m-txt-l esd-text" style="padding:0;Margin:0;padding-bottom:10px;padding-top:20px">
                                                                                        <h1 style="Margin:0;line-height:48px;mso-line-height-rule:exactly;font-family:arial, 'helvetica neue', helvetica, sans-serif;font-size:40px;font-style:normal;font-weight:bold;color:#FFFFFF">Hi, your booking is confirmed!</h1>
                                                                                    </td>
                                                                                </tr>
                                                                                <tr>
                                                                                    <td align="left" class="es-m-p20r esd-text" style="padding:0;Margin:0;padding-top:10px;padding-bottom:10px;padding-right:10px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:24px;color:#FFFFFF;font-size:16px">Please find your booking code below.<br>We look forward to seeing you soon!</p>
                                                                                    </td>
                                                                                </tr>
                                                                                <tr>
                                                                                    <td align="left" class="es-m-p20r esd-text" style="padding:0;Margin:0;padding-top:10px;padding-bottom:10px;padding-right:10px">
                                                                                        <h3 style="Margin:0;line-height:24px;mso-line-height-rule:exactly;font-family:arial, 'helvetica neue', helvetica, sans-serif;font-size:20px;font-style:normal;font-weight:bold;color:#f76b0a"><span style="background-color:#FFFFFF;border-radius:10px">&nbsp;Booking No. {{.BookingNo}}</span></h3>
                                                                                    </td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                        <!--[if mso]></td><td style="width:20px"></td><td style="width:208px" valign="top"><![endif]-->
                                                        <table cellpadding="0" cellspacing="0" class="es-right cke_show_border" align="right" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;float:right">
                                                            <tbody>
                                                                <tr>
                                                                    <td align="left" style="padding:0;Margin:0;width:208px">
                                                                        <table cellpadding="0" cellspacing="0" width="100%" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="center" style="padding:0;Margin:0;font-size:0px"><a target="_blank" href="https://traveloment.com/" style="-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;text-decoration:none;color:#FFFFFF;font-size:16px"><img class="adapt-img" src="https://tlr.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/5883100.png" alt style="display:block;border:0;outline:none;text-decoration:none;-ms-interpolation-mode:bicubic" width="208"></a></td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                        <!--[if mso]></td></tr></table><![endif]-->
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                        
                    </td>
                </tr>
            </tbody>
        </table>
        <table cellpadding="0" cellspacing="0" class="es-content" align="center" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-spacing:0px;table-layout:fixed !important;width:100%">
            <tbody>
                  <tr>
                    <td align="left" style="padding:0;Margin:0;padding-top:5px;padding-bottom:5px">
                        <table cellpadding="0" cellspacing="0" width="100%" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                            <tbody>
                                <tr>
                                    <td align="center" valign="top" style="padding:0;Margin:0;width:560px">
                                        <table cellpadding="0" cellspacing="0" width="100%" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                                            <tbody>
                                                <tr>
                                                    <td align="center" style="padding:0;Margin:0;font-size:0px"><img class="adapt-img" src="https://tlr.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group_37_ikx.png" alt style="display:block;border:0;outline:none;text-decoration:none;-ms-interpolation-mode:bicubic" width="560"></td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </td>
                </tr>
                <tr>
                    <td align="center" style="padding:0;Margin:0">
                        <table bgcolor="#ffffff" class="es-content-body" align="center" cellpadding="0" cellspacing="0" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;background-color:#1D9AD1;border-radius:10px 10px 0 0;width:560px" role="none">
                            <tbody>
                                <tr>
                                    <td align="left" style="padding:0;Margin:0;padding-left:20px;padding-right:20px;padding-top:30px">
                                        <table cellpadding="0" cellspacing="0" width="100%" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                            <tbody>
                                                <tr>
                                                    <td align="center" valign="top" style="padding:0;Margin:0;width:520px">
                                                        <table cellpadding="0" cellspacing="0" width="100%" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                                            <tbody>
                                                                <tr>
                                                                    <td align="left" style="padding:0;Margin:0" class="esd-text">
                                                                        <h2 style="Margin:0;line-height:34px;mso-line-height-rule:exactly;font-family:arial, 'helvetica neue', helvetica, sans-serif;font-size:28px;font-style:normal;font-weight:bold;color:#FFFFFF">Booking Details</h2>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                                <tr>
                                    <td align="left" style="Margin:0;padding-bottom:10px;padding-top:20px;padding-left:20px;padding-right:20px">
                                        <table cellpadding="0" cellspacing="0" width="100%" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                            <tbody>
                                                <tr>
                                                    <td class="es-m-p0r es-m-p20b" align="center" style="padding:0;Margin:0;width:520px">
                                                        <table cellpadding="0" cellspacing="0" width="100%" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:separate;border-spacing:0px;border-left:2px solid #6ABDE1;border-right:2px solid #108EC5;border-top:2px solid #6ABDE1;border-bottom:2px solid #108EC5;border-radius:10px" role="presentation">
                                                            <tbody>
                                                                <tr>
                                                                    <td style="padding:0;Margin:0">
                                                                        <table cellpadding="0" cellspacing="0" width="100%" class="es-menu" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                                                            <tbody>
                                                                                <tr class="links-images-left">
                                                                                    <td align="left" valign="top" width="100%" id="esd-menu-id-0" style="Margin:0;padding-top:10px;padding-bottom:10px;padding-left:20px;padding-right:20px;border:0"><a target="_blank" style="-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;text-decoration:none;display:flex;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;color:#FFFFFF;font-size:20px;font-weight:bold"><img src="https://fifyhaq.stripocdn.email/content/guids/CABINET_ef87daf557d03e10a7278bd7d7dd7a89f52d4697ba1e208101ce9075edfc9d69/images/hotelslogo41820.png" alt="Luggage added" title="Luggage added" align="absmiddle" width="40" style="border: 0px; outline: none; text-decoration: none; padding-right: 15px; vertical-align: middle; display: block;">Hotel Information</a></td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                                <tr>
                                                                    <td align="center" style="padding:0;Margin:0;font-size:0px"><img src="https://tlr.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group_37_fmh.png" alt style="display:block;border:0;outline:none;text-decoration:none;-ms-interpolation-mode:bicubic" width="516" class="adapt-img"></td>
                                                                </tr>
                                                                <tr>
                                                                    <td align="left" style="Margin:0;padding-top:10px;padding-bottom:10px;padding-left:20px;padding-right:20px" class="esd-text">
                                                                        <p style = "Margin:0;font-family:'trebuchet ms','lucida grande','lucida sans unicode','lucida sans',tahoma,sans-serif;line-height:24px;color:#ffffff;font-size:16px">{{.HotelName}} || {{.HotelRatting}}</p>
                                                                    </td>
                                                                </tr>
                                                                <tr>
                                                                    <td align="center" style="padding:0;Margin:0;font-size:0px"><img src="https://tlr.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group_37_fmh.png" alt style="display:block;border:0;outline:none;text-decoration:none;-ms-interpolation-mode:bicubic" width="516" class="adapt-img"></td>
                                                                </tr>
                                                                <tr>
                                                                    <td style="padding:0;Margin:0" class="esd-text" align="left">
                                                                        <table width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" class=" cke_show_border">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="left" width="50%" valign="middle" style="Margin:0;;padding-left:20px;padding-right:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:21px;color:#FFFFFF;font-size:14px">Address</p>
                                                                                    </td>
                                                                                    <td align="right" width="50%" valign="middle" style="Margin:0;;padding-left:20px;padding-right:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:21px;color:#FFFFFF;font-size:14px">J{{.HotelAddress}}</p>
                                                                                    </td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                                <tr>
                                                                    <td style="padding:0;Margin:0" class="esd-text" align="left">
                                                                        <table width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="left" width="50%" valign="middle" style="Margin:0;padding-top:5px;padding-bottom:5px;padding-left:20px;padding-right:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:21px;color:#FFFFFF;font-size:14px">Phone:</p>
                                                                                    </td>
                                                                                    <td align="right" width="50%" valign="middle" style="Margin:0;padding-top:5px;padding-bottom:5px;padding-left:20px;padding-right:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:21px;color:#FFFFFF;font-size:14px">{{.HotelPhone}}</p>
                                                                                    </td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                                <tr>
                                                                    <td style="padding:0;Margin:0" class="esd-text" align="left">
                                                                        <table width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="left" width="50%" valign="middle" style="Margin:0;padding-left:20px;padding-right:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:21px;color:#FFFFFF;font-size:14px">Email:&nbsp;</p>
                                                                                    </td>
                                                                                    <td align="right" width="50%" valign="middle" style="Margin:0;padding-left:20px;padding-right:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:27px;color:#FFFFFF;font-size:18px">{{.HotelEmail}}</p>
                                                                                    </td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                                <tr>
                                    <td align="left" style="Margin:0;padding-bottom:10px;padding-top:20px;padding-left:20px;padding-right:20px">
                                        <table cellpadding="0" cellspacing="0" width="100%" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                            <tbody>
                                                <tr>
                                                    <td class="es-m-p0r es-m-p20b" align="center" style="padding:0;Margin:0;width:520px">
                                                        <table cellpadding="0" cellspacing="0" width="100%" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:separate;border-spacing:0px;border-left:2px solid #6ABDE1;border-right:2px solid #108EC5;border-top:2px solid #6ABDE1;border-bottom:2px solid #108EC5;border-radius:10px" role="presentation">
                                                            <tbody>
                                                                <tr>
                                                                    <td style="padding:0;Margin:0">
                                                                        <table cellpadding="0" cellspacing="0" width="100%" class="es-menu" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                                                            <tbody>
                                                                                <tr class="links-images-left">
                                                                                    <td align="left" valign="top" width="100%" id="esd-menu-id-0" style="Margin:0;padding-top:10px;padding-bottom:10px;padding-left:20px;padding-right:20px;border:0"><a target="_blank" style="-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;text-decoration:none;display:flex;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;color:#FFFFFF;font-size:20px;font-weight:bold"><img src="https://fifyhaq.stripocdn.email/content/guids/CABINET_ef87daf557d03e10a7278bd7d7dd7a89f52d4697ba1e208101ce9075edfc9d69/images/traveling.png" alt="Luggage added" title="Luggage added" align="absmiddle" width="40" style="border: 0px; outline: none; text-decoration: none; padding-right: 15px; vertical-align: middle; display: block;">No. of Guests</a></td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                                <tr>
                                                                    <td align="center" style="padding:0;Margin:0;font-size:0px"><img src="https://tlr.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group_37_fmh.png" alt style="display:block;border:0;outline:none;text-decoration:none;-ms-interpolation-mode:bicubic" width="516" class="adapt-img"></td>
                                                                </tr>
                                                                
                                                                <tr>
                                                                    <td align="center" style="padding:0;Margin:0;font-size:0px"><img src="https://tlr.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group_37_fmh.png" alt style="display:block;border:0;outline:none;text-decoration:none;-ms-interpolation-mode:bicubic" width="516" class="adapt-img"></td>
                                                                </tr>
                                                                <tr>
                                                                    <td style="padding:0;Margin:0" class="esd-text" align="left">
                                                                        <table width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="left" width="50%" valign="middle" style="Margin:0;padding-top:5px;padding-bottom:5px;padding-left:20px;padding-right:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:21px;color:#FFFFFF;font-size:14px">No. Of Rooms</p>
                                                                                    </td>
                                                                                    <td align="right" width="50%" valign="middle" style="Margin:0;padding-top:5px;padding-bottom:5px;padding-left:20px;padding-right:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:27px;color:#FFFFFF;font-size:18px">{{.RoomsNo}}</p>
                                                                                    </td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                                <tr>
                                                                    <td style="padding:0;Margin:0" class="esd-text" align="left">
                                                                        <table width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="left" width="50%" valign="middle" style="Margin:0;padding-top:5px;padding-bottom:5px;padding-left:20px;padding-right:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:21px;color:#FFFFFF;font-size:14px">Adults</p>
                                                                                    </td>
                                                                                    <td align="right" width="50%" valign="middle" style="Margin:0;padding-top:5px;padding-bottom:5px;padding-left:20px;padding-right:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:27px;color:#FFFFFF;font-size:18px">{{.AdultsNo}}</p>
                                                                                    </td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                                <tr>
                                                                    <td style="padding:0;Margin:0" class="esd-text" align="left">
                                                                        <table width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="left" width="50%" valign="middle" style="Margin:0;padding-top:5px;padding-bottom:10px;padding-left:20px;padding-right:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:21px;color:#FFFFFF;font-size:14px">Childs</p>
                                                                                    </td>
                                                                                    <td align="right" width="50%" valign="middle" style="Margin:0;padding-top:5px;padding-bottom:10px;padding-left:20px;padding-right:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:27px;color:#FFFFFF;font-size:18px">{{.ChildsNo}}</p>
                                                                                    </td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                                <tr>
                                    <td align="left" style="Margin:0;padding-bottom:10px;padding-top:20px;padding-left:20px;padding-right:20px">
                                        <table cellpadding="0" cellspacing="0" width="100%" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                            <tbody>
                                                <tr>
                                                    <td class="es-m-p0r es-m-p20b" align="center" style="padding:0;Margin:0;width:520px">
                                                        <table cellpadding="0" cellspacing="0" width="100%" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:separate;border-spacing:0px;border-left:2px solid #6ABDE1;border-right:2px solid #108EC5;border-top:2px solid #6ABDE1;border-bottom:2px solid #108EC5;border-radius:10px" role="presentation">
                                                            <tbody>
                                                                <tr>
                                                                    <td style="padding:0;Margin:0">
                                                                        <table cellpadding="0" cellspacing="0" width="100%" class="es-menu" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                                                            <tbody>
                                                                                <tr class="links-images-left">
                                                                                    <td align="left" valign="top" width="100%" id="esd-menu-id-0" style="Margin:0;padding-top:10px;padding-bottom:10px;padding-left:20px;padding-right:20px;border:0"><a target="_blank" style="-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;text-decoration:none;display:flex;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;color:#FFFFFF;font-size:20px;font-weight:bold"><img src="https://fifyhaq.stripocdn.email/content/guids/CABINET_ef87daf557d03e10a7278bd7d7dd7a89f52d4697ba1e208101ce9075edfc9d69/images/location.png" alt="Luggage added" title="Luggage added" align="absmiddle" width="40" style="border: 0px; outline: none; text-decoration: none; padding-right: 15px; vertical-align: middle; display: block;">Check in - Check out Details</a></td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                                <tr>
                                                                    <td align="center" style="padding:0;Margin:0;font-size:0px"><img src="https://tlr.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group_37_fmh.png" alt style="display:block;border:0;outline:none;text-decoration:none;-ms-interpolation-mode:bicubic" width="516" class="adapt-img"></td>
                                                                </tr>
                                                                
                                                                <tr>
                                                                    <td align="center" style="padding:0;Margin:0;font-size:0px"><img src="https://tlr.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group_37_fmh.png" alt style="display:block;border:0;outline:none;text-decoration:none;-ms-interpolation-mode:bicubic" width="516" class="adapt-img"></td>
                                                                </tr>
                                                                <tr>
                                                                    <td style="padding:0;Margin:0" class="esd-text" align="left">
                                                                        <table width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="left" width="50%" valign="middle" style="Margin:0;padding-top:5px;padding-bottom:5px;padding-left:20px;padding-right:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:21px;color:#FFFFFF;font-size:14px">Check in </p>
                                                                                    </td>
                                                                                    <td align="right" width="50%" valign="middle" style="Margin:0;padding-top:5px;padding-bottom:5px;padding-left:20px;padding-right:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:27px;color:#FFFFFF;font-size:18px">{{.CheckinDate}} {{.CheckinTime}}</p>
                                                                                    </td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                                <tr>
                                                                    <td style="padding:0;Margin:0" class="esd-text" align="left">
                                                                        <table width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="left" width="50%" valign="middle" style="Margin:0;padding-top:5px;padding-bottom:5px;padding-left:20px;padding-right:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:21px;color:#FFFFFF;font-size:14px">Check out </p>
                                                                                    </td>
                                                                                    <td align="right" width="50%" valign="middle" style="Margin:0;padding-top:5px;padding-bottom:5px;padding-left:20px;padding-right:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:27px;color:#FFFFFF;font-size:18px">{{.CheckoutDate}} {{.CheckoutTime}}</p>
                                                                                    </td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                                <tr>
                                    <td align="left" style="Margin:0;padding-top:20px;padding-left:20px;padding-right:20px;padding-bottom:30px">
                                        <table cellpadding="0" cellspacing="0" width="100%" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                            <tbody>
                                                <tr>
                                                    <td class="es-m-p0r es-m-p20b" align="center" style="padding:0;Margin:0;width:520px">
                                                        <table cellpadding="0" cellspacing="0" width="100%" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:separate;border-spacing:0px;border-left:2px solid #6ABDE1;border-right:2px solid #108EC5;border-top:2px solid #6ABDE1;border-bottom:2px solid #108EC5;border-radius:10px" role="presentation">
                                                            <tbody>
                                                                <tr>
                                                                    <td style="padding:0;Margin:0">
                                                                        <table cellpadding="0" cellspacing="0" width="100%" class="es-menu" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                                                            <tbody>
                                                                                <tr class="links-images-left">
                                                                                    <td align="left" valign="top" width="100%" id="esd-menu-id-0" style="Margin:0;padding-top:10px;padding-bottom:10px;padding-left:20px;padding-right:20px;border:0"><a target="_blank" style="-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;text-decoration:none;display:flex;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;color:#FFFFFF;font-size:20px;font-weight:bold"><img src="https://fifyhaq.stripocdn.email/content/guids/CABINET_ef87daf557d03e10a7278bd7d7dd7a89f52d4697ba1e208101ce9075edfc9d69/images/pngwingcom.png" alt="Add taxi upon arrival" title="Add taxi upon arrival" align="absmiddle" width="40" style="border: 0px; outline: none; text-decoration: none; padding-right: 15px; vertical-align: middle; display: block;">Room Type</a></td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                                <tr>
                                                                    <td align="center" style="padding:0;Margin:0;font-size:0px"><img src="https://tlr.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group_37_fmh.png" alt style="display:block;border:0;outline:none;text-decoration:none;-ms-interpolation-mode:bicubic" width="516" class="adapt-img"></td>
                                                                </tr>
                                                                <tr>
                                                                    <td style="padding:0;Margin:0" class="esd-text">
                                                                        <table width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="left" width="50%" valign="middle" style="padding:0;Margin:0;padding-top:10px;padding-bottom:15px;padding-left:20px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:24px;color:#FFFFFF;font-size:16px">{{.RoomType}}</p>
                                                                                    </td>
                                                                                    
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                               
                                <tr>
                                    <td align="left" style="padding:0;Margin:0;padding-top:5px;padding-bottom:5px">
                                        <table cellpadding="0" cellspacing="0" width="100%" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                            <tbody>
                                                <tr>
                                                    <td align="center" valign="top" style="padding:0;Margin:0;width:560px">
                                                        <table cellpadding="0" cellspacing="0" width="100%" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                                            <tbody>
                                                                <tr>
                                                                    <td align="center" style="padding:0;Margin:0;font-size:0px"><img class="adapt-img" src="https://tlr.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group_37_ikx.png" alt style="display:block;border:0;outline:none;text-decoration:none;-ms-interpolation-mode:bicubic" width="560"></td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </td>
                </tr>
            </tbody>
        </table>
        <table cellpadding="0" cellspacing="0" class="es-content" align="center" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;table-layout:fixed !important;width:100%">
            <tbody>
                <tr>
                    <td align="center" style="padding:0;Margin:0">
                        <table bgcolor="#ffffff" class="es-content-body" align="center" cellpadding="0" cellspacing="0" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;background-color:#1D9AD1;width:560px">
                            <tbody>
                                <tr>
                                    <td align="left" style="padding:0;Margin:0;padding-left:20px;padding-right:20px;padding-top:30px">
                                        <table cellpadding="0" cellspacing="0" width="100%" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                            <tbody>
                                                <tr>
                                                    <td align="center" valign="top" style="padding:0;Margin:0;width:520px">
                                                        <table cellpadding="0" cellspacing="0" width="100%" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                                            <tbody>
                                                                <tr>
                                                                    <td align="left" style="padding:0;Margin:0" class="esd-text">
                                                                        <h2 style="Margin:0;line-height:34px;mso-line-height-rule:exactly;font-family:arial, 'helvetica neue', helvetica, sans-serif;font-size:28px;font-style:normal;font-weight:bold;color:#FFFFFF">Customer Details</h2>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                                
                                
                                <tr>
                                    <td align="left" style="padding:0;Margin:0;padding-top:20px;padding-left:20px;padding-right:20px">
                                        <!--[if mso]><table style="width:520px" cellpadding="0" cellspacing="0"><tr><td style="width:60px" valign="top"><![endif]-->
                                        <table cellpadding="0" cellspacing="0" class="es-left es-table-not-adapt" align="left" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;float:left">
                                            <tbody>
                                                <tr>
                                                    <td class="es-m-p20b" align="left" style="padding:0;Margin:0;border-radius:10px;overflow:hidden;width:60px">
                                                        <table cellpadding="0" cellspacing="0" width="100%" bgcolor="#41b9ef" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:separate;border-spacing:0px;border-left:2px solid #6ABDE1;border-right:2px solid #108EC5;border-top:2px solid #6ABDE1;border-bottom:2px solid #108EC5;background-color:#41b9ef;border-radius:10px" role="presentation">
                                                            <tbody>
                                                                <tr>
                                                                    <td align="center" style="padding:10px;Margin:0;font-size:0px"><a target="_blank" href="https://traveloment.com/" style="-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;text-decoration:none;color:#FFFFFF;font-size:16px"><img src="https://tlr.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/userverificationinterfacesymbol.png" alt style="display:block;border:0;outline:none;text-decoration:none;-ms-interpolation-mode:bicubic;border-radius:10px" width="36"></a></td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                        <!--[if mso]></td><td style="width:15px"></td><td style="width:445px" valign="top"><![endif]-->
                                        <table cellpadding="0" cellspacing="0" class="es-right" align="right" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;float:right">
                                            <tbody>
                                                <tr>
                                                    <td align="left" style="padding:0;Margin:0;width:445px">
                                                        <table cellpadding="0" cellspacing="0" width="100%" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                                            <tbody>
                                                                <tr>
                                                                    <td align="left" style="padding:0;Margin:0;" class="esd-text">
                                                                        <h3 style="Margin:0;line-height:24px;mso-line-height-rule:exactly;font-family:arial, 'helvetica neue', helvetica, sans-serif;font-size:20px;font-style:normal;font-weight:bold;color:#FFFFFF">Guests name</h3>
                                                                    </td>
                                                                </tr>
                                                                <tr>
                                                                    <td style="padding:0;Margin:0" class="esd-text" align="left">
                                                                        <table width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px" class=" cke_show_border">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="left" width="50%" valign="middle" style="padding:0;Margin:0;padding-top:5px;padding-bottom:5px;display:flex">
                                                                                    {{range .GuestName}}
                                                                                    <p
                                                                                      style="
                                                                                        margin: 0;
                                                                                        -webkit-text-size-adjust: none;
                                                                                        -ms-text-size-adjust: none;
                                                                                        mso-line-height-rule: exactly;
                                                                                        font-family: 'trebuchet ms',
                                                                                          'lucida grande',
                                                                                          'lucida sans unicode',
                                                                                          'lucida sans', tahoma,
                                                                                          sans-serif;
                                                                                        line-height: 21px;
                                                                                        color: #ffffff;
                                                                                        font-size: 14px;
                                                                                        
                                                                                      "
                                                                                    >
                                                                                    
                                                                                    
                                                                                    {{.}},
                                                                                    
                                                                                    
                                                                                    </p>
                                                                                    {{end}}
                                                                                    </td>
                                                                                    
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                        <!--[if mso]></td></tr></table><![endif]-->
                                    </td>
                                </tr>
                                <tr>
                                    <td align="left" style="padding:0;Margin:0;padding-top:20px;padding-left:20px;padding-right:20px">
                                        <!--[if mso]><table style="width:520px" cellpadding="0" cellspacing="0"><tr><td style="width:60px" valign="top"><![endif]-->
                                        <table cellpadding="0" cellspacing="0" class="es-left es-table-not-adapt" align="left" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;float:left">
                                            <tbody>
                                                <tr>
                                                    <td class="es-m-p20b" align="left" style="padding:0;Margin:0;border-radius:10px;overflow:hidden;width:60px">
                                                        <table cellpadding="0" cellspacing="0" width="100%" bgcolor="#41b9ef" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:separate;border-spacing:0px;border-left:2px solid #6ABDE1;border-right:2px solid #108EC5;border-top:2px solid #6ABDE1;border-bottom:2px solid #108EC5;background-color:#41b9ef;border-radius:10px" role="presentation">
                                                            <tbody>
                                                                <tr>
                                                                    <td align="center" style="padding:10px;Margin:0;font-size:0px"><a target="_blank" href="https://traveloment.com/" style="-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;text-decoration:none;color:#FFFFFF;font-size:16px"><img src="https://tlr.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/userverificationinterfacesymbol_2.png" alt style="display:block;border:0;outline:none;text-decoration:none;-ms-interpolation-mode:bicubic;border-radius:10px" width="36"></a></td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                        <!--[if mso]></td><td style="width:15px"></td><td style="width:445px" valign="top"><![endif]-->
                                        <table cellpadding="0" cellspacing="0" class="es-right" align="right" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;float:right">
                                            <tbody>
                                                <tr>
                                                    <td align="left" style="padding:0;Margin:0;width:445px">
                                                        <table cellpadding="0" cellspacing="0" width="100%" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                                            <tbody>
                                                                <tr>
                                                                    <td align="left" style="padding:0;Margin:0;" class="esd-text">
                                                                        <h3 style="Margin:0;line-height:24px;mso-line-height-rule:exactly;font-family:arial, 'helvetica neue', helvetica, sans-serif;font-size:20px;font-style:normal;font-weight:bold;color:#FFFFFF">Contact</h3>
                                                                    </td>
                                                                </tr>
                                                                <tr>
                                                                    <td style="padding:0;Margin:0" class="esd-text">
                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:24px;color:#FFFFFF;font-size:16px"></p>
                                                                        <table width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="left" width="50%" valign="middle" style="padding:0;Margin:0;padding-top:5px;padding-bottom:5px">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:21px;color:#FFFFFF;font-size:14px">{{.Name}}<br><a target="_blank" style="-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;text-decoration:none;color:#FFFFFF;font-size:14px" href="tel:+(000)123-456">{{.Phone}} </a><a target="_blank" href="mailto:tyler_sanders@mail.com" style="-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;text-decoration:none;color:#FFFFFF;font-size:14px">{{.Email}}</a></p>
                                                                                    </td>
                                                                                    
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:24px;color:#FFFFFF;font-size:16px"></p>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                        <!--[if mso]></td></tr></table><![endif]-->
                                    </td>
                                </tr>
                                <tr>
                                    <td align="left" style="Margin:0;padding-top:20px;padding-left:20px;padding-right:20px;padding-bottom:30px">
                                        <!--[if mso]><table style="width:520px" cellpadding="0" cellspacing="0"><tr><td style="width:60px" valign="top"><![endif]-->
                                        <table cellpadding="0" cellspacing="0" class="es-left es-table-not-adapt" align="left" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;float:left">
                                            <tbody>
                                                <tr>
                                                    <td class="es-m-p20b" align="left" style="padding:0;Margin:0;border-radius:10px;overflow:hidden;width:60px">
                                                        <table cellpadding="0" cellspacing="0" width="100%" bgcolor="#41b9ef" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:separate;border-spacing:0px;border-left:2px solid #6ABDE1;border-right:2px solid #108EC5;border-top:2px solid #6ABDE1;border-bottom:2px solid #108EC5;background-color:#41b9ef;border-radius:10px" role="presentation">
                                                            <tbody>
                                                                <tr>
                                                                    <td align="center" style="padding:10px;Margin:0;font-size:0px"><a target="_blank" href="https://traveloment.com/" style="-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;text-decoration:none;color:#FFFFFF;font-size:16px"><img src="https://tlr.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/userverificationinterfacesymbol_1.png" alt style="display:block;border:0;outline:none;text-decoration:none;-ms-interpolation-mode:bicubic;border-radius:10px" width="36"></a></td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                        <!--[if mso]></td><td style="width:15px"></td><td style="width:445px" valign="top"><![endif]-->
                                        <table cellpadding="0" cellspacing="0" class="es-right" align="right" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;float:right">
                                            <tbody>
                                                <tr>
                                                    <td align="left" style="padding:0;Margin:0;width:445px">
                                                        <table cellpadding="0" cellspacing="0" width="100%" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                                            <tbody>
                                                                <tr>
                                                                    <td align="left" style="padding:0;Margin:0;" class="esd-text">
                                                                        <h3 style="Margin:0;line-height:24px;mso-line-height-rule:exactly;font-family:arial, 'helvetica neue', helvetica, sans-serif;font-size:20px;font-style:normal;font-weight:bold;color:#FFFFFF">Paid by payment card</h3>
                                                                    </td>
                                                                </tr>
                                                                <tr>
                                                                    <td style="padding:0;Margin:0" class="esd-text">
                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:24px;color:#FFFFFF;font-size:16px"></p>
                                                                        <table width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="left" width="50%" valign="middle" style="padding:0;Margin:0;padding-top:5px;">
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:21px;color:#FFFFFF;font-size:14px">{{.BillingAmount}}</p></br>
                                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:21px;color:#FFFFFF;font-size:14px"> {{.Name}}</p>
                                                                                    </td>
                                                                                    <td
                                          align="right"
                                          width="50%"
                                          valign="middle"
                                          style="
                                            padding: 0;
                                            margin: 0;
                                           
                                          "
                                        >
                                        <span
                                        class="es-button-border"
                                        style="
                                          border-style: solid;
                                          border-color: #e29f33 #2cb543 #2cb543
                                            #e29f33;
                                          background: #f76b0a;
                                          border-width: 2px 0px 0px 2px;
                                          display: inline-block;
                                          border-radius: 30px;
                                          width: auto;
                                        "
                                        ><a
                                          href={{.Link}}
                                          class="es-button"
                                          target="_blank"
                                          style="
                                            mso-style-priority: 100 !important;
                                            text-decoration: none;
                                            -webkit-text-size-adjust: none;
                                            -ms-text-size-adjust: none;
                                            mso-line-height-rule: exactly;
                                            color: #ffffff;
                                            font-size: 20px;
                                            padding: 15px 30px 15px 30px;
                                            display: inline-block;
                                            background: #f76b0a;
                                            border-radius: 30px;
                                            font-family: arial, 'helvetica neue',
                                              helvetica, sans-serif;
                                            font-weight: normal;
                                            font-style: normal;
                                            line-height: 24px;
                                            width: auto;
                                            text-align: center;
                                            mso-padding-alt: 0;
                                            mso-border-alt: 10px solid #f76b0a;
                                          "
                                          >Pay Now</a
                                        >
                                      </span>
                                        </td>
                                                                                    
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                        <!--[if mso]></td></tr></table><![endif]-->
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </td>
                </tr>
            </tbody>
        </table>
        <table cellpadding="0" cellspacing="0" class="es-footer" align="center" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;table-layout:fixed !important;width:100%;background-color:transparent;background-repeat:repeat;background-position:center top">
            <tbody>
                <tr>
                    <td align="center" style="padding:0;Margin:0">
                        <table class="es-footer-body" align="center" cellpadding="0" cellspacing="0" bgcolor="#FFFFFF" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;background-color:#1D9AD1;border-radius:0 0 10px 10px;width:560px" role="none">
                            <tbody>
                                <tr>
                                    <td align="left" style="padding:0;Margin:0;padding-top:5px;padding-bottom:5px">
                                        <table cellpadding="0" cellspacing="0" width="100%" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                            <tbody>
                                                <tr>
                                                    <td align="center" valign="top" style="padding:0;Margin:0;width:560px">
                                                        <table cellpadding="0" cellspacing="0" width="100%" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                                            <tbody>
                                                                <tr>
                                                                    <td align="center" style="padding:0;Margin:0;font-size:0px"><img class="adapt-img" src="https://tlr.stripocdn.email/content/guids/CABINET_6369f7e833d10e63cc6e7964ca64418d/images/group_37_ikx.png" alt style="display:block;border:0;outline:none;text-decoration:none;-ms-interpolation-mode:bicubic" width="560"></td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                                <tr>
                                    <td align="left" style="Margin:0;padding-left:20px;padding-right:20px;padding-top:30px;padding-bottom:30px">
                                        <table cellpadding="0" cellspacing="0" width="100%" role="none" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                            <tbody>
                                                <tr>
                                                    <td align="left" style="padding:0;Margin:0;width:520px">
                                                        <table cellpadding="0" cellspacing="0" width="100%" role="presentation" style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px">
                                                            <tbody>
                                                                <tr>
                                                                    <td align="center" style="padding:0;Margin:0;padding-top:10px;padding-bottom:10px" class="esd-text">
                                                                        <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:'trebuchet ms', 'lucida grande', 'lucida sans unicode', 'lucida sans', tahoma, sans-serif;line-height:18px;color:#FFFFFF;font-size:12px">{{.CancelPolicy}}.</p>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</body>

</html>`

func sendCarEmailHandler(c *gin.Context) {
	var emailReq EmailRequest
	err := c.BindJSON(&emailReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password := os.Getenv("EMAIL_PASSWORD")
	if password == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "EMAIL_PASSWORD environment variable is not set"})
		return
	}

	tmpl, err := template.New("email").Parse(carEmailHTML)
	if err != nil {
		fmt.Println("Error parsing HTML template:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var emailBody bytes.Buffer
	err = tmpl.Execute(&emailBody, emailReq.BookingDetails)
	if err != nil {
		fmt.Println("Error executing HTML template:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sendEmail(emailReq.Subject, emailBody.String(), []string{emailReq.To}, password)

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Email sent successfully to %s", emailReq.To)})
}

func sendFlightEmailHandler(c *gin.Context) {
	var flightEmailReq FlightEmailRequest
	err := c.BindJSON(&flightEmailReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password := os.Getenv("EMAIL_PASSWORD")
	if password == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "EMAIL_PASSWORD environment variable is not set"})
		return
	}

	tmpl, err := template.New("email").Parse(flightEmailHtml)
	if err != nil {
		fmt.Println("Error parsing HTML template:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var emailBody bytes.Buffer
	err = tmpl.Execute(&emailBody, flightEmailReq.FlightBookingDetails)
	if err != nil {
		fmt.Println("Error executing HTML template:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sendEmail(flightEmailReq.Subject, emailBody.String(), []string{flightEmailReq.To}, password)

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Email sent successfully to %s", flightEmailReq.To)})
}

func sendHotelEmailHandler(c *gin.Context) {
	var hotelEmailReq HotelEmailRequest
	err := c.BindJSON(&hotelEmailReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password := os.Getenv("EMAIL_PASSWORD")
	if password == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "EMAIL_PASSWORD environment variable is not set"})
		return
	}

	tmpl, err := template.New("email").Parse(hotelEmailHtml)
	if err != nil {
		fmt.Println("Error parsing HTML template:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var emailBody bytes.Buffer
	err = tmpl.Execute(&emailBody, hotelEmailReq.HotelBookingDetails)
	if err != nil {
		fmt.Println("Error executing HTML template:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sendEmail(hotelEmailReq.Subject, emailBody.String(), []string{hotelEmailReq.To}, password)

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Email sent successfully to %s", hotelEmailReq.To)})
}

func sendEmail(subject string, body string, to []string, password string) {
	auth := smtp.PlainAuth("", "journeyendless4@gmail.com", password, "smtp.gmail.com")
	msg := "Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		"<html><body>" + body + "</body></html>"
	err := smtp.SendMail("smtp.gmail.com:587", auth, "journeyendless4@gmail.com", to, []byte(msg))

	if err != nil {
		fmt.Println(err)
	}
}

func VerifierFun(c *gin.Context) {
	var verify Verifier
	err := c.BindJSON(&verify)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode request body"})
		return
	}

	password := os.Getenv("USER_PASSWORD")
	fmt.Println(password)
	if password == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "USER_PASSWORD environment variable is not set"})
		return
	}

	if password != verify.VerifyPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Login successful"})
}

func main() {

	r := gin.Default()

	// CORS middleware
  r.Use(CORSMiddleware())

	// Routes
	r.POST("/sendCarEmail", sendCarEmailHandler)
	r.POST("/sendFightEmail", sendFlightEmailHandler)
	r.POST("/sendHotelEmail", sendHotelEmailHandler)
	r.POST("/login", VerifierFun)

	// Start server
  port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		fmt.Printf("error: %s", err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
