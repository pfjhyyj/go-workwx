# 接收消息格式

## Models

### `rxMessageCommon` 接收消息的公共部分

Name|XML|Type|Doc
:---|:--|:---|:--
`ToUserName`|`ToUserName`|`string`|企业微信CorpID
`FromUserName`|`FromUserName`|`string`|成员UserID
`CreateTime`|`CreateTime`|`int64`|消息创建时间（整型）
`MsgType`|`MsgType`|`MessageType`|消息类型
`MsgID`|`MsgId`|`int64`|消息id，64位整型
`AgentID`|`AgentID`|`int64`|企业应用的id，整型。可在应用的设置页面查看
`Event`|`Event`|`EventType`|事件类型 MsgType为event存在
`ChangeType`|`ChangeType`|`ChangeType`|变更类型 Event为change_external_contact存在

```go
// MessageType 消息类型
type MessageType string

// MessageTypeText 文本消息
const MessageTypeText MessageType = "text"

// MessageTypeImage 图片消息
const MessageTypeImage MessageType = "image"

// MessageTypeVoice 语音消息
const MessageTypeVoice MessageType = "voice"

// MessageTypeVideo 视频消息
const MessageTypeVideo MessageType = "video"

// MessageTypeLocation 位置消息
const MessageTypeLocation MessageType = "location"

// MessageTypeLink 链接消息
const MessageTypeLink MessageType = "link"

// MessageTypeEvent 事件消息
const MessageTypeEvent MessageType = "event"

// EventType 事件类型
type EventType string

// EventTypeChangeExternalContact 企业客户事件
const EventTypeChangeExternalContact EventType = "change_external_contact"

// EventTypeChangeExternalChat 客户群变更事件
const EventTypeChangeExternalChat EventType = "change_external_chat"

// EventTypeSysApprovalChange 审批申请状态变化回调通知
const EventTypeSysApprovalChange EventType = "sys_approval_change"

// EventTypeChangeContact 通讯录回调通知
const EventTypeChangeContact EventType = "change_contact"

// ChangeType 变更类型
type ChangeType string

// ChangeTypeAddExternalContact 添加企业客户事件
const ChangeTypeAddExternalContact ChangeType = "add_external_contact"

// ChangeTypeEditExternalContact 编辑企业客户事件
const ChangeTypeEditExternalContact ChangeType = "edit_external_contact"

// ChangeTypeAddHalfExternalContact 外部联系人免验证添加成员事件
const ChangeTypeAddHalfExternalContact ChangeType = "add_half_external_contact"

// ChangeTypeDelExternalContact 删除企业客户事件
const ChangeTypeDelExternalContact ChangeType = "del_external_contact"

// ChangeTypeDelFollowUser 删除跟进成员事件
const ChangeTypeDelFollowUser ChangeType = "del_follow_user"

// ChangeTypeTransferFail 客户接替失败事件
const ChangeTypeTransferFail ChangeType = "transfer_fail"

// ChangeTypeCreateUser 新增成员事件
const ChangeTypeCreateUser ChangeType = "create_user"

// ChangeTypeUpdateUser 更新成员事件
const ChangeTypeUpdateUser ChangeType = "update_user"

// EventTypeAppMenuClick 点击菜单
const EventTypeAppMenuClick = "click"

// EventTypeAppMenuCView 打开菜单链接
const EventTypeAppMenuView = "view"

// EventTypeAppMenuScanCodePush 扫码上传
const EventTypeAppMenuScanCodePush = "scancode_push"

// EventTypeAppMenuScanCodeWaitMsg 扫码等待消息
const EventTypeAppMenuScanCodeWaitMsg = "scancode_waitmsg"

// EventTypeAppMenuPicSysPhoto 弹出系统拍照发图
const EventTypeAppMenuPicSysPhoto = "pic_sysphoto"

// EventTypeAppMenuPicPhotoOrAlbum 弹出系统拍照发图
const EventTypeAppMenuPicPhotoOrAlbum = "pic_photo_or_album"

// EventTypeAppMenuPicWeixin 弹出微信相册发图器
const EventTypeAppMenuPicWeixin = "pic_weixin"

// EventTypeAppMenuLocationSelect 弹出微信位置选择器
const EventTypeAppMenuLocationSelect = "location_select"

// EventTypeAppSubscribe 应用订阅
const EventTypeAppSubscribe = "subscribe"

// EventTypeAppUnsubscribe 应用订阅取消
const EventTypeAppUnsubscribe = "unsubscribe"

// EventTypeKfMsgOrEvent 微信客服消息或事件
const EventTypeKfMsgOrEvent = "kf_msg_or_event"

// KFType 微信客服事件类型
type KFType string

// KFTypeEnterSession 微信客服用户进入会话事件
const KFTypeEnterSession = "enter_session"

// KFTypeMsgSendFail 微信客服消息发送失败事件
const KFTypeMsgSendFail = "msg_send_fail"

// KFTypeServicerStatusChange 微信客服接待人员接待状态变更事件
const KFTypeServicerStatusChange = "servicer_status_change"

// KFTypeSessionStatusChange 微信客服会话状态变更事件
const KFTypeSessionStatusChange = "session_status_change"

// KFTypeUserRecallMsg 微信客服用户撤回消息事件
const KFTypeUserRecallMsg = "user_recall_msg"

// KFTypeServicerRecallMsg 微信客服接待人员撤回消息事件
const KFTypeServicerRecallMsg = "servicer_recall_msg"

// KFTypeRejectCustomerMsgSwitchChange 微信客服拒收客户消息变更事件
const KFTypeRejectCustomerMsgSwitchChange = "reject_customer_msg_switch_change"
```

### `rxTextMessageSpecifics` 接收的文本消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`Content`|`Content`|`string`|文本消息内容

### `rxImageMessageSpecifics` 接收的图片消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`PicURL`|`PicUrl`|`string`|图片链接
`MediaID`|`MediaId`|`string`|图片媒体文件id，可以调用获取媒体文件接口拉取，仅三天内有效

### `rxVoiceMessageSpecifics` 接收的语音消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`MediaID`|`MediaId`|`string`|语音媒体文件id，可以调用获取媒体文件接口拉取数据，仅三天内有效
`Format`|`Format`|`string`|语音格式，如amr，speex等

### `rxVideoMessageSpecifics` 接收的视频消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`MediaID`|`MediaId`|`string`|视频媒体文件id，可以调用获取媒体文件接口拉取数据，仅三天内有效
`ThumbMediaID`|`ThumbMediaId`|`string`|视频消息缩略图的媒体id，可以调用获取媒体文件接口拉取数据，仅三天内有效

### `rxLocationMessageSpecifics` 接收的位置消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`Lat`|`Location_X`|`float64`|地理位置纬度
`Lon`|`Location_Y`|`float64`|地理位置经度
`Scale`|`Scale`|`int`|地图缩放大小
`Label`|`Label`|`string`|地理位置信息
`AppType`|`AppType`|`string`|app类型，在企业微信固定返回wxwork，在微信不返回该字段

### `rxLinkMessageSpecifics` 接收的链接消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`Title`|`Title`|`string`|标题
`Description`|`Description`|`string`|描述
`URL`|`Url`|`string`|链接跳转的url
`PicURL`|`PicUrl`|`string`|封面缩略图的url

### `rxEventAddExternalContact` 接收的事件消息，添加企业客户事件

Name|XML|Type|Doc
:---|:--|:---|:--
`UserID`|`UserID`|`string`|企业服务人员的UserID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号
`State`|`State`|`string`|添加此用户的「联系我」方式配置的state参数，可用于识别添加此用户的渠道
`WelcomeCode`|`WelcomeCode`|`string`|欢迎语code，可用于发送欢迎语

### `rxEventEditExternalContact` 接收的事件消息，编辑企业客户事件

Name|XML|Type|Doc
:---|:--|:---|:--
`UserID`|`UserID`|`string`|企业服务人员的UserID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号
`State`|`State`|`string`|添加此用户的「联系我」方式配置的state参数，可用于识别添加此用户的渠道

### `rxEventAddHalfExternalContact` 接收的事件消息，外部联系人免验证添加成员事件

Name|XML|Type|Doc
:---|:--|:---|:--
`UserID`|`UserID`|`string`|企业服务人员的UserID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号
`State`|`State`|`string`|添加此用户的「联系我」方式配置的state参数，可用于识别添加此用户的渠道
`WelcomeCode`|`WelcomeCode`|`string`|欢迎语code，可用于发送欢迎语

### `rxEventDelExternalContact` 接收的事件消息，删除企业客户事件

Name|XML|Type|Doc
:---|:--|:---|:--
`UserID`|`UserID`|`string`|企业服务人员的UserID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号

### `rxEventDelFollowUser` 接收的事件消息，删除跟进成员事件

Name|XML|Type|Doc
:---|:--|:---|:--
`UserID`|`UserID`|`string`|企业服务人员的UserID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号

### `rxEventTransferFail` 接收的事件消息，客户接替失败事件

Name|XML|Type|Doc
:---|:--|:---|:--
`FailReason`|`FailReason`|`string`|接替失败的原因, customer_refused-客户拒绝， customer_limit_exceed-接替成员的客户数达到上限
`UserID`|`UserID`|`string`|企业服务人员的UserID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号

### `rxEventChangeExternalChat` 接收的事件消息，客户群变更事件

Name|XML|Type|Doc
:---|:--|:---|:--
`ToUserName`|`ToUserName`|`string`|企业微信CorpID
`FromUserName`|`FromUserName`|`string`|此事件该值固定为sys，表示该消息由系统生成
`FailReason`|`FailReason`|`string`|接替失败的原因, customer_refused-客户拒绝， customer_limit_exceed-接替成员的客户数达到上限
`ChatID`|`ChatId`|`string`|群ID

### `rxEventSysApprovalChange` 接收的事件消息，审批申请状态变化回调通知

Name|XML|Type|Doc
:---|:--|:---|:--
`ApprovalInfo`|`ApprovalInfo`|`OAApprovalInfo`|审批信息、

### `rxEventChangeTypeCreateUser` 接受的事件消息，新增成员事件

Name|XML|Type|Doc
:---|:--|:---|:--
`UserID`|`UserID`|`string`|成员UserID
`Name`|`Name`|`string`|成员名称
`Department`|`Department`|`string`|成员部门列表，仅返回该应用有查看权限的部门id
`IsLeaderInDept`|`IsLeaderInDept`|`string`|表示所在部门是否为上级，0-否，1-是，顺序与Department字段的部门逐一对应
`Mobile`|`Mobile`|`string`|手机号
`Position`|`Position`|`string`|职位信息。长度为0~64个字节
`Gender`|`Gender`|`int`|性别，1表示男性，2表示女性
`Email`|`Email`|`string`|邮箱
`Status`|`Status`|`int`|激活状态：1=已激活 2=已禁用 4=未激活 已激活代表已激活企业微信或已关注微工作台（原企业号）5=成员退出
`Avatar`|`Avatar`|`string`|头像url。注：如果要获取小图将url最后的”/0”改成”/100”即可。
`Alias`|`Alias`|`string`|成员别名
`Telephone`|`Telephone`|`string`|座机
`Address`|`Address`|`string`|地址
`ExtAttr`|`ExtAttr`|`string`|扩展属性
`Type`|`Type`|`string`|扩展属性类型: 0-本文 1-网页
`Text`|`Text`|`string`|文本属性类型，扩展属性类型为0时填写
`Value`|`Value`|`string`|文本属性内容
`Web`|`Web`|`string`|网页类型属性，扩展属性类型为1时填写
`Title`|`Title`|`string`|网页的展示标题
`Url`|`Url`|`string`|网页的url

### `rxEventChangeTypeUpdateUser` 接受的事件消息，更新成员事件

Name|XML|Type|Doc
:---|:--|:---|:--
`UserID`|`UserID`|`string`|成员UserID
`NewUserID`|`NewUserID`|`string`|新的UserID，变更时推送（userid由系统生成时可更改一次）
`Name`|`Name`|`string`|成员名称
`Department`|`Department`|`string`|成员部门列表，仅返回该应用有查看权限的部门id
`IsLeaderInDept`|`IsLeaderInDept`|`string`|表示所在部门是否为上级，0-否，1-是，顺序与Department字段的部门逐一对应
`Mobile`|`Mobile`|`string`|手机号
`Position`|`Position`|`string`|职位信息。长度为0~64个字节
`Gender`|`Gender`|`int`|性别，1表示男性，2表示女性
`Email`|`Email`|`string`|邮箱
`Status`|`Status`|`int`|激活状态：1=已激活 2=已禁用 4=未激活 已激活代表已激活企业微信或已关注微工作台（原企业号）5=成员退出
`Avatar`|`Avatar`|`string`|头像url。注：如果要获取小图将url最后的”/0”改成”/100”即可。
`Alias`|`Alias`|`string`|成员别名
`Telephone`|`Telephone`|`string`|座机
`Address`|`Address`|`string`|地址
`ExtAttr`|`ExtAttr`|`string`|扩展属性
`Type`|`Type`|`string`|扩展属性类型: 0-本文 1-网页
`Text`|`Text`|`string`|文本属性类型，扩展属性类型为0时填写
`Value`|`Value`|`string`|文本属性内容
`Web`|`Web`|`string`|网页类型属性，扩展属性类型为1时填写
`Title`|`Title`|`string`|网页的展示标题
`Url`|`Url`|`string`|网页的url

### `rxEventAppMenuClick` 接受的事件消息，应用菜单点击事件

Name|XML|Type|Doc
:---|:--|:---|:--
`EventKey`|`EventKey`|`string`|事件key

### `rxEventAppMenuView ` 接受的事件消息，应用菜单点击链接事件

Name|XML|Type|Doc
:---|:--|:---|:--
`EventKey`|`EventKey`|`string`|事件key

### `rxEventAppSubscribe` 接受的事件消息，用户订阅事件

Name|XML|Type|Doc
:---|:--|:---|:--
`EventKey`|`EventKey`|`string`|事件key

### `rxEventAppUnsubscribe` 接受的事件消息，用户取消订阅事件

Name|XML|Type|Doc
:---|:--|:---|:--
`EventKey`|`EventKey`|`string`|事件key

### `rxKfEventEnterSession` 接受的事件消息，用户进入会话事件

Name|XML|Type|Doc
:---|:--|:---|:--
`OpenKFID`|`OpenKfid`|`string`|客服账号ID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号
`Scene`|`Scene`|`string`|进入会话的场景值
`SceneParam`|`SceneParam`|`string`|进入会话的自定义参数
`WelcomeCode`|`WelcomeCode`|`string`|如果满足发送欢迎语条件（条件为：用户在过去48小时里未收过欢迎语，且未向客服发过消息），会返回该字段
`WechatChannels`|`WechatChannels`|`KfWechatChannel`|进入会话的视频号信息，从视频号进入会话才有值

### `rxKfEventMsgSendFail` 接受的事件消息，消息发送失败事件

Name|XML|Type|Doc
:---|:--|:---|:--
`OpenKFID`|`OpenKfid`|`string`|客服账号ID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号
`FailMsgID`|`FailMsgId`|`string`|发送失败的消息ID
`FailType`|`FailType`|`uint32`|失败类型。0-未知原因 1-客服账号已删除 2-应用已关闭 4-会话已过期，超过48小时 5-会话已关闭 6-超过5条限制 8-主体未验证 10-用户拒收 11-企业未有成员登录企业微信App（排查方法：企业至少一个成员通过手机号验证/微信授权登录企业微信App即可）12-发送的消息为客服组件禁发的消息类型

### `rxKfEventServicerStatusChange` 接受的事件消息，接待人员接待状态变更事件

Name|XML|Type|Doc
:---|:--|:---|:--
`ServicerUserID`|`ServicerUserID`|`string`|接待人员的userid
`Status`|`Status`|`uint32`|状态类型。1-接待中 2-停止接待
`StopType`|`StopType`|`uint32`|接待人员的状态为「停止接待」的子类型。0:停止接待,1:暂时挂起
`OpenKFID`|`OpenKfid`|`string`|客服账号ID

### `rxKfEventSessionStatusChange` 接受的事件消息，会话状态变更事件

Name|XML|Type|Doc
:---|:--|:---|:--
`OpenKFID`|`OpenKfid`|`string`|客服账号ID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号
`ChangeType`|`ChangeType`|`uint32`|变更类型，均为接待人员在企业微信客户端操作触发。1-从接待池接入会话 2-转接会话 3-结束会话 4-重新接入已结束/已转接会话
`OldServicerUserID`|`OldServicerUserID`|`string`|原接待人员的userid，仅change_type为2、3和4有值
`NewServicerUserID`|`NewServicerUserID`|`string`|新接待人员的userid，仅change_type为1、2和4有值
`MsgCode`|`MsgCode`|`string`|用于发送事件响应消息的code，仅change_type为1和3时，会返回该字段。可用该msg_code调用发送事件响应消息接口给客户发送回复语或结束语。

### `rxKfEventUserRecallMsg` 接受的事件消息，用户撤回消息事件

Name|XML|Type|Doc
:---|:--|:---|:--
`OpenKFID`|`OpenKfid`|`string`|客服账号ID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号
`RecallMsgID`|`RecallMsgId`|`string`|撤回的消息ID

### `rxKfEventServicerRecallMsg` 接受的事件消息，接待人员撤回消息事件

Name|XML|Type|Doc
:---|:--|:---|:--
`OpenKFID`|`OpenKfid`|`string`|客服账号ID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号
`RecallMsgID`|`RecallMsgId`|`string`|撤回的消息ID
`ServicerUserID`|`ServicerUserID`|`string`|接待人员的userid

### `rxKfEventRejectCustomerMsgSwitchChange` 接受的事件消息，拒收客户消息变更事件

Name|XML|Type|Doc
:---|:--|:---|:--
`OpenKFID`|`OpenKfid`|`string`|客服账号ID
`ExternalUserID`|`ExternalUserID`|`string`|外部联系人的userid，注意不是企业成员的帐号
`ServicerUserID`|`ServicerUserID`|`string`|接待人员的userid
`RejectSwitch`|`RejectSwitch`|`uint32`|拒收客户消息，1表示接待人员拒收了客户消息，0表示接待人员取消拒收客户消息

### `rxEventUnknown` 接受的事件消息，未定义的事件类型

Name|XML|Type|Doc
:---|:--|:---|:--
`EventType`|`-`|`string`|事件类型
`Raw`|`-`|`string`|原始的消息体
