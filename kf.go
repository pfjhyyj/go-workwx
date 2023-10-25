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

// GetKfServiceState 获取会话状态
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

// GetKfUpgradeServiceConfig 获取配置的专员与客户群
func (c *WorkwxApp) GetKfUpgradeServiceConfig() (*KfMemberRange, *KfGroupChatRange, error) {
	resp, err := c.execGetUpgradeServiceConfig(reqGetUpgradeServiceConfig{})
	if err != nil {
		return nil, nil, err
	}

	return &resp.MemberRange, &resp.GroupChatRange, nil
}

// SetKfUpgradeServiceMember 为客户升级为专员
func (c *WorkwxApp) SetKfUpgradeServiceMember(openKfId string, externalUserID string, userID string, wording string) error {
	_, err := c.execUpgradeService(reqUpgradeService{
		OpenKfId:       openKfId,
		ExternalUserID: externalUserID,
		Type:           1,
		Member: &KfUpgradeMember{
			UserID:  userID,
			Wording: wording,
		},
	})

	return err
}

// SetKfUpgradeServiceGroupChat 为客户升级为客户群
func (c *WorkwxApp) SetKfUpgradeServiceGroupChat(openKfId string, externalUserID string, chatID string, wording string) error {
	_, err := c.execUpgradeService(reqUpgradeService{
		OpenKfId:       openKfId,
		ExternalUserID: externalUserID,
		Type:           2,
		GroupChat: &KfUpgradeGroupChat{
			ChatID:  chatID,
			Wording: wording,
		},
	})

	return err
}

// CancelKfUpgradeService 为客户取消推荐
func (c *WorkwxApp) CancelKfUpgradeService(openKfId string, externalUserId string) error {
	_, err := c.execUpgradeService(reqUpgradeService{
		OpenKfId:       openKfId,
		ExternalUserID: externalUserId,
	})

	return err
}

// GetKfCustomers 获取客户基础信息
func (c *WorkwxApp) GetKfCustomers(externalUserIDs []string, needEnterSessionContext uint8) ([]KfCustomer, error) {
	resp, err := c.execBatchGetCustomer(reqBatchGetCustomer{
		ExternalUserIDList:      externalUserIDs,
		NeedEnterSessionContext: needEnterSessionContext,
	})

	if err != nil {
		return nil, err
	}

	return resp.CustomerList, nil
}

// GetKfCorpStatistic 获取「客户数据统计」企业汇总数据
func (c *WorkwxApp) GetKfCorpStatistic(startTime uint32, endTime uint32) ([]*KfCorpStatisticDetail, error) {
	resp, err := c.execGetCorpStatistic(reqGetCorpStatistic{
		StartTime: startTime,
		EndTime:   endTime,
	})

	if err != nil {
		return nil, err
	}

	return resp.StatisticList, nil
}

// GetKfServicerStatistic 获取「客户数据统计」接待人员明细数据
func (c *WorkwxApp) GetKfServicerStatistic(openKfId string, startTime uint32, endTime uint32) ([]*KfServicerStatisticDetail, error) {
	resp, err := c.execGetServicerStatistic(reqGetServicerStatistic{
		StartTime: startTime,
		EndTime:   endTime,
		OpenKfId:  openKfId,
	})

	if err != nil {
		return nil, err
	}

	return resp.StatisticList, nil
}

// AddKfKnowledgeGroup 添加知识库分组
func (c *WorkwxApp) AddKfKnowledgeGroup(name string) (string, error) {
	resp, err := c.execAddKnowledgeGroup(reqAddKnowledgeGroup{
		Name: name,
	})

	if err != nil {
		return "", err
	}

	return resp.GroupId, nil
}

// DelKfKnowledgeGroup 删除知识库分组
func (c *WorkwxApp) DelKfKnowledgeGroup(groupId string) error {
	_, err := c.execDelKnowledgeGroup(reqDelKnowledgeGroup{
		GroupId: groupId,
	})

	return err
}

// UpdateKfKnowledgeGroup 更新知识库分组
func (c *WorkwxApp) UpdateKfKnowledgeGroup(groupId string, name string) error {
	_, err := c.execModKnowledgeGroup(reqModKnowledgeGroup{
		GroupId: groupId,
		Name:    name,
	})

	return err
}

// GetKfKnowledgeGroupList 获取知识库分组
func (c *WorkwxApp) GetKfKnowledgeGroupList(cursor string, limit uint32, groupId string) ([]*KfKnowledgeGroup, error) {
	resp, err := c.execGetKnowledgeGroupList(reqGetKnowledgeGroupList{
		Cursor:  cursor,
		Limit:   limit,
		GroupId: groupId,
	})
	if err != nil {
		return nil, err
	}

	return resp.GroupList, nil
}

// AddKfKnowledgeIntent 添加知识库问答
func (c *WorkwxApp) AddKfKnowledgeIntent(
	groupId string,
	question *KfQuestion,
	similarQuestions *KfSimilarQuestions,
	answers []*KfQuestionAnswer,
) (string, error) {
	resp, err := c.execAddKnowledgeIntent(reqAddKnowledgeIntent{
		GroupId:          groupId,
		Question:         question,
		SimilarQuestions: similarQuestions,
		Answers:          answers,
	})

	if err != nil {
		return "", err
	}

	return resp.IntentId, nil
}

// DelKfKnowledgeIntent 删除知识库问答
func (c *WorkwxApp) DelKfKnowledgeIntent(intentId string) error {
	_, err := c.execDelKnowledgeIntent(reqDelKnowledgeIntent{
		IntentId: intentId,
	})

	return err
}

// UpdateKfKnowledgeIntent 更新知识库问答
func (c *WorkwxApp) UpdateKfKnowledgeIntent(
	intentId string,
	question *KfQuestion,
	similarQuestions *KfSimilarQuestions,
	answers []*KfQuestionAnswer,
) error {
	_, err := c.execModKnowledgeIntent(reqModKnowledgeIntent{
		IntentId:         intentId,
		Question:         question,
		SimilarQuestions: similarQuestions,
		Answers:          answers,
	})

	return err
}

// GetKfKnowledgeIntentList 获取知识库问答
func (c *WorkwxApp) GetKfKnowledgeIntentList(
	cursor string,
	limit uint32,
	groupId string,
	intentId string,
) ([]*KfKnowledgeIntent, error) {
	resp, err := c.execGetKnowledgeIntentList(reqGetKnowledgeIntentList{
		Cursor:   cursor,
		Limit:    limit,
		GroupId:  groupId,
		IntentId: intentId,
	})
	if err != nil {
		return nil, err
	}

	return resp.IntentList, nil
}
