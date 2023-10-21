package workwx

import (
	"fmt"
)

type kfMessageDetail interface {
}

type kfTextMessage struct {
	Content string `json:"content"`
	MenuId  string `json:"menu_id"`
}

var _ kfMessageDetail = (*kfTextMessage)(nil)

type kfImageMessage struct {
	MediaId string `json:"media_id"`
}

var _ kfMessageDetail = (*kfImageMessage)(nil)

type kfVoiceMessage struct {
	MediaId string `json:"media_id"`
}

var _ kfMessageDetail = (*kfVoiceMessage)(nil)

type kfVideoMessage struct {
	MediaId string `json:"media_id"`
}

var _ kfMessageDetail = (*kfVideoMessage)(nil)

type kfFileMessage struct {
	MediaId string `json:"media_id"`
}

var _ kfMessageDetail = (*kfFileMessage)(nil)

type kfLocationMessage struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Name      string `json:"name"`
	Address   string `json:"address"`
}

var _ kfMessageDetail = (*kfLocationMessage)(nil)

type kfLinkMessage struct {
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Url    string `json:"url"`
	PicUrl string `json:"pic_url"`
}

var _ kfMessageDetail = (*kfLinkMessage)(nil)

type kfBusinessCardMessage struct {
	UserId string `json:"userid"`
}

var _ kfMessageDetail = (*kfBusinessCardMessage)(nil)

type kfMiniProgramMessage struct {
	Title        string `json:"title"`
	Appid        string `json:"appid"`
	PagePath     string `json:"pagepath"`
	ThumbMediaId string `json:"thumb_media_id"`
}

var _ kfMessageDetail = (*kfMiniProgramMessage)(nil)

type kfMsgMenuMessage struct {
	HeadContent string `json:"head_content"`
	List        []struct {
		Type  string `json:"type"`
		Click struct {
			Id      string `json:"id"`
			Content string `json:"content"`
		} `json:"click"`
		View struct {
			Url     string `json:"url"`
			Content string `json:"content"`
		} `json:"view"`
		MiniProgram struct {
			Appid    string `json:"appid"`
			Pagepath string `json:"pagepath"`
			Content  string `json:"content"`
		}
	} `json:"list"`
	TailContent string `json:"tail_content"`
}

var _ kfMessageDetail = (*kfMsgMenuMessage)(nil)

type kfChannelsShopProductMessage struct {
	ProductId     string `json:"product_id"`
	HeadImage     string `json:"head_image"`
	Title         string `json:"title"`
	SalesPrice    string `json:"sales_price"`
	ShopNickname  string `json:"shop_nickname"`
	ShopHeadImage string `json:"shop_head_image"`
}

var _ kfMessageDetail = (*kfChannelsShopProductMessage)(nil)

type kfChannelsShopOrder struct {
	OrderId       string `json:"order_id"`
	ProductTitles string `json:"product_titles"`
	PriceWording  string `json:"price_wording"`
	State         string `json:"state"`
	ImageUrl      string `json:"image_url"`
	ShopNickname  string `json:"shop_nickname"`
}

var _ kfMessageDetail = (*kfChannelsShopOrder)(nil)

type kfMergedMessage struct {
	Title string `json:"title"`
	Item  []struct {
		SendTime   uint32 `json:"send_time"`
		MsgType    uint32 `json:"msgtype"`
		SenderName string `json:"sender_name"`
		MsgContent string `json:"msg_content"`
	} `json:"item"`
}

var _ kfMessageDetail = (*kfMergedMessage)(nil)

type kfChannelsMessage struct {
	SubType  uint32 `json:"sub_type"`
	Nickname string `json:"nickname"`
	Title    string `json:"title"`
}

var _ kfMessageDetail = (*kfChannelsMessage)(nil)

type kfEventEnterSession struct {
	EventType      KfEventType `json:"event_type"`
	OpenKfId       string      `json:"open_kfid"`
	ExternalUserId string      `json:"external_userid"`
	Scene          string      `json:"scene"`
	SceneParam     string      `json:"scene_param"`
	WelcomeCode    string      `json:"welcome_code"`
	WechatChannels struct {
		Nickname     string `json:"nickname"`
		ShopNickname string `json:"shop_nickname"`
		Scene        string `json:"scene"`
	} `json:"wechat_channels"`
}

type kfEventMsgSendFail struct {
	EventType      KfEventType `json:"event_type"`
	OpenKfId       string      `json:"open_kfid"`
	ExternalUserId string      `json:"external_userid"`
	FailMsgId      string      `json:"fail_msgid"`
	FailType       uint32      `json:"fail_type"`
}

type kfEventServicerStatusChange struct {
	EventType      KfEventType `json:"event_type"`
	OpenKfId       string      `json:"open_kfid"`
	ServicerUserId string      `json:"servicer_userid"`
	Status         uint32      `json:"status"`
	StopType       uint32      `json:"stop_type"`
}

type kfEventSessionStatusChange struct {
	EventType         KfEventType `json:"event_type"`
	OpenKfId          string      `json:"open_kfid"`
	ExternalUserId    string      `json:"external_userid"`
	ChangeType        uint32      `json:"change_type"`
	OldServicerUserId string      `json:"old_servicer_userid"`
	NewServicerUserId string      `json:"new_servicer_userid"`
	MsgCode           string      `json:"msg_code"`
}

type kfEventUserRecallMsg struct {
	EventType      KfEventType `json:"event_type"`
	OpenKfId       string      `json:"open_kfid"`
	ExternalUserId string      `json:"external_userid"`
	RecallMsgId    string      `json:"recall_msgid"`
}

type kfEventServicerRecallMsg struct {
	EventType      KfEventType `json:"event_type"`
	OpenKfId       string      `json:"open_kfid"`
	ExternalUserId string      `json:"external_userid"`
	RecallMsgId    string      `json:"recall_msgid"`
	ServicerUserId string      `json:"servicer_userid"`
}

type kfEventRejectCustomerMsgSwitchChange struct {
	EventType      KfEventType `json:"event_type"`
	OpenKfId       string      `json:"open_kfid"`
	ServicerUserId string      `json:"servicer_userid"`
	ExternalUserId string      `json:"external_userid"`
	RejectSwitch   uint32      `json:"reject_switch"`
}

func extractKfMessageExtras(common kfMsgCommon, msg KfMsgRaw) (kfMessageDetail, error) {
	switch common.MsgType {
	case KfMessageTypeText:
		return msg.Text, nil
	case KfMessageTypeImage:
		return msg.Image, nil
	case KfMessageTypeVoice:
		return msg.Voice, nil
	case KfMessageTypeVideo:
		return msg.Video, nil
	case KfMessageTypeFile:
		return msg.File, nil
	case KfMessageTypeLocation:
		return msg.Location, nil
	case KfMessageTypeLink:
		return msg.Link, nil
	case KfMessageTypeBusinessCard:
		return msg.BusinessCard, nil
	case KfMessageTypeMiniProgram:
		return msg.MiniProgram, nil
	case KfMessageTypeMsgMenu:
		return msg.MsgMenu, nil
	case KfMessageTypeChannelsShopProduct:
		return msg.ChannelsShopProduct, nil
	case KfMessageTypeChannelsShopOrder:
		return msg.ChannelsShopOrder, nil
	case KfMessageTypeMergedMsg:
		return msg.MergedMsg, nil
	case KfMessageTypeChannels:
		return msg.Channels, nil
	case KfMessageMeeting:
		return nil, nil
	case KfMessageSchedule:
		return nil, nil
	case KfMessageTypeEvent:
		switch common.EventType {
		case KfEventTypeEnterSession:
			return kfEventEnterSession{
				EventType:      common.EventType,
				OpenKfId:       common.OpenKfId,
				ExternalUserId: common.ExternalUserID,
				Scene:          msg.Event.Scene,
				SceneParam:     msg.Event.SceneParam,
				WelcomeCode:    msg.Event.WelcomeCode,
				WechatChannels: msg.Event.WechatChannels,
			}, nil
		case KfEventTypeMsgSendFail:
			return kfEventMsgSendFail{
				EventType:      common.EventType,
				OpenKfId:       common.OpenKfId,
				ExternalUserId: common.ExternalUserID,
				FailMsgId:      msg.Event.FailMsgId,
				FailType:       msg.Event.FailType,
			}, nil
		case KfEventTypeServicerStatusChange:
			return kfEventServicerStatusChange{
				EventType:      common.EventType,
				OpenKfId:       common.OpenKfId,
				ServicerUserId: common.ServicerUserId,
				Status:         msg.Event.Status,
				StopType:       msg.Event.StopType,
			}, nil
		case KfEventTypeSessionStatusChange:
			return kfEventSessionStatusChange{
				EventType:         common.EventType,
				OpenKfId:          common.OpenKfId,
				ExternalUserId:    common.ExternalUserID,
				ChangeType:        msg.Event.ChangeType,
				OldServicerUserId: msg.Event.OldServicerUserId,
				NewServicerUserId: msg.Event.NewServicerUserId,
				MsgCode:           msg.Event.MsgCode,
			}, nil
		case KfEventTypeUserRecallMsg:
			return kfEventUserRecallMsg{
				EventType:      common.EventType,
				OpenKfId:       common.OpenKfId,
				ExternalUserId: common.ExternalUserID,
				RecallMsgId:    msg.Event.RecallMsgId,
			}, nil
		case KfEventTypeServicerRecallMsg:
			return kfEventServicerRecallMsg{
				EventType:      common.EventType,
				OpenKfId:       common.OpenKfId,
				ExternalUserId: common.ExternalUserID,
				RecallMsgId:    msg.Event.RecallMsgId,
				ServicerUserId: msg.Event.ServicerUserId,
			}, nil
		case KfEventTypeRejectCustomerMsgSwitchChange:
			return kfEventRejectCustomerMsgSwitchChange{
				EventType:      common.EventType,
				OpenKfId:       common.OpenKfId,
				ServicerUserId: common.ServicerUserId,
				ExternalUserId: common.ExternalUserID,
				RejectSwitch:   msg.Event.RejectSwitch,
			}, nil
		default:
			return nil, fmt.Errorf("unknown event type: %s", common.EventType)
		}
	default:
		return nil, fmt.Errorf("unknown message type: %s", common.MsgType)
	}
}
