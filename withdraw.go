package gdax

import (
	"log"
	"fmt"
)

type CryptoWithdrawalRequest struct {
	Amount      float64 `json:"amount"`
	Currency      string `json:"currency"`
	CryptoAddress      string `json:"crypto_address"`
}

/*
sample
{
    "amount": 10.00,
    "currency": "BTC",
    "crypto_address": "0x5ad5769cd04681FeD900BCE3DDc877B50E83d469"
}
*/


type CryptoWithdrawalResponse struct {
	Id      string `json:"id"`
	Amount      string `json:"amount"`
	Currency      string `json:"currency"`
}

type PaymentMethodWithdrawalRequest struct {
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
	PaymentMethodId string `json:"payment_method_id"`
}

/*
{
    "id":"593533d2-ff31-46e0-b22e-ca754147a96a",
    "amount": "10.00",
    "currency": "USD",
    "payout_at": "2016-08-20T00:31:09Z"
}
 */
type PaymentMethodWithdrawalResponse struct {
	Id string `json:"id"`
	Amount string `json:"amount"`
	Currency string `json:"currency"`
	PayoutAt string `json:"payout_at"`
}


/*
sample
{
    "id":"593533d2-ff31-46e0-b22e-ca754147a96a",
    "amount":"10.00",
    "currency": "BTC",
}
*/

// go get test

func (c *Client) WithdrawCrypto(amount string, currency string, cryptoAddress string) (CryptoWithdrawalResponse, error) {
	amountFloat := StringToFloat(amount)
	newWithdrawalReq := CryptoWithdrawalRequest{amountFloat, currency, cryptoAddress}

	log.Printf("DEBUG > newWithdrawalReq %+v", newWithdrawalReq)

	resp := CryptoWithdrawalResponse{}

	url := fmt.Sprintf("/withdrawals/crypto")
	_, err := c.Request("POST", url, newWithdrawalReq, &resp)
	return resp, err
}

func (c *Client) WithdrawToPaymentMethods(payload PaymentMethodWithdrawalRequest) (PaymentMethodWithdrawalResponse, error) {
	// /withdrawals/payment-method
	/*
	{
    "amount": 10.00,
    "currency": "USD",
    "payment_method_id": "bc677162-d934-5f1a-968c-a496b1c1270b"
	 */
	resp := PaymentMethodWithdrawalResponse{}

	url := fmt.Sprintf("/withdrawals/payment-method")
	_, err := c.Request("POST", url, payload, &resp)

	return resp, err
}
