package model

import (
	"encoding/xml"
)

//发送的报文实体
type MybankReq struct {
	XMLName xml.Name            `xml:"document"`
	Req     *MyBankReqInterface `xml:"request"`
	Sign    *MybankSignature
}

// PayReq 条码支付请求
type MyBankReqInterface struct {
	XMLName xml.Name `xml:"request"`
	Head    MybankReqHead
	Body    MybankBody `xml:"body"`
}

//网商签名域.未确定，无明确格式定义。TODO
type MybankSignature struct {
	XMLName        xml.Name `xml:"Signature"`
	Xmlns          string   `xml:"xmlns,attr"`
	SInfo          MybankSignInfo
	SignatureValue string `xml:"SignatureValue"`
}

type MybankSignInfo struct {
	XMLName xml.Name `xml:"SignedInfo"`
}

//网商报文头
type MybankReqHead struct {
	XMLName      xml.Name `xml:"head"`
	Version      string   `xml:"Version"`
	Appid        string   `xml:"Appid,omitempty"`        //网商分配appid
	Function     string   `xml:"Function,omitempty"`     //接口名字
	ReqTime      string   `xml:"ReqTime,omitempty"`      //yyyyMMddHHmmss
	ReqTimeZone  string   `xml:"ReqTimeZone,omitempty"`  //时区
	ReqMsgId     string   `xml:"ReqMsgId,omitempty"`     //全局唯一uuid,报文定位唯一标识
	InputCharset string   `xml:"InputCharset,omitempty"` //暂时只支持UTF-8
	Reserve      string   `xml:"Reserve,omitempty"`      //KV格式的保留字段
	SignType     string   `xml:"SignType,omitempty"`     //RSA
}

// 全部放进去，随意处理
type MybankBody struct {
	XMLName              xml.Name    `xml:"body"`
	AuthCode             string      `xml:"AuthCode" validate:"nonzero" bson:"authCode"`                          // 第三方支付授权码
	OutTradeNo           string      `xml:"OutTradeNo" validate:"nonzero" bson:"outTradeNo"`                      // 外部交易号
	Body                 string      `xml:"Body" validate:"nonzero" bson:"body"`                                  // 商品描述
	GoodsTag             string      `xml:"GoodsTag,omitempty" bson:"goodsTag,omitempty"`                         // 商品标记。微信支付代金券或立减优惠功能的参数。
	GoodsDetail          interface{} `xml:"GoodsDetail,omitempty" bson:"goodsDetail,omitempty"`                   // 商品明细列表信息,json格式
	TotalAmount          int64       `xml:"TotalAmount" validate:"nonzero" bson:"totalAmount"`                    // 交易总额度。货币最小单位，人民币：分
	Currency             string      `xml:"Currency" validate:"nonzero" bson:"currency"`                          // 币种。默认CNY
	MerchantID           string      `xml:"MerchantId" validate:"nonzero" bson:"merchantId"`                      // 商户号。网商为商户分配的商户号，通过商户入驻结果查询接口获取
	IsvOrgID             string      `xml:"IsvOrgId" validate:"nonzero" bson:"isvOrgId"`                          // 合作方机构号（网商银行分配）
	ChannelType          string      `xml:"ChannelType" validate:"nonzero" bson:"channelType"`                    // 支付渠道类型。该笔支付走的第三方支付渠道：ALI/WX/QQ/JD
	OperatorID           string      `xml:"OperatorId,omitempty" bson:"operatorId,omitempty"`                     // 操作员ID。门店操作员ID
	StoreID              string      `xml:"StoreId,omitempty" bson:"storeId,omitempty"`                           // 门店ID
	DeviceID             string      `xml:"DeviceId,omitempty" bson:"deviceId,omitempty"`                         // 终端设备号。门店收银设备ID
	DeviceCreateIP       string      `xml:"DeviceCreateIp" validate:"nonzero" bson:"deviceCreateIp"`              // 终端IP。订单生成的机器IP。
	ExpireExpress        string      `xml:"ExpireExpress,omitempty" bson:"expireExpress,omitempty"`               // 订单有效期。指定订单的支付有效时间（以分钟计算），超过有效时间用户将无法支付。若不指定该参数则系统默认设置1小时支付有效时间。参数允许设置范围：1-1440区间的整数值。
	SettleType           string      `xml:"SettleType" validate:"nonzero" bson:"settleType"`                      // 清算方式。可选值：T0：T+0清算按笔清算，目前仅支持清算到余利宝，不支持清算到银行卡。T1：T+1汇总清算，可支持清算到余利宝及清算到银行卡。
	Attach               string      `xml:"Attach,omitempty" bson:"attach,omitempty"`                             // 附加信息，原样返回
	PayLimit             string      `xml:"PayLimit,omitempty" bson:"payLimit,omitempty"`                         // 禁用支付方式。商户禁受理支付方式列表，多个用逗号隔开。可选值：credit：信用卡 pcredit：花呗（仅支付宝）
	DiscountableAmount   int64       `xml:"DiscountableAmount,omitempty" bson:"discountableAmount,omitempty"`     // 可打折金额。货币最小单位，人民币：分。仅支付宝交易有效。如果同时传入了【可打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【交易总额度】=【可打折金额】+【不可打折金额】
	UndiscountableAmount int64       `xml:"UndiscountableAmount,omitempty" bson:"undiscountableAmount,omitempty"` // 不可打折金额。货币最小单位，人民币：分。仅支付宝交易有效。如果同时传入了【可打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【交易总额度】=【可打折金额】+【不可打折金额】
	AlipayStoreID        string      `xml:"AlipayStoreId,omitempty" bson:"alipayStoreId,omitempty"`               // 支付宝的店铺编号，用于精准店铺营销
	SysServiceProviderID string      `xml:"SysServiceProviderId,omitempty" bson:"sysServiceProviderId,omitempty"` // 支付宝支持系统商返佣，该参数作为系统商返佣数据提取的依据，请填写系统商签约协议的PID
	CheckLaterNm         int         `xml:"CheckLaterNm,omitempty" bson:"checkLaterNm,omitempty"`                 // 花呗交易分期数，可选值：3：3期 6：6期 12：12期 每期间隔为一个月。例如，选择3期，所垫付的资金及利息按3个月等额本息还款，每月还款一笔
	GoodsID              string      `xml:"Goodsid" validate:"nonzero"`                                           // 商品ID
	MerchantId           string      `xml:"MerchantId,omitempty" json:"-" `
	IsvOrgId             string      `xml:"IsvOrgId,omitempty" json:"-" `
	GmtPayment           string      `xml:"GmtPayment,omitempty" json:"-" `
	BankType             string      `xml:"BankType,omitempty" json:"-" `
	IsSubscribe          string      `xml:"IsSubscribe,omitempty" json:"-" `
	PayChannelOrderNo    string      `xml:"PayChannelOrderNo,omitempty" json:"-" `
	MerchantOrderNo      string      `xml:"MerchantOrderNo,omitempty" json:"-" `
	SubAppId             string      `xml:"SubAppId,omitempty" json:"-" `
	CouponFee            string      `xml:"CouponFee,omitempty" json:"-" `
	OpenId               string      `xml:"OpenId,omitempty" json:"-" `
	SubOpenId            string      `xml:"SubOpenId,omitempty" json:"-" `
	BuyerLogonId         string      `xml:"BuyerLogonId,omitempty" json:"-" `
	BuyerUserId          string      `xml:"BuyerUserId,omitempty" json:"-" `
	Credit               string      `xml:"Credit,omitempty" json:"-" `
	ReceiptAmount        string      `xml:"ReceiptAmount,omitempty" json:"-" `
	BuyerPayAmount       string      `xml:"BuyerPayAmount,omitempty" json:"-" `
	InvoiceAmount        string      `xml:"InvoiceAmount,omitempty" json:"-" `
}

//应答的报文实体
type MybankResp struct {
	XMLName xml.Name    `xml:"document"`
	Resp    interface{} `xml:"response"`
	Sign    *MybankSignature
}

// RespInfo 应答码组件
type MybankRespInfo struct {
	ResultStatus string `xml:"ResultStatus,omitempty"` // 本次业务处理的状态，默认以下3个状态：S：成功，F：失败，U：未知
	ResultCode   string `xml:"ResultCode,omitempty"`   // 当resultStatus为S时，该字段必定为0000 当resultStatus为F或U时，该字段可以为全局返回码，也可以为业务返回码。 如果为业务返回码，参见业务接口部分。
	ResultMsg    string `xml:"ResultMsg,omitempty"`    // 当resultStatus为S时，该字段可为空当resultStatus为F或U时，需要描述该错误的原因
}

// PayResp 条码支付请求结果
type MybankPayResp struct {
	XMLName xml.Name `xml:"response"`
	Id      string   `xml:"id,attr"`
	Head    MybankRespHead
	Body    MybankPayRespBody
}
