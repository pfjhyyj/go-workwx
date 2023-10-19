package workwx

import "errors"

// SendKfTextMessage 发送微信客服文本消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendKfTextMessage(
	recipient *Recipient,
	content string,
	isSafe bool,
) error {
	return c.sendMessage(recipient, "text", map[string]interface{}{"content": content}, isSafe)
}

// SendKfImageMessage 发送微信客服图片消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendKfImageMessage(
	recipient *Recipient,
	mediaID string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"image",
		map[string]interface{}{
			"media_id": mediaID,
		}, isSafe,
	)
}

// SendKfVoiceMessage 发送微信客服语音消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendKfVoiceMessage(
	recipient *Recipient,
	mediaID string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"voice",
		map[string]interface{}{
			"media_id": mediaID,
		}, isSafe,
	)
}

// SendKfVideoMessage 发送微信客服视频消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendKfVideoMessage(
	recipient *Recipient,
	mediaID string,
	description string,
	title string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"video",
		map[string]interface{}{
			"media_id":    mediaID,
			"description": description, // TODO: 零值
			"title":       title,       // TODO: 零值
		}, isSafe,
	)
}

// SendKfFileMessage 发送微信客服文件消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendKfFileMessage(
	recipient *Recipient,
	mediaID string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"file",
		map[string]interface{}{
			"media_id": mediaID,
		}, isSafe,
	)
}

// SendKfNewsMessage 发送微信客服图文消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendKfNewsMessage(
	recipient *Recipient,
	title string,
	description string,
	url string,
	picURL string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"news",
		map[string]interface{}{
			"title":       title,
			"description": description, // TODO: 零值
			"url":         url,
			"picurl":      picURL, // TODO: 零值
		}, isSafe,
	)
}

// sendKfMessage 发送微信客服消息底层接口
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) sendKfMessage(
	recipient *Recipient,
	msgtype string,
	content map[string]interface{},
	isSafe bool,
) error {
	isApichatSendRequest := false
	if !recipient.isValidForMessageSend() {
		if !recipient.isValidForAppchatSend() {
			// TODO: better error
			return errors.New("recipient invalid for message sending")
		}

		// 发送给群聊
		isApichatSendRequest = true
	}

	req := reqMessage{
		ToUser:  recipient.UserIDs,
		ToParty: recipient.PartyIDs,
		ToTag:   recipient.TagIDs,
		ChatID:  recipient.ChatID,
		AgentID: c.AgentID,
		MsgType: msgtype,
		Content: content,
		IsSafe:  isSafe,
	}

	var resp respMessageSend
	var err error
	if isApichatSendRequest {
		resp, err = c.execAppchatSend(req)
	} else {
		resp, err = c.execMessageSend(req)
	}

	if err != nil {
		return err
	}

	// TODO: what to do with resp?
	_ = resp
	return nil
}
