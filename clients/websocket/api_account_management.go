package websocket

import (
	"github.com/BestNathan/deribit-api/pkg/models"
)

func (c *DeribitWSClient) GetAnnouncements() (result []models.Announcement, err error) {
	err = c.Call("public/get_announcements", nil, &result)
	return
}

func (c *DeribitWSClient) ChangeSubaccountName(params *models.ChangeSubaccountNameParams) (result string, err error) {
	err = c.Call("private/change_subaccount_name", params, &result)
	return
}

func (c *DeribitWSClient) CreateSubaccount() (result models.Subaccount, err error) {
	err = c.Call("private/create_subaccount", nil, &result)
	return
}

func (c *DeribitWSClient) DisableTfaForSubaccount(params *models.DisableTfaForSubaccountParams) (result string, err error) {
	err = c.Call("private/disable_tfa_for_subaccount", params, &result)
	return
}

func (c *DeribitWSClient) GetAccountSummary(params *models.GetAccountSummaryParams) (result models.AccountSummary, err error) {
	err = c.Call("private/get_account_summary", params, &result)
	return
}

func (c *DeribitWSClient) GetEmailLanguage() (result string, err error) {
	err = c.Call("private/get_email_language", nil, &result)
	return
}

func (c *DeribitWSClient) GetNewAnnouncements() (result []models.Announcement, err error) {
	err = c.Call("private/get_new_announcements", nil, &result)
	return
}

func (c *DeribitWSClient) GetPosition(params *models.GetPositionParams) (result models.Position, err error) {
	err = c.Call("private/get_position", params, &result)
	return
}

func (c *DeribitWSClient) GetPositions(params *models.GetPositionsParams) (result []models.Position, err error) {
	err = c.Call("private/get_positions", params, &result)
	return
}

func (c *DeribitWSClient) GetSubaccounts(params *models.GetSubaccountsParams) (result []models.Subaccount, err error) {
	err = c.Call("private/get_subaccounts", params, &result)
	return
}

func (c *DeribitWSClient) GetSubaccountsDetails(params *models.GetSubaccountsDetailsParams) (result []models.SubaccountsDetails, err error) {
	err = c.Call("private/get_subaccounts_details", params, &result)
	return
}

func (c *DeribitWSClient) SetAnnouncementAsRead(params *models.SetAnnouncementAsReadParams) (result string, err error) {
	err = c.Call("private/set_announcement_as_read", params, &result)
	return
}

func (c *DeribitWSClient) SetEmailForSubaccount(params *models.SetEmailForSubaccountParams) (result string, err error) {
	err = c.Call("private/set_email_for_subaccount", params, &result)
	return
}

func (c *DeribitWSClient) SetEmailLanguage(params *models.SetEmailLanguageParams) (result string, err error) {
	err = c.Call("private/set_email_language", params, &result)
	return
}

func (c *DeribitWSClient) SetPasswordForSubaccount(params *models.SetPasswordForSubaccountParams) (result string, err error) {
	err = c.Call("private/set_password_for_subaccount", params, &result)
	return
}

func (c *DeribitWSClient) ToggleNotificationsFromSubaccount(params *models.ToggleNotificationsFromSubaccountParams) (result string, err error) {
	err = c.Call("private/toggle_notifications_from_subaccount", params, &result)
	return
}

func (c *DeribitWSClient) ToggleSubaccountLogin(params *models.ToggleSubaccountLoginParams) (result string, err error) {
	err = c.Call("private/toggle_subaccount_login", params, &result)
	return
}
