package xendit

import (
	"context"
	"os"

	"github.com/MuhAndriJP/gateway-service.git/helper"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
)

type IXendit interface {
	CreateEWalletCharge(context.Context, *CreateEWalletChargeRequest) (*xendit.EWalletCharge, error)
	GetEWalletChargeStatus(context.Context, *GetEWalletChargeStatusRequest) (*xendit.EWalletCharge, error)
	CreateEWalletCallback(context.Context) error
}

type Xendit struct {
}

func NewXendit() IXendit {
	return &Xendit{}
}

func (x *Xendit) CreateEWalletCharge(c context.Context, req *CreateEWalletChargeRequest) (res *xendit.EWalletCharge, err error) {
	res = &xendit.EWalletCharge{}

	xendit.Opt.SecretKey = os.Getenv("XENDIT_SECRET_KEY")

	if req.RedirectURL == "" {
		req.RedirectURL = "https://avicenna.up.railway.app/index"
	}

	if req.CheckoutMethod == "" {
		req.CheckoutMethod = OTP
	}

	charge := ewallet.CreateEWalletChargeParams{
		ReferenceID:    helper.RandomString(10),
		Currency:       IDR,
		Amount:         req.Amount,
		CheckoutMethod: req.CheckoutMethod,
		ChannelCode:    req.ChannelCode,
		ChannelProperties: map[string]string{
			"success_redirect_url": req.RedirectURL,
		},
		// Customer: &xendit.EwalletCustomer{},
		// Basket:         []xendit.EWalletBasketItem{},
	}

	resp, errEwallet := ewallet.CreateEWalletCharge(&charge)
	if errEwallet != nil {
		err = errEwallet
		return
	}

	res = resp

	return
}

func (x *Xendit) GetEWalletChargeStatus(ctx context.Context, req *GetEWalletChargeStatusRequest) (res *xendit.EWalletCharge, err error) {
	res = &xendit.EWalletCharge{}

	xendit.Opt.SecretKey = os.Getenv("XENDIT_SECRET_KEY")

	data := ewallet.GetEWalletChargeStatusParams{
		ChargeID: req.ChargeID,
	}

	resp, errEwallet := ewallet.GetEWalletChargeStatus(&data)
	if errEwallet != nil {
		err = errEwallet
		return
	}

	res = resp

	return
}

func (x *Xendit) CreateEWalletCallback(ctx context.Context) (err error) {

	return
}
