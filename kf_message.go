package workwx

//type KfSession struct {
//	ToUser   string `json:"touser"`
//	OpenKfId string `json:"open_kfid"`
//}

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

// SyncKfMsg 读取消息
func (c *WorkwxApp) SyncKfMsg(
	cursor string,
	token string,
	limit uint32,
	voiceFormat uint32,
	openKfId string,
) ([]*KfMsg, error) {
	resp, err := c.execSyncKfMsg(reqSyncKfMsg{
		Cursor:      cursor,
		Token:       token,
		Limit:       limit,
		VoiceFormat: voiceFormat,
		OpenKfId:    openKfId,
	})

	if err != nil {
		return nil, err
	}

	return resp.MsgList, nil
}

// SendKfTextMessage 发送微信客服文本消息
func (c *WorkwxApp) SendKfTextMessage(
	toUser string,
	openKfId string,
	msgId string,
	content string,
) error {
	return c.sendKfMessage(
		toUser,
		openKfId,
		msgId,
		"text",
		map[string]interface{}{
			"content": content,
		},
	)
}

// SendKfImageMessage 发送微信客服图片消息
func (c *WorkwxApp) SendKfImageMessage(
	toUser string,
	openKfId string,
	msgId string,
	mediaID string,
) error {
	return c.sendKfMessage(
		toUser,
		openKfId,
		msgId,
		"image",
		map[string]interface{}{
			"media_id": mediaID,
		},
	)
}

// SendKfVoiceMessage 发送微信客服语音消息
func (c *WorkwxApp) SendKfVoiceMessage(
	toUser string,
	openKfId string,
	msgId string,
	mediaID string,
) error {
	return c.sendKfMessage(
		toUser,
		openKfId,
		msgId,
		"voice",
		map[string]interface{}{
			"media_id": mediaID,
		},
	)
}

// SendKfVideoMessage 发送微信客服视频消息
func (c *WorkwxApp) SendKfVideoMessage(
	toUser string,
	openKfId string,
	msgId string,
	mediaID string,
) error {
	return c.sendKfMessage(
		toUser,
		openKfId,
		msgId,
		"video",
		map[string]interface{}{
			"media_id": mediaID,
		},
	)
}

// SendKfFileMessage 发送微信客服文件消息
func (c *WorkwxApp) SendKfFileMessage(
	toUser string,
	openKfId string,
	msgId string,
	mediaID string,
) error {
	return c.sendKfMessage(
		toUser,
		openKfId,
		msgId,
		"file",
		map[string]interface{}{
			"media_id": mediaID,
		},
	)
}

// SendKfLinkMessage 发送微信客服图文链接消息
func (c *WorkwxApp) SendKfLinkMessage(
	toUser string,
	openKfId string,
	msgId string,
	title string,
	desc string,
	url string,
	thumbUrl string,
) error {
	return c.sendKfMessage(
		toUser,
		openKfId,
		msgId,
		"link",
		map[string]interface{}{
			"title":          title,
			"desc":           desc,
			"url":            url,
			"thumb_media_id": thumbUrl,
		},
	)
}

// SendKfMiniProgramMessage 发送微信客服小程序卡片消息
func (c *WorkwxApp) SendKfMiniProgramMessage(
	toUser string,
	openKfId string,
	msgId string,
	title string,
	thumbUrl string,
	appid string,
	pagepath string,
) error {
	return c.sendKfMessage(
		toUser,
		openKfId,
		msgId,
		"miniprogram",
		map[string]interface{}{
			"title":          title,
			"thumb_media_id": thumbUrl,
			"appid":          appid,
			"pagepath":       pagepath,
		},
	)
}

// SendKfMenuMessage 发送微信客服菜单消息
func (c *WorkwxApp) SendKfMenuMessage(
	toUser string,
	openKfId string,
	msgId string,
	headContent string,
	list []KfMsgMenu,
	tailContent string,
) error {
	return c.sendKfMessage(
		toUser,
		openKfId,
		msgId,
		"menu",
		map[string]interface{}{
			"head_content": headContent,
			"list":         list,
			"tail_content": tailContent,
		},
	)
}

// SendKfLocationMessage 发送微信客服地理位置消息
func (c *WorkwxApp) SendKfLocationMessage(
	toUser string,
	openKfId string,
	msgId string,
	name string,
	address string,
	latitude float64,
	longitude float64,
) error {
	return c.sendKfMessage(
		toUser,
		openKfId,
		msgId,
		"location",
		map[string]interface{}{
			"name":      name,
			"address":   address,
			"latitude":  latitude,
			"longitude": longitude,
		},
	)
}

// sendKfMessage 发送微信客服消息底层接口
func (c *WorkwxApp) sendKfMessage(
	toUser string,
	openKfId string,
	msgId string,
	msgType string,
	content map[string]interface{},
) error {
	req := reqSendKfMsg{
		ToUser:   toUser,
		OpenKfId: openKfId,
		MsgId:    msgId,
		MsgType:  msgType,
		Content:  content,
	}

	resp, err := c.execSendKfMsg(req)

	if err != nil {
		return err
	}

	// TODO: what to do with resp?
	_ = resp
	return nil
}

// SendKfMsgOnEvent 发送欢迎语等事件响应消息
func (c *WorkwxApp) SendKfMsgOnEvent(
	code string,
	msgId string,
	msgType string,
) (string, error) {
	resp, err := c.execSendKfMsgOnEvent(reqSendKfMsgOnEvent{
		Code:    code,
		MsgId:   msgId,
		MsgType: msgType,
	})

	if err != nil {
		return "", err
	}

	return resp.MsgId, nil
}
