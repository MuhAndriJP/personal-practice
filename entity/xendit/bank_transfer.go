package xendit

import (
	"time"

	"github.com/xendit/xendit-go"
)

const (
	WHATSAPP = "whatsapp"
	EMAIL    = "email"
	SMS      = "sms"
)

type CreateInvoiceRequest struct {
	ForUserID                      string                                       `json:"-"`
	ExternalID                     string                                       `json:"external_id" validate:"required"`
	Amount                         float64                                      `json:"amount" validate:"required"`
	Description                    string                                       `json:"description,omitempty"`
	PayerEmail                     string                                       `json:"payer_email,omitempty"`
	ShouldSendEmail                *bool                                        `json:"should_send_email,omitempty"`
	Customer                       xendit.InvoiceCustomer                       `json:"customer,omitempty"`
	CustomerAddress                xendit.InvoiceCustomerAddress                `json:"addresses,omitempty"`
	CustomerNotificationPreference xendit.InvoiceCustomerNotificationPreference `json:"customer_notification_preference,omitempty"`
	InvoiceDuration                int                                          `json:"invoice_duration,omitempty"`
	SuccessRedirectURL             string                                       `json:"success_redirect_url,omitempty"`
	FailureRedirectURL             string                                       `json:"failure_redirect_url,omitempty"`
	PaymentMethods                 []string                                     `json:"payment_methods,omitempty"`
	Currency                       string                                       `json:"currency,omitempty"`
	FixedVA                        *bool                                        `json:"fixed_va,omitempty"`
	CallbackVirtualAccountID       string                                       `json:"callback_virtual_account_id,omitempty"`
	MidLabel                       string                                       `json:"mid_label,omitempty"`
	ReminderTimeUnit               string                                       `json:"reminder_time_unit,omitempty"`
	ReminderTime                   int                                          `json:"reminder_time,omitempty"`
	Locale                         string                                       `json:"locale,omitempty"`
	Items                          []xendit.InvoiceItem                         `json:"items,omitempty"`
	Fees                           []xendit.InvoiceFee                          `json:"fees,omitempty"`
	ShouldAuthenticateCreditCard   bool
}

type CreateInvoiceResponse struct {
	ID                        string                       `json:"id"`
	UserID                    string                       `json:"user_id"`
	ExternalID                string                       `json:"external_id"`
	Status                    string                       `json:"status"`
	MerchantName              string                       `json:"merchant_name"`
	MerchantProfilePictureURL string                       `json:"merchant_profile_picture_url"`
	Amount                    int                          `json:"amount"`
	PayerEmail                string                       `json:"payer_email"`
	Description               string                       `json:"description"`
	InvoiceURL                string                       `json:"invoice_url"`
	ExpiryDate                time.Time                    `json:"expiry_date"`
	AvailableBanks            []xendit.InvoiceBank         `json:"available_banks"`
	AvailableRetailOutlets    []xendit.InvoiceRetailOutlet `json:"available_retail_outlets"`
	ShouldExcludeCreditCard   bool                         `json:"should_exclude_credit_card"`
	ShouldSendEmail           bool                         `json:"should_send_email"`
	Created                   time.Time                    `json:"created"`
	Updated                   time.Time                    `json:"updated"`
	MidLabel                  string                       `json:"mid_label"`
	Currency                  string                       `json:"currency"`
	FixedVa                   bool                         `json:"fixed_va"`
	Locale                    string                       `json:"locale"`
	Customer                  xendit.InvoiceCustomer       `json:"customer"`
	Items                     []xendit.InvoiceItem         `json:"items"`
	Fees                      []xendit.InvoiceFee          `json:"fees"`
}
