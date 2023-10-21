package workwx

type kfMsgCommon struct {
	MsgId          string        `json:"msgid"`
	OpenKfId       string        `json:"open_kfid"`
	ExternalUserID string        `json:"external_userid"`
	SendTime       uint64        `json:"send_time"`
	Origin         uint32        `json:"origin"`
	ServicerUserId string        `json:"servicer_userid"`
	MsgType        KfMessageType `json:"msgtype"`
	EventType      KfEventType   `json:"event_type"`
}

type KfMessage struct {
	kfMsgCommon

	extras kfMessageDetail
}

func parseKfMsg(msgRaw *kfMsgRaw) (*KfMessage, error) {
	common := kfMsgCommon{
		MsgId:          msgRaw.MsgId,
		OpenKfId:       msgRaw.OpenKfId,
		ExternalUserID: msgRaw.ExternalUserID,
		SendTime:       msgRaw.SendTime,
		Origin:         msgRaw.Origin,
		ServicerUserId: msgRaw.ServicerUserId,
		MsgType:        msgRaw.MsgType,
		EventType:      msgRaw.Event.EventType,
	}
	extras, err := extractKfMessageExtras(common, *msgRaw)
	if err != nil {
		return nil, err
	}
	msg := KfMessage{
		kfMsgCommon: common,
		extras:      extras,
	}
	return &msg, nil
}

func (m *KfMessage) Text() (*KfTextMessage, bool) {
	msg, ok := m.extras.(KfTextMessage)
	return &msg, ok
}

func (m *KfMessage) Image() (*KfImageMessage, bool) {
	msg, ok := m.extras.(KfImageMessage)
	return &msg, ok
}

func (m *KfMessage) Voice() (*KfVoiceMessage, bool) {
	msg, ok := m.extras.(KfVoiceMessage)
	return &msg, ok
}

func (m *KfMessage) Video() (*KfVideoMessage, bool) {
	msg, ok := m.extras.(KfVideoMessage)
	return &msg, ok
}

func (m *KfMessage) File() (*KfFileMessage, bool) {
	msg, ok := m.extras.(KfFileMessage)
	return &msg, ok
}

func (m *KfMessage) Location() (*KfLocationMessage, bool) {
	msg, ok := m.extras.(KfLocationMessage)
	return &msg, ok
}

func (m *KfMessage) Link() (*KfLinkMessage, bool) {
	msg, ok := m.extras.(KfLinkMessage)
	return &msg, ok
}

func (m *KfMessage) BusinessCard() (*KfBusinessCardMessage, bool) {
	msg, ok := m.extras.(KfBusinessCardMessage)
	return &msg, ok
}

func (m *KfMessage) MiniProgram() (*KfMiniProgramMessage, bool) {
	msg, ok := m.extras.(KfMiniProgramMessage)
	return &msg, ok
}

func (m *KfMessage) MsgMenu() (*KfMsgMenuMessage, bool) {
	msg, ok := m.extras.(KfMsgMenuMessage)
	return &msg, ok
}

func (m *KfMessage) ChannelsShopProduct() (*KfChannelsShopProductMessage, bool) {
	msg, ok := m.extras.(KfChannelsShopProductMessage)
	return &msg, ok
}

func (m *KfMessage) ChannelsShopOrder() (*KfChannelsShopOrderMessage, bool) {
	msg, ok := m.extras.(KfChannelsShopOrderMessage)
	return &msg, ok
}

func (m *KfMessage) MergedM() (*KfMergedMessage, bool) {
	msg, ok := m.extras.(KfMergedMessage)
	return &msg, ok
}

func (m *KfMessage) Channels() (*KfChannelsMessage, bool) {
	msg, ok := m.extras.(KfChannelsMessage)
	return &msg, ok
}

func (m *KfMessage) EventEnterSession() (*KfEventEnterSession, bool) {
	msg, ok := m.extras.(KfEventEnterSession)
	return &msg, ok
}

func (m *KfMessage) EventMsgSendFail() (*KfEventMsgSendFail, bool) {
	msg, ok := m.extras.(KfEventMsgSendFail)
	return &msg, ok
}

func (m *KfMessage) EventServicerStatusChange() (*KfEventServicerStatusChange, bool) {
	msg, ok := m.extras.(KfEventServicerStatusChange)
	return &msg, ok
}

func (m *KfMessage) EventSessionStatusChange() (*KfEventSessionStatusChange, bool) {
	msg, ok := m.extras.(KfEventSessionStatusChange)
	return &msg, ok
}

func (m *KfMessage) EventUserRecallMsg() (*KfEventUserRecallMsg, bool) {
	msg, ok := m.extras.(KfEventUserRecallMsg)
	return &msg, ok
}

func (m *KfMessage) EventServicerRecallMsg() (*KfEventServicerRecallMsg, bool) {
	msg, ok := m.extras.(KfEventServicerRecallMsg)
	return &msg, ok
}

func (m *KfMessage) EventRejectCustomerMsgSwitchChange() (*KfEventRejectCustomerMsgSwitchChange, bool) {
	msg, ok := m.extras.(KfEventRejectCustomerMsgSwitchChange)
	return &msg, ok
}

// KfMessageType 消息类型
type KfMessageType string

// KfMessageTypeText 文本消息
const KfMessageTypeText KfMessageType = "text"

// KfMessageTypeImage 图片消息
const KfMessageTypeImage KfMessageType = "image"

// KfMessageTypeVoice 语音消息
const KfMessageTypeVoice KfMessageType = "voice"

// KfMessageTypeVideo 视频消息
const KfMessageTypeVideo KfMessageType = "video"

// KfMessageTypeFile 文件消息
const KfMessageTypeFile KfMessageType = "file"

// KfMessageTypeLocation 位置消息
const KfMessageTypeLocation KfMessageType = "location"

// KfMessageTypeLink 链接消息
const KfMessageTypeLink KfMessageType = "link"

// KfMessageTypeBusinessCard 名片消息
const KfMessageTypeBusinessCard KfMessageType = "business_card"

// KfMessageTypeMiniProgram 小程序消息
const KfMessageTypeMiniProgram KfMessageType = "miniprogram"

// KfMessageTypeMsgMenu 菜单消息
const KfMessageTypeMsgMenu = "msgmenu"

// KfMessageTypeChannelsShopProduct 视频号商品消息
const KfMessageTypeChannelsShopProduct = "channels_shop_product"

// KfMessageTypeChannelsShopOrder 视频号订单消息
const KfMessageTypeChannelsShopOrder = "channels_shop_order"

// KfMessageTypeMergedMsg 合并转发消息
const KfMessageTypeMergedMsg = "merged_msg"

// KfMessageTypeChannels 视频号消息
const KfMessageTypeChannels = "channels"

// KfMessageMeeting 会议消息
const KfMessageMeeting = "meeting"

// KfMessageSchedule 日程消息
const KfMessageSchedule = "schedule"

// KfMessageTypeEvent 事件消息
const KfMessageTypeEvent KfMessageType = "event"

// KfEventType 事件类型
type KfEventType string

// KfEventTypeEnterSession 用户进入会话事件
const KfEventTypeEnterSession KfEventType = "enter_session"

// KfEventTypeMsgSendFail 消息发送失败事件
const KfEventTypeMsgSendFail KfEventType = "msg_send_fail"

// KfEventTypeServicerStatusChange 接待人员接待状态变更事件
const KfEventTypeServicerStatusChange KfEventType = "servicer_status_change"

// KfEventTypeSessionStatusChange 会话状态变更事件
const KfEventTypeSessionStatusChange KfEventType = "session_status_change"

// KfEventTypeUserRecallMsg 用户撤回消息事件
const KfEventTypeUserRecallMsg KfEventType = "user_recall_msg"

// KfEventTypeServicerRecallMsg 接待人员撤回消息事件
const KfEventTypeServicerRecallMsg KfEventType = "servicer_recall_msg"

// KfEventTypeRejectCustomerMsgSwitchChange 拒收客户消息变更事件
const KfEventTypeRejectCustomerMsgSwitchChange KfEventType = "reject_customer_msg_switch_change"

// SyncKfMsg 读取消息
func (c *WorkwxApp) SyncKfMsg(
	cursor string,
	token string,
	limit uint32,
	voiceFormat uint32,
	openKfId string,
) ([]*KfMessage, error) {
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

	msgs := make([]*KfMessage, len(resp.MsgList))
	for i, msg := range resp.MsgList {
		msg, err := parseKfMsg(msg)
		if err != nil {
			return nil, err
		}
		msgs[i] = msg
	}
	return msgs, nil
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
