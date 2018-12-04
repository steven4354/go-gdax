package gdax

import (
	"fmt"
)

type CryptoWithdrawal struct {
	Amount      string `json:"amount"`
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

/*
sample
{
    "id":"593533d2-ff31-46e0-b22e-ca754147a96a",
    "amount":"10.00",
    "currency": "BTC",
}
*/

func (c *Client) WithdrawCrypto(newWithdrawal *CryptoWithdrawal) (CryptoWithdrawalResponse, error) {
	resp := CryptoWithdrawalResponse{}

	url := fmt.Sprintf("/withdrawals/crypto")
	_, err := c.Request("POST", url, newWithdrawal, &resp)
	return resp, err
}
