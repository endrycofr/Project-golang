package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {

	formatter := CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt

	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return nil // instead of an empty slice, return nil to indicate no results
	}

	formatters := make([]CampaignTransactionFormatter, 0, len(transactions)) // pre-allocate the slice

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		formatters = append(formatters, formatter)
	}

	return formatters
}
