package workwx

// AddKfAccount 添加客服帐号
func (c *WorkwxApp) AddKfAccount(name string, mediaId string) (string, error) {
	resp, err := c.execAddKfAccount(reqAddKfAccount{
		Name:    name,
		MediaID: mediaId,
	})

	if err != nil {
		return "", err
	}

	return resp.OpenKfId, nil
}

// DelKfAccount 删除客服帐号
func (c *WorkwxApp) DelKfAccount(openKfId string) error {
	_, err := c.execDelKfAccount(reqDelKfAccount{
		OpenKfId: openKfId,
	})

	return err
}

// UpdateKfAccount 更新客服帐号
func (c *WorkwxApp) UpdateKfAccount(openKfId string, name string, mediaId string) error {
	_, err := c.execUpdateKfAccount(reqUpdateKfAccount{
		OpenKfId: openKfId,
		Name:     name,
		MediaID:  mediaId,
	})

	return err
}

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

// AddKfContactWay 获取客服账号链接
func (c *WorkwxApp) AddKfContactWay(openKfId string, scene string) (string, error) {
	resp, err := c.execAddKfContactWay(reqAddKfContactWay{
		OpenKfId: openKfId,
		Scene:    scene,
	})

	if err != nil {
		return "", err
	}

	return resp.Url, nil
}

// AddKfServicer 添加接待人员
func (c *WorkwxApp) AddKfServicer(openKfId string, userIds []string, departmentIds []uint) ([]KfAccountResult, error) {
	resp, err := c.execAddServicer(reqAddServicer{
		OpenKfId:         openKfId,
		UserIdList:       userIds,
		DepartmentIdList: departmentIds,
	})

	if err != nil {
		return nil, err
	}

	return resp.ResultList, nil
}

// DelKfServicer 删除接待人员
func (c *WorkwxApp) DelKfServicer(openKfId string, userIds []string, departmentIds []uint) ([]KfAccountResult, error) {
	resp, err := c.execDelServicer(reqDelServicer{
		OpenKfId:         openKfId,
		UserIdList:       userIds,
		DepartmentIdList: departmentIds,
	})

	if err != nil {
		return nil, err
	}

	return resp.ResultList, nil
}

// GetKfServicerList 获取接待人员列表
func (c *WorkwxApp) GetKfServicerList(openKfId string) ([]KfServicer, error) {
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
