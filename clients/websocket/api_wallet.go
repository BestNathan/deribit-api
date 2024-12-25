package websocket

import (
	"github.com/BestNathan/deribit-api/pkg/models"
)

func (c *DeribitWSClient) CancelTransferByID(params *models.CancelTransferByIDParams) (result models.Transfer, err error) {
	err = c.Call("private/cancel_transfer_by_id", params, &result)
	return
}

func (c *DeribitWSClient) CancelWithdrawal(params *models.CancelWithdrawalParams) (result models.Withdrawal, err error) {
	err = c.Call("private/cancel_withdrawal", params, &result)
	return
}

func (c *DeribitWSClient) CreateDepositAddress(params *models.CreateDepositAddressParams) (result models.DepositAddress, err error) {
	err = c.Call("private/create_deposit_address", params, &result)
	return
}

func (c *DeribitWSClient) GetCurrentDepositAddress(params *models.GetCurrentDepositAddressParams) (result models.DepositAddress, err error) {
	err = c.Call("private/get_current_deposit_address", params, &result)
	return
}

func (c *DeribitWSClient) GetDeposits(params *models.GetDepositsParams) (result models.GetDepositsResponse, err error) {
	err = c.Call("private/get_deposits", params, &result)
	return
}

func (c *DeribitWSClient) GetTransfers(params *models.GetTransfersParams) (result models.GetTransfersResponse, err error) {
	err = c.Call("private/get_transfers", params, &result)
	return
}

func (c *DeribitWSClient) GetWithdrawals(params *models.GetWithdrawalsParams) (result []models.Withdrawal, err error) {
	err = c.Call("private/get_withdrawals", params, &result)
	return
}

func (c *DeribitWSClient) Withdraw(params *models.WithdrawParams) (result models.Withdrawal, err error) {
	err = c.Call("private/withdraw", params, &result)
	return
}
