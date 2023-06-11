package xendit

import (
	"time"

	"gorm.io/gorm"
)

const (
	IDR = "IDR"

	OTP        = "ONE_TIME_PAYMENT"
	SUBSCRIBED = "TOKENIZED_PAYMENT"

	OVO      = "ID_OVO"
	DANA     = "ID_DANA"
	LINK_AJA = "ID_LINKAJA"
	SHOPEE   = "ID_SHOPEEPAY"
	ASTRA    = "ID_ASTRAPAY"
	JENIUS   = "ID_JENIUSPAY"
	SAKUKU   = "ID_SAKUKU"
)

type CreateEWalletChargeRequest struct {
	Amount         float64 `json:"amount"`
	ChannelCode    string  `json:"channel_code"`
	RedirectURL    string  `json:"redirect_url"`
	CheckoutMethod string  `json:"checkout_method"`
}

type GetEWalletChargeStatusRequest struct {
	ChargeID string `json:"charge_id"`
}

type CreateEWalletCallbackRequest struct {
	CallbackData Data      `json:"data"`
	Event        string    `json:"event"`
	Created      time.Time `json:"created"`
	BusinessID   string    `json:"business_id"`
}

type Data struct {
	ID                 string            `json:"id"`
	Basket             interface{}       `json:"basket"`
	Status             string            `json:"status"`
	Actions            Actions           `json:"actions"`
	Created            time.Time         `json:"created"`
	Updated            time.Time         `json:"updated"`
	Currency           string            `json:"currency"`
	Metadata           Metadata          `json:"metadata"`
	VoidedAt           interface{}       `json:"voided_at"`
	CaptureNow         bool              `json:"capture_now"`
	CustomerID         interface{}       `json:"customer_id"`
	CallbackURL        string            `json:"callback_url"`
	ChannelCode        string            `json:"channel_code"`
	FailureCode        interface{}       `json:"failure_code"`
	ReferenceID        string            `json:"reference_id"`
	ChargeAmount       int               `json:"charge_amount"`
	CaptureAmount      int               `json:"capture_amount"`
	CheckoutMethod     string            `json:"checkout_method"`
	PaymentMethodID    interface{}       `json:"payment_method_id"`
	ChannelProperties  ChannelProperties `json:"channel_properties"`
	IsRedirectRequired bool              `json:"is_redirect_required"`
}

type Actions struct {
	MobileWebCheckoutURL      string `json:"mobile_web_checkout_url"`
	DesktopWebCheckoutURL     string `json:"desktop_web_checkout_url"`
	MobileDeeplinkCheckoutURL string `json:"mobile_deeplink_checkout_url"`
}

type Metadata struct {
	BranchCode string `json:"branch_code"`
}

type ChannelProperties struct {
	SuccessRedirectURL string `json:"success_redirect_url"`
}

type EWalletPayment struct {
	gorm.Model
	ID             string    `json:"id"`
	Status         string    `json:"status"`
	Created        time.Time `json:"created"`
	Updated        time.Time `json:"updated"`
	Currency       string    `json:"currency"`
	CallbackURL    string    `json:"callback_url"`
	ChannelCode    string    `json:"channel_code"`
	ReferenceID    string    `json:"reference_id"`
	ChargeAmount   int       `json:"charge_amount"`
	CaptureAmount  int       `json:"capture_amount"`
	CheckoutMethod string    `json:"checkout_method"`
	Event          string    `json:"event"`
	BusinessID     string    `json:"business_id"`
}
