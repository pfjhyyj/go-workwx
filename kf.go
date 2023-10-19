package workwx

// KfWechatChannel 微信客服进入会话的视频号信息
type KfWechatChannel struct {
	// Nickname 视频号名称，视频号场景值为1、2、3时返回此项
	Nickname string `xml:"Nickname"`
	// ShopNickname 视频号小店名称，视频号场景值为4、5时返回此项
	ShopNickname string `xml:"ShopNickname"`
	// Scene 视频号场景值。1：视频号主页，2：视频号直播间商品列表页，3：视频号商品橱窗页，4：视频号小店商品详情页，5：视频号小店订单页
	Scene string `xml:"Scene"`
}
