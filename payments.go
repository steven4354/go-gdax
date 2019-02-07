package gdax

import "fmt"

type Remaining struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Total struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Deposit struct {
	PeriodInDays int `json:"period_in_days"`
	Total Total `json:"total"`
	Remaining Remaining `json:"remaining"`
}

type Sell struct {
	PeriodInDays int `json:"period_in_days"`
	Total Total `json:"total"`
	Remaining Remaining `json:"remaining"`
}

type InstantBuy struct {
	PeriodInDays int `json:"period_in_days"`
	Total Total `json:"total"`
	Remaining Remaining `json:"remaining"`
}

type Buy struct {
	PeriodInDays int `json:"period_in_days"`
	Total Total `json:"total"`
	Remaining Remaining `json:"remaining"`
}

type Limits struct {
	Buy 			 []Buy         `json:"buy"`
	InstantBuy []InstantBuy  `json:"instant_buy"`
	Sell 			 []Sell 			 `json:"sell"`
	Deposit 	 []Deposit 	   `json:"deposit"`
}

type PaymentMethod struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	Currency      string `json:"currency"`
	PrimaryBuy    bool   `json:"primary_buy"`
	PrimarySell   bool   `json:"primary_sell"`
	AllowBuy      bool   `json:"allow_buy"`
	AllowSell     bool   `json:"allow_sell"`
	AllowDeposit  bool   `json:"allow_deposit"`
	AllowWithdraw bool   `json:"allow_withdraw"`
	Limits        Limits `json:"limits"`
}

func (c *Client) GetPaymentMethods() ([]PaymentMethod, error) {
	var paymentMethods []PaymentMethod

	url := fmt.Sprintf("/payment-methods")
	_, err := c.Request("GET", url, nil, &paymentMethods)
	return paymentMethods, err
}