package workwx

// GetKfAccountList 获取客服列表
func (c *WorkwxApp) GetKfAccountList(offset uint32, limit uint32) ([]KfAccount, error) {
	resp, err := c.execListKfAccount(reqListKfAccount{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}

	return resp.AccountList, nil
}

// GetServicerList 获取接待人员列表
func (c *WorkwxApp) GetServicerList(openKfId string) ([]KfServicer, error) {
	resp, err := c.execListServicer(reqListServicer{
		OpenKfId: openKfId,
	})

	if err != nil {
		return nil, err
	}

	return resp.ServicerList, nil
}

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
