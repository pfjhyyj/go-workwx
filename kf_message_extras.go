package workwx

import (
	"fmt"
)

type kfMessageDetail interface {
}

type KfTextMessage struct {
	Content string `json:"content"`
	MenuId  string `json:"menu_id"`
}

var _ kfMessageDetail = (*KfTextMessage)(nil)

type KfImageMessage struct {
	MediaId string `json:"media_id"`
}

var _ kfMessageDetail = (*KfImageMessage)(nil)

type KfVoiceMessage struct {
	MediaId string `json:"media_id"`
}

var _ kfMessageDetail = (*KfVoiceMessage)(nil)

type KfVideoMessage struct {
	MediaId string `json:"media_id"`
}

var _ kfMessageDetail = (*KfVideoMessage)(nil)

type KfFileMessage struct {
	MediaId string `json:"media_id"`
}

var _ kfMessageDetail = (*KfFileMessage)(nil)

type KfLocationMessage struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Name      string `json:"name"`
	Address   string `json:"address"`
}

var _ kfMessageDetail = (*KfLocationMessage)(nil)

type KfLinkMessage struct {
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Url    string `json:"url"`
	PicUrl string `json:"pic_url"`
}

var _ kfMessageDetail = (*KfLinkMessage)(nil)

type KfBusinessCardMessage struct {
	UserId string `json:"userid"`
}

var _ kfMessageDetail = (*KfBusinessCardMessage)(nil)

type KfMiniProgramMessage struct {
	Title        string `json:"title"`
	Appid        string `json:"appid"`
	PagePath     string `json:"pagepath"`
	ThumbMediaId string `json:"thumb_media_id"`
}

var _ kfMessageDetail = (*KfMiniProgramMessage)(nil)

type KfMsgMenuMessage struct {
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

var _ kfMessageDetail = (*KfMsgMenuMessage)(nil)

type KfChannelsShopProductMessage struct {
	ProductId     string `json:"product_id"`
	HeadImage     string `json:"head_image"`
	Title         string `json:"title"`
	SalesPrice    string `json:"sales_price"`
	ShopNickname  string `json:"shop_nickname"`
	ShopHeadImage string `json:"shop_head_image"`
}

var _ kfMessageDetail = (*KfChannelsShopProductMessage)(nil)

type KfChannelsShopOrderMessage struct {
	OrderId       string `json:"order_id"`
	ProductTitles string `json:"product_titles"`
	PriceWording  string `json:"price_wording"`
	State         string `json:"state"`
	ImageUrl      string `json:"image_url"`
	ShopNickname  string `json:"shop_nickname"`
}

var _ kfMessageDetail = (*KfChannelsShopOrderMessage)(nil)

type KfMergedMessage struct {
	Title string `json:"title"`
	Item  []struct {
		SendTime   uint32 `json:"send_time"`
		MsgType    uint32 `json:"msgtype"`
		SenderName string `json:"sender_name"`
		MsgContent string `json:"msg_content"`
	} `json:"item"`
}

var _ kfMessageDetail = (*KfMergedMessage)(nil)

type KfChannelsMessage struct {
	SubType  uint32 `json:"sub_type"`
	Nickname string `json:"nickname"`
	Title    string `json:"title"`
}

var _ kfMessageDetail = (*KfChannelsMessage)(nil)

type KfEventEnterSession struct {
	EventType      KfEventType `json:"event_type"`
	OpenKfId       string      `json:"open_Kfid"`
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

type KfEventMsgSendFail struct {
	EventType      KfEventType `json:"event_type"`
	OpenKfId       string      `json:"open_Kfid"`
	ExternalUserId string      `json:"external_userid"`
	FailMsgId      string      `json:"fail_msgid"`
	FailType       uint32      `json:"fail_type"`
}

type KfEventServicerStatusChange struct {
	EventType      KfEventType `json:"event_type"`
	OpenKfId       string      `json:"open_Kfid"`
	ServicerUserId string      `json:"servicer_userid"`
	Status         uint32      `json:"status"`
	StopType       uint32      `json:"stop_type"`
}

type KfEventSessionStatusChange struct {
	EventType         KfEventType `json:"event_type"`
	OpenKfId          string      `json:"open_Kfid"`
	ExternalUserId    string      `json:"external_userid"`
	ChangeType        uint32      `json:"change_type"`
	OldServicerUserId string      `json:"old_servicer_userid"`
	NewServicerUserId string      `json:"new_servicer_userid"`
	MsgCode           string      `json:"msg_code"`
}

type KfEventUserRecallMsg struct {
	EventType      KfEventType `json:"event_type"`
	OpenKfId       string      `json:"open_Kfid"`
	ExternalUserId string      `json:"external_userid"`
	RecallMsgId    string      `json:"recall_msgid"`
}

type KfEventServicerRecallMsg struct {
	EventType      KfEventType `json:"event_type"`
	OpenKfId       string      `json:"open_Kfid"`
	ExternalUserId string      `json:"external_userid"`
	RecallMsgId    string      `json:"recall_msgid"`
	ServicerUserId string      `json:"servicer_userid"`
}

type KfEventRejectCustomerMsgSwitchChange struct {
	EventType      KfEventType `json:"event_type"`
	OpenKfId       string      `json:"open_Kfid"`
	ServicerUserId string      `json:"servicer_userid"`
	ExternalUserId string      `json:"external_userid"`
	RejectSwitch   uint32      `json:"reject_switch"`
}

func extractKfMessageExtras(common kfMsgCommon, msg kfMsgRaw) (kfMessageDetail, error) {
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
			return KfEventEnterSession{
				EventType:      common.EventType,
				OpenKfId:       common.OpenKfId,
				ExternalUserId: common.ExternalUserID,
				Scene:          msg.Event.Scene,
				SceneParam:     msg.Event.SceneParam,
				WelcomeCode:    msg.Event.WelcomeCode,
				WechatChannels: msg.Event.WechatChannels,
			}, nil
		case KfEventTypeMsgSendFail:
			return KfEventMsgSendFail{
				EventType:      common.EventType,
				OpenKfId:       common.OpenKfId,
				ExternalUserId: common.ExternalUserID,
				FailMsgId:      msg.Event.FailMsgId,
				FailType:       msg.Event.FailType,
			}, nil
		case KfEventTypeServicerStatusChange:
			return KfEventServicerStatusChange{
				EventType:      common.EventType,
				OpenKfId:       common.OpenKfId,
				ServicerUserId: common.ServicerUserId,
				Status:         msg.Event.Status,
				StopType:       msg.Event.StopType,
			}, nil
		case KfEventTypeSessionStatusChange:
			return KfEventSessionStatusChange{
				EventType:         common.EventType,
				OpenKfId:          common.OpenKfId,
				ExternalUserId:    common.ExternalUserID,
				ChangeType:        msg.Event.ChangeType,
				OldServicerUserId: msg.Event.OldServicerUserId,
				NewServicerUserId: msg.Event.NewServicerUserId,
				MsgCode:           msg.Event.MsgCode,
			}, nil
		case KfEventTypeUserRecallMsg:
			return KfEventUserRecallMsg{
				EventType:      common.EventType,
				OpenKfId:       common.OpenKfId,
				ExternalUserId: common.ExternalUserID,
				RecallMsgId:    msg.Event.RecallMsgId,
			}, nil
		case KfEventTypeServicerRecallMsg:
			return KfEventServicerRecallMsg{
				EventType:      common.EventType,
				OpenKfId:       common.OpenKfId,
				ExternalUserId: common.ExternalUserID,
				RecallMsgId:    msg.Event.RecallMsgId,
				ServicerUserId: msg.Event.ServicerUserId,
			}, nil
		case KfEventTypeRejectCustomerMsgSwitchChange:
			return KfEventRejectCustomerMsgSwitchChange{
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
