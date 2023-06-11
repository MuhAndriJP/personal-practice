package ewallet

import (
	"context"
	"os"

	entity "github.com/MuhAndriJP/personal-practice.git/entity/xendit"
	"github.com/MuhAndriJP/personal-practice.git/helper"
	"github.com/MuhAndriJP/personal-practice.git/repo/mysql"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
)

type IXendit interface {
	CreateEWalletCharge(context.Context, *entity.CreateEWalletChargeRequest) (*xendit.EWalletCharge, error)
	GetEWalletChargeStatus(context.Context, *entity.GetEWalletChargeStatusRequest) (*xendit.EWalletCharge, error)
	CreateEWalletCallback(context.Context, *entity.CreateEWalletCallbackRequest) error
}

type Xendit struct {
	mysql mysql.IEWallet
}

func NewXenditEwallet() IXendit {
	return &Xendit{
		mysql: mysql.NewEWallet(),
	}
}

func (x *Xendit) CreateEWalletCharge(c context.Context, req *entity.CreateEWalletChargeRequest) (res *xendit.EWalletCharge, err error) {
	res = &xendit.EWalletCharge{}

	xendit.Opt.SecretKey = os.Getenv("XENDIT_SECRET_KEY")

	if req.RedirectURL == "" {
		req.RedirectURL = "https://avicenna.up.railway.app/index"
	}

	if req.CheckoutMethod == "" {
		req.CheckoutMethod = entity.OTP
	}

	charge := ewallet.CreateEWalletChargeParams{
		ReferenceID:    helper.RandomString(10),
		Currency:       entity.IDR,
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

func (x *Xendit) GetEWalletChargeStatus(ctx context.Context, req *entity.GetEWalletChargeStatusRequest) (res *xendit.EWalletCharge, err error) {
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

func (x *Xendit) CreateEWalletCallback(ctx context.Context, req *entity.CreateEWalletCallbackRequest) (err error) {
	callbackData := entity.EWalletPayment{
		ID:             req.CallbackData.ID,
		Status:         req.CallbackData.Status,
		Currency:       req.CallbackData.Currency,
		ChannelCode:    req.CallbackData.ChannelCode,
		ReferenceID:    req.CallbackData.ReferenceID,
		ChargeAmount:   req.CallbackData.ChargeAmount,
		CaptureAmount:  req.CallbackData.CaptureAmount,
		CheckoutMethod: req.CallbackData.CheckoutMethod,
		Event:          req.Event,
		BusinessID:     req.BusinessID,
		Created:        req.CallbackData.Created,
		Updated:        req.CallbackData.Updated,
	}

	err = x.mysql.Insert(ctx, &callbackData)

	return
}
