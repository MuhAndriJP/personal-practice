package bank_transfer

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	entity "github.com/MuhAndriJP/personal-practice.git/entity/xendit"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type IXenditBankTransfer interface {
	CreateInvoice(context.Context, *entity.CreateInvoiceRequest) (res *xendit.Invoice, err error)
}

type xenditBankTransfer struct{}

func NewXenditBankTransfer() IXenditBankTransfer {
	return &xenditBankTransfer{}
}

func (x *xenditBankTransfer) CreateInvoice(ctx context.Context, req *entity.CreateInvoiceRequest) (res *xendit.Invoice, err error) {
	customerAddress := xendit.InvoiceCustomerAddress{
		Country:     req.CustomerAddress.Country,
		StreetLine1: req.CustomerAddress.StreetLine1,
		StreetLine2: req.CustomerAddress.StreetLine2,
		City:        req.CustomerAddress.City,
		State:       req.CustomerAddress.State,
		PostalCode:  req.CustomerAddress.PostalCode,
	}

	customer := xendit.InvoiceCustomer{
		GivenNames:   req.Customer.GivenNames,
		Surname:      req.Customer.Surname,
		Email:        req.Customer.Email,
		MobileNumber: req.Customer.MobileNumber,
		Address:      []xendit.InvoiceCustomerAddress{customerAddress},
	}

	items := make([]xendit.InvoiceItem, 0)
	for _, v := range req.Items {
		items = append(items, xendit.InvoiceItem{
			Name:     v.Name,
			Quantity: v.Quantity,
			Price:    v.Price,
			Category: v.Category,
		})
	}

	fees := make([]xendit.InvoiceFee, 0)
	for _, v := range req.Fees {
		fees = append(fees, xendit.InvoiceFee{
			Type:  v.Type,
			Value: v.Value,
		})
	}

	NotificationType := []string{entity.WHATSAPP, entity.EMAIL, entity.SMS}

	customerNotificationPreference := xendit.InvoiceCustomerNotificationPreference{
		InvoiceCreated:  NotificationType,
		InvoiceReminder: NotificationType,
		InvoicePaid:     NotificationType,
		InvoiceExpired:  NotificationType,
	}

	data := invoice.CreateParams{
		ExternalID:                     req.ExternalID,
		Amount:                         req.Amount,
		Description:                    req.Description,
		InvoiceDuration:                req.InvoiceDuration * 24,
		Customer:                       customer,
		CustomerNotificationPreference: customerNotificationPreference,
		SuccessRedirectURL:             "https://www.google.com",
		FailureRedirectURL:             "https://www.google.com",
		Currency:                       entity.IDR,
		Items:                          items,
		Fees:                           fees,
	}

	payload, errMarshal := json.Marshal(data)
	if err != nil {
		err = errMarshal
		return
	}

	apiKey := os.Getenv("XENDIT_BASE64")

	reqHttp, errReq := http.NewRequest("POST", os.Getenv("XENDIT_URL_INVOICE"), bytes.NewBuffer(payload))
	reqHttp.Header.Set("Content-Type", "application/json")
	reqHttp.Header.Set("Authorization", "Basic "+apiKey)
	if err != nil {
		err = errReq
		return
	}

	client := http.Client{}
	resp, errDo := client.Do(reqHttp)
	if errDo != nil {
		err = errDo
		return
	}

	defer resp.Body.Close()
	body, errBody := io.ReadAll(resp.Body)
	if err != nil {
		err = errBody
		return
	}

	errUnmarshal := json.Unmarshal(body, &res)
	if err != nil {
		err = errUnmarshal
		return
	}

	return
}
