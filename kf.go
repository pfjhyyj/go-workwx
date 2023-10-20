package workwx

type KfSession struct {
	ToUser   string `json:"touser"`
	OpenKFID string `json:"open_kfid"`
}

// SendKfTextMessage 发送微信客服文本消息
func (c *WorkwxApp) SendKfTextMessage(
	session *KfSession,
	msgId string,
	content string,
) error {
	return c.sendKfMessage(
		session,
		msgId,
		"text",
		map[string]interface{}{
			"content": content,
		},
	)
}

// SendKfImageMessage 发送微信客服图片消息
func (c *WorkwxApp) SendKfImageMessage(
	session *KfSession,
	msgId string,
	mediaID string,
) error {
	return c.sendKfMessage(
		session,
		msgId,
		"image",
		map[string]interface{}{
			"media_id": mediaID,
		},
	)
}

// SendKfVoiceMessage 发送微信客服语音消息
func (c *WorkwxApp) SendKfVoiceMessage(
	session *KfSession,
	msgId string,
	mediaID string,
) error {
	return c.sendKfMessage(
		session,
		msgId,
		"voice",
		map[string]interface{}{
			"media_id": mediaID,
		},
	)
}

// SendKfVideoMessage 发送微信客服视频消息
func (c *WorkwxApp) SendKfVideoMessage(
	session *KfSession,
	msgId string,
	mediaID string,
) error {
	return c.sendKfMessage(
		session,
		msgId,
		"video",
		map[string]interface{}{
			"media_id": mediaID,
		},
	)
}

// SendKfFileMessage 发送微信客服文件消息
func (c *WorkwxApp) SendKfFileMessage(
	session *KfSession,
	msgId string,
	mediaID string,
) error {
	return c.sendKfMessage(
		session,
		msgId,
		"file",
		map[string]interface{}{
			"media_id": mediaID,
		},
	)
}

// SendKfLinkMessage 发送微信客服图文链接消息
func (c *WorkwxApp) SendKfLinkMessage(
	session *KfSession,
	msgId string,
	title string,
	desc string,
	url string,
	thumbUrl string,
) error {
	return c.sendKfMessage(
		session,
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
	session *KfSession,
	msgId string,
	title string,
	thumbUrl string,
	appid string,
	pagepath string,
) error {
	return c.sendKfMessage(
		session,
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
	session *KfSession,
	msgId string,
	headContent string,
	list []KfMsgMenu,
	tailContent string,
) error {
	return c.sendKfMessage(
		session,
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
	session *KfSession,
	msgId string,
	name string,
	address string,
	latitude float64,
	longitude float64,
) error {
	return c.sendKfMessage(
		session,
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
	session *KfSession,
	msgId string,
	msgType string,
	content map[string]interface{},
) error {
	req := reqSendKfMsg{
		ToUser:   session.ToUser,
		OpenKfId: session.OpenKFID,
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
