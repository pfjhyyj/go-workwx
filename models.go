package workwx

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func marshalIntoJSONBody(x interface{}) ([]byte, error) {
	y, err := json.Marshal(x)
	if err != nil {
		// should never happen unless OOM or similar bad things
		return nil, makeReqMarshalErr(err)
	}

	return y, nil
}

type reqAccessToken struct {
	CorpID     string
	CorpSecret string
}

var _ urlValuer = reqAccessToken{}

func (x reqAccessToken) intoURLValues() url.Values {
	return url.Values{
		"corpid":     {x.CorpID},
		"corpsecret": {x.CorpSecret},
	}
}

type respCommon struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// IsOK 响应体是否为一次成功请求的响应
//
// 实现依据: https://work.weixin.qq.com/api/doc#10013
//
// > 企业微信所有接口，返回包里都有errcode、errmsg。
// > 开发者需根据errcode是否为0判断是否调用成功(errcode意义请见全局错误码)。
// > 而errmsg仅作参考，后续可能会有变动，因此不可作为是否调用成功的判据。
func (x *respCommon) IsOK() bool {
	return x.ErrCode == 0
}

func (x *respCommon) TryIntoErr() error {
	if x.IsOK() {
		return nil
	}

	return &WorkwxClientError{
		Code: x.ErrCode,
		Msg:  x.ErrMsg,
	}
}

type respAccessToken struct {
	respCommon

	AccessToken   string `json:"access_token"`
	ExpiresInSecs int64  `json:"expires_in"`
}

type reqJSAPITicketAgentConfig struct{}

var _ urlValuer = reqJSAPITicketAgentConfig{}

func (x reqJSAPITicketAgentConfig) intoURLValues() url.Values {
	return url.Values{
		"type": {"agent_config"},
	}
}

type reqJSAPITicket struct{}

var _ urlValuer = reqJSAPITicket{}

func (x reqJSAPITicket) intoURLValues() url.Values {
	return url.Values{}
}

type respJSAPITicket struct {
	respCommon

	Ticket        string `json:"ticket"`
	ExpiresInSecs int64  `json:"expires_in"`
}

// reqMessage 消息发送请求
type reqMessage struct {
	ToUser  []string
	ToParty []string
	ToTag   []string
	ChatID  string
	AgentID int64
	MsgType string
	Content map[string]interface{}
	IsSafe  bool
}

var _ bodyer = reqMessage{}

func (x reqMessage) intoBody() ([]byte, error) {
	// fuck
	safeInt := 0
	if x.IsSafe {
		safeInt = 1
	}

	obj := map[string]interface{}{
		"msgtype": x.MsgType,
		"agentid": x.AgentID,
		"safe":    safeInt,
	}

	// msgtype polymorphism
	obj[x.MsgType] = x.Content

	// 复用这个结构体，因为是 package-private 的所以这么做没风险
	if x.ChatID != "" {
		obj["chatid"] = x.ChatID
	} else {
		obj["touser"] = strings.Join(x.ToUser, "|")
		obj["toparty"] = strings.Join(x.ToParty, "|")
		obj["totag"] = strings.Join(x.ToTag, "|")
	}

	return marshalIntoJSONBody(obj)
}

// respMessageSend 消息发送响应
type respMessageSend struct {
	respCommon

	InvalidUsers   string `json:"invaliduser"`
	InvalidParties string `json:"invalidparty"`
	InvalidTags    string `json:"invalidtag"`
}

type reqUserGet struct {
	UserID string
}

var _ urlValuer = reqUserGet{}

func (x reqUserGet) intoURLValues() url.Values {
	return url.Values{
		"userid": {x.UserID},
	}
}

// respUserGet 读取成员响应
type respUserGet struct {
	respCommon

	UserDetail
}

// reqUserUpdate 更新成员请求
type reqUserUpdate struct {
	UserDetail *UserDetail
}

var _ bodyer = reqUserUpdate{}

func (x reqUserUpdate) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x.UserDetail)
}

// respUserUpdate 更新成员响应
type respUserUpdate struct {
	respCommon
}

// reqUserList 部门成员请求
type reqUserList struct {
	DeptID     int64
	FetchChild bool
}

var _ urlValuer = reqUserList{}

func (x reqUserList) intoURLValues() url.Values {
	var fetchChild int64
	if x.FetchChild {
		fetchChild = 1
	}

	return url.Values{
		"department_id": {strconv.FormatInt(x.DeptID, 10)},
		"fetch_child":   {strconv.FormatInt(fetchChild, 10)},
	}
}

// respUsersByDeptID 部门成员详情响应
type respUserList struct {
	respCommon

	Users []*UserDetail `json:"userlist"`
}

// reqConvertUserIDToOpenID userid转openid 请求
type reqConvertUserIDToOpenID struct {
	UserID string `json:"userid"`
}

var _ bodyer = reqConvertUserIDToOpenID{}

// respConvertUserIDToOpenID userid转openid 响应
type respConvertUserIDToOpenID struct {
	respCommon

	OpenID string `json:"openid"`
}

func (x reqConvertUserIDToOpenID) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

// reqConvertOpenIDToUserID openid转userid 请求
type reqConvertOpenIDToUserID struct {
	OpenID string `json:"openid"`
}

var _ bodyer = reqConvertOpenIDToUserID{}

// respConvertUserIDToOpenID openid转userid 响应
type respConvertOpenIDToUserID struct {
	respCommon

	UserID string `json:"userid"`
}

func (x reqConvertOpenIDToUserID) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

// reqUserIDByMobile 手机号获取 userid 请求
type reqUserIDByMobile struct {
	Mobile string `json:"mobile"`
}

var _ bodyer = reqUserIDByMobile{}

func (x reqUserIDByMobile) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

// respUserIDByMobile 手机号获取 userid 响应
type respUserIDByMobile struct {
	respCommon

	UserID string `json:"userid"`
}

// EmailType 用户邮箱的类型
//
// 1表示用户邮箱是企业邮箱（默认）
// 2表示用户邮箱是个人邮箱
type EmailType int

const (
	// EmailTypeCorporate 企业邮箱
	EmailTypeCorporate EmailType = 1
	// EmailTypePersonal 个人邮箱
	EmailTypePersonal EmailType = 2
)

// reqUserIDByEmail 邮箱获取 userid 请求
type reqUserIDByEmail struct {
	Email     string    `json:"email"`
	EmailType EmailType `json:"email_type"`
}

var _ bodyer = reqUserIDByEmail{}

func (x reqUserIDByEmail) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

// respUserIDByEmail 邮箱获取 userid 响应
type respUserIDByEmail struct {
	respCommon

	UserID string `json:"userid"`
}

// reqDeptCreate 创建部门
type reqDeptCreate struct {
	DeptInfo *DeptInfo
}

var _ bodyer = reqDeptCreate{}

func (x reqDeptCreate) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x.DeptInfo)
}

// respDeptCreate 创建部门响应
type respDeptCreate struct {
	respCommon

	ID int64 `json:"id"`
}

// reqDeptList 获取部门列表
// 从2022年8月15日10点开始，“企业管理后台 - 管理工具 - 通讯录同步”的新增IP将不能再调用此接口，企业可通过「获取部门ID列表」接口获取部门ID列表。查看调整详情。
// https://developer.work.weixin.qq.com/document/path/96079
type reqDeptList struct {
	HaveID bool
	ID     int64
}

var _ urlValuer = reqDeptList{}

func (x reqDeptList) intoURLValues() url.Values {
	if !x.HaveID {
		return url.Values{}
	}

	return url.Values{
		"id": {strconv.FormatInt(x.ID, 10)},
	}
}

// respDeptList 部门列表响应
type respDeptList struct {
	respCommon

	// TODO: 不要懒惰，把 API 层的类型写好
	Department []*DeptInfo `json:"department"`
}

// reqDeptSimpleList 获取子部门ID列表
type reqDeptSimpleList struct {
	HaveID bool
	ID     int64
}

var _ urlValuer = reqDeptSimpleList{}

func (x reqDeptSimpleList) intoURLValues() url.Values {
	if !x.HaveID {
		return url.Values{}
	}

	return url.Values{
		"id": {strconv.FormatInt(x.ID, 10)},
	}
}

// respDeptSimpleList 部门列表响应
type respDeptSimpleList struct {
	respCommon

	DepartmentIDs []*DeptInfo `json:"department_id"`
}

// reqAppchatGet 获取群聊会话请求
type reqAppchatGet struct {
	ChatID string
}

var _ urlValuer = reqAppchatGet{}

func (x reqAppchatGet) intoURLValues() url.Values {
	return url.Values{
		"chatid": {x.ChatID},
	}
}

// respAppchatGet 获取群聊会话响应
type respAppchatGet struct {
	respCommon

	ChatInfo *ChatInfo `json:"chat_info"`
}

// reqAppchatCreate 创建群聊会话请求
type reqAppchatCreate struct {
	ChatInfo *ChatInfo
}

var _ bodyer = reqAppchatCreate{}

func (x reqAppchatCreate) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x.ChatInfo)
}

// respAppchatCreate 创建群聊会话响应
type respAppchatCreate struct {
	respCommon

	ChatID string `json:"chatid"`
}

// reqAppchatUpdate 修改群聊会话请求
type reqAppchatUpdate struct {
	ChatInfo
	AddMemberUserIDs []string `json:"add_user_list"`
	DelMemberUserIDs []string `json:"del_user_list"`
}

var _ bodyer = reqAppchatUpdate{}

func (x reqAppchatUpdate) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

// respAppchatUpdate 修改群聊会话响应
type respAppchatUpdate struct {
	respCommon
}

// reqMediaUpload 临时素材上传请求
type reqMediaUpload struct {
	Type  string
	Media *Media
}

var _ urlValuer = reqMediaUpload{}
var _ mediaUploader = reqMediaUpload{}

func (x reqMediaUpload) intoURLValues() url.Values {
	return url.Values{
		"type": {x.Type},
	}
}

func (x reqMediaUpload) getMedia() *Media {
	return x.Media
}

// respMediaUpload 临时素材上传响应
type respMediaUpload struct {
	respCommon

	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

// reqMediaUploadImg 永久图片素材上传请求
type reqMediaUploadImg struct {
	Media *Media
}

var _ urlValuer = reqMediaUploadImg{}
var _ mediaUploader = reqMediaUploadImg{}

func (x reqMediaUploadImg) intoURLValues() url.Values {
	return url.Values{}
}

func (x reqMediaUploadImg) getMedia() *Media {
	return x.Media
}

// respMediaUploadImg 永久图片素材上传响应
type respMediaUploadImg struct {
	respCommon

	URL string `json:"url"`
}

// reqExternalContactList 获取客户列表
type reqExternalContactList struct {
	UserID string `json:"userid"`
}

var _ urlValuer = reqExternalContactList{}

func (x reqExternalContactList) intoURLValues() url.Values {
	return url.Values{
		"userid": {x.UserID},
	}
}

// respExternalContactList 获取客户列表
type respExternalContactList struct {
	respCommon

	ExternalUserID []string `json:"external_userid"`
}

// reqExternalContactGet 获取客户详情
type reqExternalContactGet struct {
	ExternalUserID string `json:"external_userid"`
}

var _ urlValuer = reqExternalContactGet{}

func (x reqExternalContactGet) intoURLValues() url.Values {
	return url.Values{
		"external_userid": {x.ExternalUserID},
	}
}

// respExternalContactGet 获取客户详情
type respExternalContactGet struct {
	respCommon
	ExternalContactInfo
}

// ExternalContactInfo 外部联系人信息
type ExternalContactInfo struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowUser      []FollowUser    `json:"follow_user"`
}

// ExternalContactBatchInfo 外部联系人信息
type ExternalContactBatchInfo struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowInfo      FollowInfo      `json:"follow_info"`
}

// BatchListExternalContactsResp 外部联系人信息
type BatchListExternalContactsResp struct {
	Result     []ExternalContactBatchInfo
	NextCursor string
}

// reqExternalContactBatchList 批量获取客户详情
type reqExternalContactBatchList struct {
	UserID string `json:"userid"`
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

var _ bodyer = reqExternalContactBatchList{}

func (x reqExternalContactBatchList) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

// respExternalContactBatchList 批量获取客户详情
type respExternalContactBatchList struct {
	respCommon
	NextCursor          string                     `json:"next_cursor"`
	ExternalContactList []ExternalContactBatchInfo `json:"external_contact_list"`
}

// reqExternalContactRemark 获取客户详情
type reqExternalContactRemark struct {
	Remark *ExternalContactRemark
}

var _ bodyer = reqExternalContactRemark{}

func (x reqExternalContactRemark) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x.Remark)
}

// respExternalContactRemark 获取客户详情
type respExternalContactRemark struct {
	respCommon
}

// reqUserInfoGet 获取访问用户身份
type reqUserInfoGet struct {
	// 通过成员授权获取到的code，最大为512字节。每次成员授权带上的code将不一样，code只能使用一次，5分钟未被使用自动过期。
	Code string
}

var _ urlValuer = reqUserInfoGet{}

func (x reqUserInfoGet) intoURLValues() url.Values {
	return url.Values{
		"code": {x.Code},
	}
}

// respUserInfoGet 部门列表响应
type respUserInfoGet struct {
	respCommon
	UserIdentityInfo
}

// reqExternalContactListCorpTags 获取企业标签库
type reqExternalContactListCorpTags struct {
	// 要查询的标签id，如果不填则获取该企业的所有客户标签，目前暂不支持标签组id
	TagIDs []string `json:"tag_id"`
}

var _ bodyer = reqExternalContactListCorpTags{}

func (x reqExternalContactListCorpTags) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

// respExternalContactListCorpTags 获取企业标签库
type respExternalContactListCorpTags struct {
	respCommon
	// 标签组列表
	TagGroup []ExternalContactCorpTagGroup `json:"tag_group"`
}

// reqExternalContactAddCorpTag 添加企业客户标签
type reqExternalContactAddCorpTag struct {
	ExternalContactCorpTagGroup
}

var _ bodyer = reqExternalContactAddCorpTag{}

func (x reqExternalContactAddCorpTag) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x.ExternalContactCorpTagGroup)
}

// respExternalContactAddCorpTag 添加企业客户标签
type respExternalContactAddCorpTag struct {
	respCommon
	// 标签组列表
	TagGroup ExternalContactCorpTagGroup `json:"tag_group"`
}

// reqExternalContactEditCorpTag 编辑企业客户标签
type reqExternalContactEditCorpTag struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Order uint32 `json:"order"`
}

var _ bodyer = reqExternalContactEditCorpTag{}

func (x reqExternalContactEditCorpTag) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

// respExternalContactEditCorpTag 编辑企业客户标签
type respExternalContactEditCorpTag struct {
	respCommon
}

// reqExternalContactDelCorpTag 删除企业客户标签
type reqExternalContactDelCorpTag struct {
	TagID   []string `json:"tag_id"`
	GroupID []string `json:"group_id"`
}

var _ bodyer = reqExternalContactDelCorpTag{}

func (x reqExternalContactDelCorpTag) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

// respExternalContactDelCorpTag 删除企业客户标签
type respExternalContactDelCorpTag struct {
	respCommon
}

// reqExternalContactMarkTag 编辑企业客户标签
type reqExternalContactMarkTag struct {
	UserID         string   `json:"userid"`
	ExternalUserID string   `json:"external_userid"`
	AddTag         []string `json:"add_tag"`
	RemoveTag      []string `json:"remove_tag"`
}

var _ bodyer = reqExternalContactMarkTag{}

func (x reqExternalContactMarkTag) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

// respExternalContactMarkTag 编辑企业客户标签
type respExternalContactMarkTag struct {
	respCommon
}

// reqJSCode2Session 临时登录凭证校验
type reqJSCode2Session struct {
	JSCode string
}

var _ urlValuer = reqJSCode2Session{}

func (x reqJSCode2Session) intoURLValues() url.Values {
	return url.Values{
		"js_code":    {x.JSCode},
		"grant_type": {"authorization_code"},
	}
}

// respJSCode2Session 临时登录凭证校验
type respJSCode2Session struct {
	respCommon
	JSCodeSession
}

// JSCodeSession 临时登录凭证
type JSCodeSession struct {
	CorpID     string `json:"corpid"`
	UserID     string `json:"userid"`
	SessionKey string `json:"session_key"`
}

type reqMsgAuditListPermitUser struct {
	MsgAuditEdition MsgAuditEdition `json:"type"`
}

var _ bodyer = reqMsgAuditListPermitUser{}

func (x reqMsgAuditListPermitUser) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respMsgAuditListPermitUser struct {
	respCommon
	IDs []string `json:"ids"`
}

type reqMsgAuditCheckSingleAgree struct {
	Infos []CheckMsgAuditSingleAgreeUserInfo `json:"info"`
}

var _ bodyer = reqMsgAuditCheckSingleAgree{}

func (x reqMsgAuditCheckSingleAgree) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respMsgAuditCheckSingleAgree struct {
	respCommon
	AgreeInfo []struct {
		UserID           string              `json:"userid"`
		ExternalOpenID   string              `json:"exteranalopenid"`
		AgreeStatus      MsgAuditAgreeStatus `json:"agree_status"`
		StatusChangeTime int                 `json:"status_change_time"`
	} `json:"agreeinfo"`
}

func (x respMsgAuditCheckSingleAgree) intoCheckSingleAgreeInfoList() (resp []CheckMsgAuditSingleAgreeInfo) {
	for _, agreeInfo := range x.AgreeInfo {
		resp = append(resp, CheckMsgAuditSingleAgreeInfo{
			CheckMsgAuditSingleAgreeUserInfo: CheckMsgAuditSingleAgreeUserInfo{
				UserID:         agreeInfo.UserID,
				ExternalOpenID: agreeInfo.ExternalOpenID,
			},
			AgreeStatus:      agreeInfo.AgreeStatus,
			StatusChangeTime: time.Unix(int64(agreeInfo.StatusChangeTime), 0),
		})
	}
	return resp
}

type reqMsgAuditCheckRoomAgree struct {
	RoomID string `json:"roomid"`
}

var _ bodyer = reqMsgAuditCheckRoomAgree{}

func (x reqMsgAuditCheckRoomAgree) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respMsgAuditCheckRoomAgree struct {
	respCommon
	AgreeInfo []struct {
		StatusChangeTime int                 `json:"status_change_time"`
		AgreeStatus      MsgAuditAgreeStatus `json:"agree_status"`
		ExternalOpenID   string              `json:"exteranalopenid"`
	} `json:"agreeinfo"`
}

func (x respMsgAuditCheckRoomAgree) intoCheckRoomAgreeInfoList() (resp []CheckMsgAuditRoomAgreeInfo) {
	for _, agreeInfo := range x.AgreeInfo {
		resp = append(resp, CheckMsgAuditRoomAgreeInfo{
			StatusChangeTime: time.Unix(int64(agreeInfo.StatusChangeTime), 0),
			AgreeStatus:      agreeInfo.AgreeStatus,
			ExternalOpenID:   agreeInfo.ExternalOpenID,
		})
	}
	return resp
}

type reqMsgAuditGetGroupChat struct {
	RoomID string `json:"roomid"`
}

var _ bodyer = reqMsgAuditGetGroupChat{}

func (x reqMsgAuditGetGroupChat) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respMsgAuditGetGroupChat struct {
	respCommon
	Members []struct {
		MemberID string `json:"memberid"`
		JoinTime int    `json:"jointime"`
	} `json:"members"`
	RoomName       string `json:"roomname"`
	Creator        string `json:"creator"`
	RoomCreateTime int    `json:"room_create_time"`
	Notice         string `json:"notice"`
}

func (x respMsgAuditGetGroupChat) intoGroupChat() (resp MsgAuditGroupChat) {
	resp.Creator = x.Creator
	resp.Notice = x.Notice
	resp.RoomName = x.RoomName
	resp.RoomCreateTime = time.Unix(int64(x.RoomCreateTime), 0)
	for _, member := range x.Members {
		resp.Members = append(resp.Members, MsgAuditGroupChatMember{
			MemberID: member.MemberID,
			JoinTime: time.Unix(int64(member.JoinTime), 0),
		})
	}
	return resp
}

type reqListUnassignedExternalContact struct {
	// PageID 分页查询，要查询页号，从0开始
	PageID uint32 `json:"page_id"`
	// PageSize 每次返回的最大记录数，默认为1000，最大值为1000
	PageSize uint32 `json:"page_size"`
	// Cursor 分页查询游标，字符串类型，适用于数据量较大的情况，如果使用该参数则无需填写page_id，该参数由上一次调用返回
	Cursor string `json:"cursor"`
}

var _ bodyer = reqListUnassignedExternalContact{}

func (x reqListUnassignedExternalContact) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respListUnassignedExternalContact struct {
	respCommon
	Info []struct {
		HandoverUserID string `json:"handover_userid"`
		ExternalUserID string `json:"external_userid"`
		DemissionTime  int    `json:"dimission_time"`
	} `json:"info"`
	IsLast     bool   `json:"is_last"`
	NextCursor string `json:"next_cursor"`
}

func (x respListUnassignedExternalContact) intoExternalContactUnassignedList() (resp ExternalContactUnassignedList) {
	list := make([]ExternalContactUnassigned, 0, len(x.Info))
	for _, info := range x.Info {
		list = append(list, ExternalContactUnassigned{
			HandoverUserID: info.HandoverUserID,
			ExternalUserID: info.ExternalUserID,
			DemissionTime:  time.Unix(int64(info.DemissionTime), 0),
		})
	}
	resp.Info = list
	resp.IsLast = x.IsLast
	resp.NextCursor = x.NextCursor
	return resp
}

type reqTransferExternalContact struct {
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `json:"external_userid"`
	// HandoverUserID 原跟进成员的userid
	HandoverUserID string `json:"handover_userid"`
	// TakeoverUserID 接替成员的userid
	TakeoverUserID string `json:"takeover_userid"`
	// TransferSuccessMsg 转移成功后发给客户的消息，最多200个字符，不填则使用默认文案，目前只对在职成员分配客户的情况生效
	TransferSuccessMsg string `json:"transfer_success_msg"`
}

var _ bodyer = reqTransferExternalContact{}

func (x reqTransferExternalContact) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respTransferExternalContact struct {
	respCommon
}

type reqGetTransferExternalContactResult struct {
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `json:"external_userid"`
	// HandoverUserID 原跟进成员的userid
	HandoverUserID string `json:"handover_userid"`
	// TakeoverUserID 接替成员的userid
	TakeoverUserID string `json:"takeover_userid"`
}

var _ bodyer = reqGetTransferExternalContactResult{}

func (x reqGetTransferExternalContactResult) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respGetTransferExternalContactResult struct {
	respCommon
	Status       uint8 `json:"status"`
	TakeoverTime int   `json:"takeover_time"`
}

func (x respGetTransferExternalContactResult) intoExternalContactTransferResult() ExternalContactTransferResult {
	return ExternalContactTransferResult{
		Status:       ExternalContactTransferStatus(x.Status),
		TakeoverTime: time.Unix(int64(x.TakeoverTime), 0),
	}
}

type reqTransferGroupChatExternalContact struct {
	// ChatIDList 需要转群主的客户群ID列表。取值范围： 1 ~ 100
	ChatIDList []string `json:"chat_id_list"`
	// NewOwner 新群主ID
	NewOwner string `json:"new_owner"`
}

var _ bodyer = reqTransferGroupChatExternalContact{}

func (x reqTransferGroupChatExternalContact) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respTransferGroupChatExternalContact struct {
	respCommon
	FailedChatList []ExternalContactGroupChatTransferFailed `json:"failed_chat_list"`
}

type reqOAGetTemplateDetail struct {
	TemplateID string `json:"template_id"`
}

var _ bodyer = reqOAGetTemplateDetail{}

func (x reqOAGetTemplateDetail) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respOAGetTemplateDetail struct {
	respCommon
	OATemplateDetail
}

type reqOAApplyEvent struct {
	OAApplyEvent
}

var _ bodyer = reqOAApplyEvent{}

func (x reqOAApplyEvent) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respOAApplyEvent struct {
	respCommon
	// SpNo 表单提交成功后，返回的表单编号
	SpNo string `json:"sp_no"`
}

type reqOAGetApprovalInfo struct {
	StartTime string                 `json:"starttime"`
	EndTime   string                 `json:"endtime"`
	Cursor    int                    `json:"cursor"`
	Size      uint32                 `json:"size"`
	Filters   []OAApprovalInfoFilter `json:"filters"`
}

var _ bodyer = reqOAGetApprovalInfo{}

func (x reqOAGetApprovalInfo) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respOAGetApprovalInfo struct {
	respCommon
	// SpNoList 审批单号列表，包含满足条件的审批申请
	SpNoList []string `json:"sp_no_list"`
}

type reqOAGetApprovalDetail struct {
	// SpNo 审批单编号。
	SpNo string `json:"sp_no"`
}

var _ bodyer = reqOAGetApprovalDetail{}

func (x reqOAGetApprovalDetail) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respOAGetApprovalDetail struct {
	respCommon
	// Info 审批申请详情
	Info OAApprovalDetail `json:"info"`
}

// TaskCardBtn 任务卡片消息按钮
type TaskCardBtn struct {
	// Key 按钮key值，用户点击后，会产生任务卡片回调事件，回调事件会带上该key值，只能由数字、字母和“_-@”组成，最长支持128字节
	Key string `json:"key"`
	// Name 按钮名称
	Name string `json:"name"`
	// ReplaceName 点击按钮后显示的名称，默认为“已处理”
	ReplaceName string `json:"replace_name"`
	// Color 按钮字体颜色，可选“red”或者“blue”,默认为“blue”
	Color string `json:"color"`
	// IsBold 按钮字体是否加粗，默认false
	IsBold bool `json:"is_bold"`
}

type reqTransferCustomer struct {
	// HandoverUserID 原跟进成员的userid
	HandoverUserID string `json:"handover_userid"`
	// TakeoverUserID 接替成员的userid
	TakeoverUserID string `json:"takeover_userid"`
	// ExternalUserID 客户的external_userid列表，每次最多分配100个客户
	ExternalUserID []string `json:"external_userid"`
	// TransferSuccessMsg 转移成功后发给客户的消息，最多200个字符，不填则使用默认文案
	TransferSuccessMsg string `json:"transfer_success_msg"`
}

var _ bodyer = reqTransferCustomer{}

func (x reqTransferCustomer) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respTransferCustomer struct {
	respCommon
	Customer []struct {
		// ExternalUserID 转接客户的外部联系人userid
		ExternalUserID string `json:"external_userid"`
		// Errcode 对此客户进行分配的结果, 具体可参考全局错误码(https://developer.work.weixin.qq.com/document/path/90475), 0表示成功发起接替,待24小时后自动接替,并不代表最终接替成功
		Errcode int `json:"errcode"`
	} `json:"customer"`
}

func (x respTransferCustomer) intoTransferCustomerResult() TransferCustomerResult {
	return x.Customer
}

type reqGetTransferCustomerResult struct {
	// HandoverUserID 原跟进成员的userid
	HandoverUserID string `json:"handover_userid"`
	// TakeoverUserID 接替成员的userid
	TakeoverUserID string `json:"takeover_userid"`
	// Cursor 分页查询的cursor，每个分页返回的数据不会超过1000条；不填或为空表示获取第一个分页
	Cursor string `json:"cursor"`
}

var _ bodyer = reqGetTransferCustomerResult{}

func (x reqGetTransferCustomerResult) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respGetTransferCustomerResult struct {
	respCommon
	Customer []struct {
		// ExternalUserID 转接客户的外部联系人userid
		ExternalUserID string `json:"external_userid"`
		// Status 接替状态， 1-接替完毕 2-等待接替 3-客户拒绝 4-接替成员客户达到上限 5-无接替记录
		Status int `json:"status"`
		// TakeoverTime 接替客户的时间，如果是等待接替状态，则为未来的自动接替时间
		TakeoverTime int `json:"takeover_time"`
	} `json:"customer"`
	// NextCursor 下个分页的起始cursor
	NextCursor string `json:"next_cursor"`
}

func (x respGetTransferCustomerResult) intoCustomerTransferResult() CustomerTransferResult {
	return CustomerTransferResult{
		Customer:   x.Customer,
		NextCursor: x.NextCursor,
	}
}

type reqListFollowUserExternalContact struct {
}

var _ urlValuer = reqListFollowUserExternalContact{}

func (x reqListFollowUserExternalContact) intoURLValues() url.Values {
	return url.Values{}
}

type respListFollowUserExternalContact struct {
	respCommon
	ExternalContactFollowUserList
}

type reqAddContactExternalContact struct {
	ExternalContactWay
}

var _ bodyer = reqAddContactExternalContact{}

func (x reqAddContactExternalContact) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respAddContactExternalContact struct {
	respCommon
	ExternalContactAddContact
}

type ExternalContactAddContact struct {
	ConfigID string `json:"config_id"`
	QRCode   string `json:"qr_code"`
}

type reqGetContactWayExternalContact struct {
	ConfigID string `json:"config_id"`
}

var _ bodyer = reqGetContactWayExternalContact{}

func (x reqGetContactWayExternalContact) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respGetContactWayExternalContact struct {
	respCommon
	ContactWay ExternalContactContactWay `json:"contact_way"`
}

type ExternalContactContactWay struct {
	ConfigID string `json:"config_id"`
	QRCode   string `json:"qr_code"`
	ExternalContactWay
}

var _ bodyer = reqListContactWayExternalContact{}

func (x reqListContactWayExternalContact) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respListContactWayChatExternalContact struct {
	respCommon
	ExternalContactListContactWayChat
}

type ExternalContactListContactWayChat struct {
	NextCursor string       `json:"next_cursor"`
	ContactWay []contactWay `json:"contact_way"`
}

type contactWay struct {
	ConfigID string `json:"config_id"`
}

var _ bodyer = reqUpdateContactWayExternalContact{}

func (x reqUpdateContactWayExternalContact) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respUpdateContactWayExternalContact struct {
	respCommon
}

type reqDelContactWayExternalContact struct {
	ConfigID string `json:"config_id"`
}

var _ bodyer = reqDelContactWayExternalContact{}

func (x reqDelContactWayExternalContact) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respDelContactWayExternalContact struct {
	respCommon
}

type reqGroupChatList struct {
	ReqChatList ReqChatList
}

func (x reqGroupChatList) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x.ReqChatList)
}

var _ bodyer = reqGroupChatList{}

type respGroupChatList struct {
	respCommon
	*RespGroupChatList
}

type reqGroupChatInfo struct {
	ChatID   string `json:"chat_id"`
	NeedName int64  `json:"need_name"`
}

func (x reqGroupChatInfo) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

var _ bodyer = reqGroupChatInfo{}

type respGroupChatInfo struct {
	respCommon
	GroupChat *RespGroupChatInfo `json:"group_chat"`
}

type reqAddGroupChatJoinWayExternalContact struct {
	ExternalGroupChatJoinWay
}

var _ bodyer = reqAddGroupChatJoinWayExternalContact{}

func (x reqAddGroupChatJoinWayExternalContact) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respAddGroupChatJoinWayExternalContact struct {
	respCommon

	ConfigID string `json:"config_id"`
}

type reqGetGroupChatJoinWayExternalContact struct {
	ConfigID string `json:"config_id"`
}

var _ bodyer = reqGetGroupChatJoinWayExternalContact{}

func (x reqGetGroupChatJoinWayExternalContact) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respGetGroupChatJoinWayExternalContact struct {
	respCommon
	JoinWay ExternalContactGroupChatJoinWay `json:"join_way"`
}

type ExternalContactGroupChatJoinWay struct {
	ConfigID string `json:"config_id"`
	QRCode   string `json:"qr_code"`
	ExternalGroupChatJoinWay
}

type reqUpdateGroupChatJoinWayExternalContact struct {
	ConfigID string `json:"config_id"`
	ExternalGroupChatJoinWay
}

var _ bodyer = reqUpdateGroupChatJoinWayExternalContact{}

func (x reqUpdateGroupChatJoinWayExternalContact) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respUpdateGroupChatJoinWayExternalContact struct {
	respCommon
}

type reqDelGroupChatJoinWayExternalContact struct {
	ConfigID string `json:"config_id"`
}

var _ bodyer = reqDelGroupChatJoinWayExternalContact{}

func (x reqDelGroupChatJoinWayExternalContact) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respDelGroupChatJoinWayExternalContact struct {
	respCommon
}

type reqCloseTempChatExternalContact struct {
	UserID         string `json:"userid"`
	ExternalUserID string `json:"external_userid"`
}

var _ bodyer = reqCloseTempChatExternalContact{}

func (x reqCloseTempChatExternalContact) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respCloseTempChatExternalContact struct {
	respCommon
}

type reqAddMsgTemplateExternalContact struct {
	AddMsgTemplateExternalContact
}

var _ bodyer = reqAddMsgTemplateExternalContact{}

func (x reqAddMsgTemplateExternalContact) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respAddMsgTemplateExternalContact struct {
	respCommon
	AddMsgTemplateDetail
}

type AddMsgTemplateDetail struct {
	FailList []string `json:"fail_list"`
	MsgId    string   `json:"msgid"`
}

// reqSendWelcomeMsgExternalContact 发送新客户欢迎语
type reqSendWelcomeMsgExternalContact struct {
	SendWelcomeMsgExternalContact
}

var _ bodyer = reqSendWelcomeMsgExternalContact{}

func (x reqSendWelcomeMsgExternalContact) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respSendWelcomeMsgExternalContact struct {
	respCommon
}

// reqExternalContactAddCorpTag 添加企业客户标签
type reqExternalContactAddCorpTagGroup struct {
	ExternalContactAddCorpTagGroup
}

var _ bodyer = reqExternalContactAddCorpTagGroup{}

func (x reqExternalContactAddCorpTagGroup) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x.ExternalContactAddCorpTagGroup)
}

// reqAddKfAccount 添加客服账号
type reqAddKfAccount struct {
	Name    string `json:"name"`
	MediaID string `json:"media_id"`
}

var _ bodyer = reqAddKfAccount{}

func (x reqAddKfAccount) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respAddKfAccount struct {
	respCommon
	OpenKfId string `json:"open_kfid"`
}

// reqDeleteKfAccount 删除客服账号
type reqDelKfAccount struct {
	OpenKfId string `json:"open_kfid"`
}

var _ bodyer = reqDelKfAccount{}

func (x reqDelKfAccount) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respDelKfAccount struct {
	respCommon
}

// reqUpdateKfAccount 修改客服账号
type reqUpdateKfAccount struct {
	OpenKfId string `json:"open_kfid"`
	Name     string `json:"name"`
	MediaID  string `json:"media_id"`
}

var _ bodyer = reqUpdateKfAccount{}

func (x reqUpdateKfAccount) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respUpdateKfAccount struct {
	respCommon
}

// reqListKfAccount 获取客服列表
type reqListKfAccount struct {
	Offset uint32 `json:"offset"`
	Limit  uint32 `json:"limit"`
}

var _ bodyer = reqListKfAccount{}

func (x reqListKfAccount) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respListKfAccount struct {
	respCommon
	AccountList []KfAccount `json:"account_list"`
}

type KfAccount struct {
	OpenKfId        string `json:"open_kfid"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	ManagePrivilege bool   `json:"manage_privilege"`
}

// reqAddKfContactWay 获取客服账号链接息
type reqAddKfContactWay struct {
	OpenKfId string `json:"open_kfid"`
	Scene    string `json:"scene"`
}

var _ bodyer = reqAddKfContactWay{}

func (x reqAddKfContactWay) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respAddKfContactWay struct {
	respCommon
	Url string `json:"url"`
}

// reqAddServicer 添加接待人员
type reqAddServicer struct {
	OpenKfId         string   `json:"open_kfid"`
	UserIdList       []string `json:"user_id_list"`
	DepartmentIdList []uint   `json:"department_id_list"`
}

var _ bodyer = reqAddServicer{}

func (x reqAddServicer) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respAddServicer struct {
	respCommon
	ResultList []KfAccountResult `json:"result_list"`
}

type KfAccountResult struct {
	UserId       string `json:"userid"`
	DepartmentId uint   `json:"department_id"`
	respCommon
}

// reqDelServicer 删除接待人员
type reqDelServicer struct {
	OpenKfId         string   `json:"open_kfid"`
	UserIdList       []string `json:"userid_list"`
	DepartmentIdList []uint   `json:"department_id_list"`
}

var _ bodyer = reqDelServicer{}

func (x reqDelServicer) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respDelServicer struct {
	respCommon
	ResultList []KfAccountResult `json:"result_list"`
}

// reqListServicer 获取接待人员列表
type reqListServicer struct {
	OpenKfId string `json:"open_kfid"`
}

var _ urlValuer = reqListServicer{}

func (x reqListServicer) intoURLValues() url.Values {
	return url.Values{
		"open_kfid": {x.OpenKfId},
	}
}

type respListServicer struct {
	respCommon
	ServicerList []KfServicer `json:"servicer_list"`
}

type KfServicer struct {
	UserId       string `json:"userid"`
	DepartmentId uint   `json:"department_id"`
	Status       uint   `json:"status"`
	StopType     uint   `json:"stopType"`
}

// reqGetKfServiceState 获取微信客服会话状态
type reqGetKfServiceState struct {
	OpenKfId       string `json:"open_kfid"`
	ExternalUserID string `json:"external_userid"`
}

var _ bodyer = reqGetKfServiceState{}

func (x reqGetKfServiceState) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respGetKfServiceState struct {
	respCommon
	ServiceState  int    `json:"service_state"`
	ServiceUserID string `json:"servicer_userid"`
}

// reqTransKfServiceState 变更会话状态
type reqTransKfServiceState struct {
	OpenKfId       string `json:"open_kfid"`
	ExternalUserID string `json:"external_userid"`
	ServiceState   int    `json:"service_state"`
	ServiceUserID  string `json:"servicer_userid"`
}

var _ bodyer = reqTransKfServiceState{}

func (x reqTransKfServiceState) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respTransKfServiceState struct {
	respCommon
	MsgCode string `json:"msg_code"`
}

// reqSyncKfMsg 读取消息
type reqSyncKfMsg struct {
	Cursor      string `json:"cursor"`
	Token       string `json:"token"`
	Limit       uint32 `json:"limit"`
	VoiceFormat uint32 `json:"voice_format"`
	OpenKfId    string `json:"open_kfid"`
}

var _ bodyer = reqSyncKfMsg{}

func (x reqSyncKfMsg) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respSyncKfMsg struct {
	respCommon
	NextCursor string      `json:"next_cursor"`
	HasMore    uint32      `json:"has_more"`
	MsgList    []*kfMsgRaw `json:"msg_list"`
}

type kfMsgRaw struct {
	MsgId               string                       `json:"msgid"`
	OpenKfId            string                       `json:"open_kfid"`
	ExternalUserID      string                       `json:"external_userid"`
	SendTime            uint64                       `json:"send_time"`
	Origin              KfOriginType                 `json:"origin"`
	ServicerUserId      string                       `json:"servicer_userid"`
	MsgType             KfMessageType                `json:"msgtype"`
	Text                KfTextMessage                `json:"text"`
	Image               KfImageMessage               `json:"image"`
	Voice               KfVoiceMessage               `json:"voice"`
	Video               KfVideoMessage               `json:"video"`
	File                KfFileMessage                `json:"file"`
	Location            KfLocationMessage            `json:"location"`
	Link                KfLinkMessage                `json:"link"`
	BusinessCard        KfBusinessCardMessage        `json:"business_card"`
	MiniProgram         KfMiniProgramMessage         `json:"miniprogram"`
	MsgMenu             KfMsgMenu                    `json:"msgmenu"`
	ChannelsShopProduct KfChannelsShopProductMessage `json:"channels_shop_product"`
	ChannelsShopOrder   KfChannelsShopOrderMessage   `json:"channels_shop_order"`
	MergedMsg           KfMergedMessage              `json:"merged_msg"`
	Channels            KfChannelsMessage            `json:"channels"`
	Event               struct {
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
		FailMsgId         string `json:"fail_msgid"`
		FailType          uint32 `json:"fail_type"`
		Status            uint32 `json:"status"`
		StopType          uint32 `json:"stop_type"`
		ServicerUserId    string `json:"servicer_userid"`
		ChangeType        uint32 `json:"change_type"`
		OldServicerUserId string `json:"old_servicer_userid"`
		NewServicerUserId string `json:"new_servicer_userid"`
		MsgCode           string `json:"msg_code"`
		RecallMsgId       string `json:"recall_msgid"`
		RejectSwitch      uint32 `json:"reject_switch"`
	} `json:"event"`
}

// reqSendKfMsg 发送客服消息
type reqSendKfMsg struct {
	ToUser   string `json:"touser"`
	OpenKfId string `json:"open_kfid"`
	MsgId    string `json:"msgid"`
	MsgType  string `json:"msgtype"`
	Content  map[string]interface{}
}

type KfMsgMenu struct {
	Type        string               `json:"type"`
	Click       KfMsgMenuClick       `json:"click"`
	View        KfMsgMenuView        `json:"view"`
	MiniProgram KfMsgMenuMiniProgram `json:"miniProgram"`
	Text        KfMsgMenuText        `json:"text"`
}

type KfMsgMenuClick struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

type KfMsgMenuView struct {
	Url     string `json:"url"`
	Content string `json:"content"`
}

type KfMsgMenuMiniProgram struct {
	Appid    string `json:"appid"`
	Pagepath string `json:"pagepath"`
	Content  string `json:"content"`
}

type KfMsgMenuText struct {
	Content   string `json:"content"`
	NoNewLine bool   `json:"no_new_line"`
}

var _ bodyer = reqSendKfMsg{}

func (x reqSendKfMsg) intoBody() ([]byte, error) {
	obj := map[string]interface{}{
		"touser":    x.ToUser,
		"open_kfid": x.OpenKfId,
		"msgid":     x.MsgId,
		"msgtype":   x.MsgType,
	}

	// msgtype polymorphism
	obj[x.MsgType] = x.Content

	return marshalIntoJSONBody(obj)
}

type respSendKfMsg struct {
	respCommon
	MsgId string `json:"msgid"`
}

// reqSendKfMsgOnEvent 发送欢迎语等事件响应消息
type reqSendKfMsgOnEvent struct {
	Code    string `json:"code"`
	MsgId   string `json:"msgid"`
	MsgType string `json:"msgtype"`
	Content map[string]interface{}
}

var _ bodyer = reqSendKfMsgOnEvent{}

func (x reqSendKfMsgOnEvent) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respSendKfMsgOnEvent struct {
	respCommon
	MsgId string `json:"msgid"`
}

// reqGetUpgradeServiceConfig 获取配置的专员与客户群
type reqGetUpgradeServiceConfig struct{}

var _ urlValuer = reqGetUpgradeServiceConfig{}

func (x reqGetUpgradeServiceConfig) intoURLValues() url.Values {
	return url.Values{}
}

type respGetUpgradeServiceConfig struct {
	respCommon
}

// reqUpgradeService 为客户升级为专员或客户群服务
type reqUpgradeService struct {
	OpenKfId       string             `json:"open_kfid"`
	ExternalUserID string             `json:"external_userid"`
	Type           uint32             `json:"type"`
	Member         KfUpgradeMember    `json:"member"`
	GroupChat      KfUpgradeGroupChat `json:"groupchat"`
}

type KfUpgradeMember struct {
	UserId  string `json:"userid"`
	Wording string `json:"wording"`
}

type KfUpgradeGroupChat struct {
	ChatId  string `json:"chat_id"`
	Wording string `json:"wording"`
}

var _ bodyer = reqUpgradeService{}

func (x reqUpgradeService) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respUpgradeService struct {
	respCommon
}

// reqCancelUpgradeService 为客户取消推荐
type reqCancelUpgradeService struct {
	OpenKfId       string `json:"open_kfid"`
	ExternalUserID string `json:"external_userid"`
}

var _ bodyer = reqCancelUpgradeService{}

func (x reqCancelUpgradeService) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respCancelUpgradeService struct {
	respCommon
}

// reqBatchGetCustomer 获取客户基础信息
type reqBatchGetCustomer struct {
	ExternalUserIDList      []string `json:"external_userid_list"`
	NeedEnterSessionContext uint8    `json:"need_enter_session_context"`
}

var _ bodyer = reqBatchGetCustomer{}

func (x reqBatchGetCustomer) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respBatchGetCustomer struct {
	respCommon
	CustomerList          []KfCustomer `json:"customer_list"`
	InvalidExternalUserID []string     `json:"invalid_external_userid"`
}

type KfCustomer struct {
	ExternalUserID      string                `json:"external_userid"`
	Nickname            string                `json:"nickname"`
	Avatar              string                `json:"avatar"`
	Gender              int                   `json:"gender"`
	UnionId             string                `json:"unionid"`
	EnterSessionContext KfEnterSessionContext `json:"enter_session_context"`
}

type KfEnterSessionContext struct {
	Scene          string           `json:"scene"`
	SceneParam     string           `json:"scene_param"`
	WechatChannels KfWechatChannels `json:"wechat_channels"`
}

type KfWechatChannels struct {
	Nickname     string `json:"nickname"`
	ShopNickname string `json:"shop_nickname"`
	Scene        uint32 `json:"scene"`
}

// reqGetCorpStatistic 获取「客户数据统计」企业汇总数据
type reqGetCorpStatistic struct {
	OpenKfId  string `json:"open_kfid"`
	StartTime uint32 `json:"start_time"`
	EndTime   uint32 `json:"end_time"`
}

var _ bodyer = reqGetCorpStatistic{}

func (x reqGetCorpStatistic) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respGetCorpStatistic struct {
	respCommon
	StatisticList []KfCorpStatistic `json:"statistic_list"`
}

type KfCorpStatistic struct {
	StartTime uint32 `json:"start_time"`
}

type KfCorpStatisticDetail struct {
	SessionCnt                uint64  `json:"session_cnt"`
	CustomerCnt               uint64  `json:"customer_cnt"`
	CustomerMsgCnt            uint64  `json:"customer_msg_cnt"`
	UpgradeServiceCustomerCnt uint64  `json:"upgrade_service_customer_cnt"`
	AiSessionReplyCnt         uint64  `json:"ai_session_reply_cnt"`
	AIKnowledgeHitRate        float64 `json:"ai_knowledge_hit_rate"`
	MsgRejectedCustomerCnt    uint64  `json:"msg_rejected_customer_cnt"`
}

// reqGetServicerStatistic 获取「客户数据统计」
type reqGetServicerStatistic struct {
	OpenKfId       string `json:"open_kfid"`
	ServicerUserId string `json:"servicer_userid"`
	StartTime      uint32 `json:"start_time"`
	EndTime        uint32 `json:"end_time"`
}

var _ bodyer = reqGetServicerStatistic{}

func (x reqGetServicerStatistic) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respGetServicerStatistic struct {
	respCommon
}

type KfServicerStatistic struct {
	StartTime uint32                    `json:"start_time"`
	Statistic KfServicerStatisticDetail `json:"statistic"`
}

type KfServicerStatisticDetail struct {
	SessionCnt                         uint64  `json:"session_cnt"`
	CustomerCnt                        uint64  `json:"customer_cnt"`
	CustomerMsgCnt                     uint64  `json:"customer_msg_cnt"`
	ReplyRate                          float64 `json:"reply_rate"`
	FirstReplyAverageSec               float64 `json:"first_reply_avg_sec"`
	SatisfactionInvestgateCnt          uint64  `json:"satisfaction_investgate_cnt"`
	SatisfactionParticipationRate      float64 `json:"satisfaction_participation_rate"`
	SatisfiedRate                      float64 `json:"satisfied_rate"`
	MiddlingRate                       float64 `json:"middling_rate"`
	DissatisfiedRate                   float64 `json:"dissatisfied_rate"`
	UpgradeServiceCustomerCnt          uint64  `json:"upgrade_service_customer_cnt"`
	UpgradeServiceMemberInviteCnt      uint64  `json:"upgrade_service_member_invite_cnt"`
	UpgradeServiceMemberCustomerCnt    uint64  `json:"upgrade_service_member_customer_cnt"`
	UpgradeServiceGroupchatInviteCnt   uint64  `json:"upgrade_service_groupchat_invite_cnt"`
	UpgradeServiceGroupchatCustomerCnt uint64  `json:"upgrade_service_groupchat_customer_cnt"`
	MsgRejectedCustomerCnt             uint64  `json:"msg_rejected_customer_cnt"`
}

// reqAddKnowledgeGroup 添加知识库分组
type reqAddKnowledgeGroup struct {
	Name string `json:"name"`
}

var _ bodyer = reqAddKnowledgeGroup{}

func (x reqAddKnowledgeGroup) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respAddKnowledgeGroup struct {
	respCommon
	GroupId string `json:"group_id"`
}

// reqDelKnowledgeGroup 删除知识库分组
type reqDelKnowledgeGroup struct {
	GroupId string `json:"group_id"`
}

var _ bodyer = reqDelKnowledgeGroup{}

func (x reqDelKnowledgeGroup) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respDelKnowledgeGroup struct {
	respCommon
}

// reqModKnowledgeGroup 修改知识库分组
type reqModKnowledgeGroup struct {
	GroupId string `json:"group_id"`
	Name    string `json:"name"`
}

var _ bodyer = reqModKnowledgeGroup{}

func (x reqModKnowledgeGroup) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respModKnowledgeGroup struct {
	respCommon
}

// reqGetKnowledgeGroupList 获取知识库分组列表
type reqGetKnowledgeGroupList struct {
	Cursor  string `json:"cursor"`
	Limit   uint32 `json:"limit"`
	GroupId string `json:"group_id"`
}

var _ bodyer = reqGetKnowledgeGroupList{}

func (x reqGetKnowledgeGroupList) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respGetKnowledgeGroupList struct {
	respCommon
	NextCursor string             `json:"next_cursor"`
	HasMore    uint32             `json:"has_more"`
	GroupList  []KfKnowledgeGroup `json:"group_list"`
}

type KfKnowledgeGroup struct {
	GroupId   string `json:"group_id"`
	Name      string `json:"name"`
	IsDefault uint32 `json:"is_default"`
}

// reqAddKnowledgeIntent 添加知识库问答
type reqAddKnowledgeIntent struct {
	GroupId          string             `json:"group_id"`
	Question         KfQuestion         `json:"question"`
	SimilarQuestions KfSimilarQuestions `json:"similar_questions"`
	Answers          []KfQuestionAnswer `json:"answers"`
}

type KfQuestion struct {
	Text KfQuestionText `json:"text"`
}

type KfQuestionText struct {
	Content string `json:"content"`
}

type KfSimilarQuestions struct {
	Items []KfQuestionText `json:"items"`
}

type KfQuestionAnswer struct {
	Text KfQuestionAnswerText `json:"text"`
}

type KfQuestionAnswerText struct {
	Content string `json:"content"`
}

type KfQuestionAnswerAttachment struct {
	MsgType string `json:"msgtype"`
	Content map[string]interface{}
}

type KfQuestionAnswerAttachmentImage struct {
	MediaID string `json:"media_id"`
	Name    string `json:"name"`
}

type KfQuestionAnswerAttachmentVideo struct {
	MediaID string `json:"media_id"`
	Name    string `json:"name"`
}

type KfQuestionAnswerAttachmentLink struct {
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Url    string `json:"url"`
	PicUrl string `json:"pic_url"`
}

type KfQuestionAnswerAttachmentMiniprogram struct {
	Title        string `json:"title"`
	Appid        string `json:"appid"`
	ThumbMediaId string `json:"thumb_media_id"`
}

var _ bodyer = reqAddKnowledgeIntent{}

func (x reqAddKnowledgeIntent) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respAddKnowledgeIntent struct {
	respCommon
	IntentId string `json:"intent_id"`
}

// reqDelKnowledgeIntent 删除知识库问答
type reqDelKnowledgeIntent struct {
	IntentId string `json:"intent_id"`
}

var _ bodyer = reqDelKnowledgeIntent{}

func (x reqDelKnowledgeIntent) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respDelKnowledgeIntent struct {
	respCommon
}

// reqModKnowledgeIntent 修改知识库问答
type reqModKnowledgeIntent struct {
	IntentId         string             `json:"intent_id"`
	Question         KfQuestion         `json:"question"`
	SimilarQuestions KfSimilarQuestions `json:"similar_questions"`
	Answers          []KfQuestionAnswer `json:"answers"`
}

var _ bodyer = reqModKnowledgeIntent{}

func (x reqModKnowledgeIntent) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respModKnowledgeIntent struct {
	respCommon
}

// reqGetKnowledgeIntentList 获取知识库问答列表
type reqGetKnowledgeIntentList struct {
	Cursor   string `json:"cursor"`
	Limit    uint32 `json:"limit"`
	GroupId  string `json:"group_id"`
	IntentId string `json:"intent_id"`
}

var _ bodyer = reqGetKnowledgeIntentList{}

func (x reqGetKnowledgeIntentList) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(x)
}

type respGetKnowledgeIntentList struct {
	respCommon
	NextCursor string                    `json:"next_cursor"`
	HasMore    uint32                    `json:"has_more"`
	IntentList []KfKnowledgeIntentDetail `json:"intent_list"`
}

type KfKnowledgeIntentDetail struct {
	IntentId string     `json:"intent_id"`
	GroupId  string     `json:"group_id"`
	Question KfQuestion `json:"question"`
}
