package workwx

// GetKfServiceState 获取微信客服状态
func (c *WorkwxApp) GetKfServiceState(
	openKfId string,
	externalUserID string,
) (int, string, error) {
	resp, err := c.execGetKfServiceState(reqGetKfServiceState{
		OpenKfId:       openKfId,
		ExternalUserID: externalUserID,
	})
	if err != nil {
		return 0, "", err
	}

	return resp.ServiceState, resp.ServiceUserID, nil
}

// TransKfServiceState 变更会话状态
func (c *WorkwxApp) TransKfServiceState(
	openKfId string,
	externalUserID string,
	serviceState int,
	ServiceUserID string,
) (string, error) {
	resp, err := c.execTransKfServiceState(reqTransKfServiceState{
		OpenKfId:       openKfId,
		ExternalUserID: externalUserID,
		ServiceState:   serviceState,
		ServiceUserID:  ServiceUserID,
	})

	if err != nil {
		return "", err
	}

	return resp.MsgCode, nil
}
