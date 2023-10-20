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
	//return c.sendMessage(recipient, "text", map[string]interface{}{"content": content}, isSafe)
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
