package payment

import (
	"bwastartup/user"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)

type service struct{}

// Service interface untuk mendapatkan token pembayaran.
type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

// NewService mengembalikan instance baru dari service.
func NewService() *service {
	return &service{}
}

// GetToken generates a payment token for Midtrans Snap.
func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	// Inisialisasi Midtrans client
	midClient := midtrans.NewClient()
	midClient.ServerKey = ""
	midClient.ClientKey = ""
	midClient.APIEnvType = midtrans.Sandbox

	// Inisialisasi Snap Gateway
	snapGateway := midtrans.SnapGateway{
		Client: midClient,
	}

	// Membuat Snap request
	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	// Memproses transaksi dan mendapatkan token
	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
