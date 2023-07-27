package application

import (
	"clockwork-server/config"
	"clockwork-server/domain/model"
	"clockwork-server/domain/repository"
	"errors"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransService interface {
	GenerateSnapUrl(model.Order) (string, error)
	VerifyPayment(data map[string]interface{}) error
}

type midtransService struct {
	client    snap.Client
	midtrans  config.Midtrans
	orderRepo repository.OrderRepository
}

func NewMidtransService(config *config.Config, orderRepo repository.OrderRepository) MidtransService {
	var client snap.Client

	env := midtrans.Sandbox

	if config.Midtrans.IsProduction {
		env = midtrans.Production
	}

	client.New(config.Midtrans.Key, env)
	return &midtransService{
		client:    client,
		midtrans:  config.Midtrans,
		orderRepo: orderRepo,
	}
}

func (m midtransService) GenerateSnapUrl(order model.Order) (string, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.FormatUint(uint64(order.ID), 10),
			GrossAmt: int64(order.GrandTotal),
		},
		EnabledPayments: []snap.SnapPaymentType{
			"gopay",
		},
	}

	snapResp, err := m.client.CreateTransaction(req)
	if err != nil {
		return "", err
	}

	return snapResp.RedirectURL, nil
}

func (m midtransService) VerifyPayment(data map[string]interface{}) error {
	var coreClient coreapi.Client

	env := midtrans.Sandbox
	if m.midtrans.IsProduction {
		env = midtrans.Production
	}

	coreClient.New(m.client.ServerKey, env)

	orderId, exists := data["order_id"].(string)
	if !exists {
		return errors.New("Invalid payload")
	}

	order, err := m.orderRepo.FindById(data["order_id"].(int))
	if err != nil {
		return err
	}

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := coreClient.CheckTransaction(orderId)
	if e != nil {
		return e
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					order.Status = "payment_accept"
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				order.Status = "payment_accept"
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				order.Status = "payment_expired"
			} else if transactionStatusResp.TransactionStatus == "pending" {
				order.Status = "payment_pending"
			}
		}
	}

	_, err = m.orderRepo.Update(order)
	if err != nil {
		return err
	}

	return nil
}
