package channelMock

import (
	"encoding/json"

	"github.com/CardInfoLink/bubble-gum/channelMock/model"
	"github.com/CardInfoLink/log"
)

type alpPayResp struct {
	Resp model.AlpCommonResponse `json:"alipay_trade_pay_response"`
}

type alpPreCreateResp struct {
	Resp model.AlpCommonResponse `json:"alipay_trade_precreate_response"`
}

type alpQueryResp struct {
	Resp model.AlpCommonResponse `json:"alipay_trade_query_response"`
}

var alpCommonResp_Success = model.AlpCommonResponse{
	Code:        "10000",
	Msg:         "Success",
	TradeStatus: "TRADE_SUCCESS",
}
var alpPayResp_Success = alpPayResp{
	Resp: alpCommonResp_Success,
}

var alpPreCreateResp_Process = alpPreCreateResp{
	Resp: alpCommonResp_Success,
}

var alpQueryResp_Process = alpQueryResp{
	Resp: alpCommonResp_Success,
}

// json: 	alipay_trade_pay_response
func alpServive(req *model.AlpComonRequest) []byte {
	log.Debugf("[rcv req]%+v", req)

	switch req.Method {
	case "alipay.trade.query":
		return alpQueryService(req)
	case "alipay.trade.precreate":
		return alpPreCreateService(req)
	case "alipay.trade.pay":
		return alpPayService(req)

	}
	return nil
}

func alpPayService(req *model.AlpComonRequest) []byte {
	bytes, _ := json.Marshal(alpPayResp_Success)
	return bytes
}

func alpPreCreateService(req *model.AlpComonRequest) []byte {
	bytes, _ := json.Marshal(alpPreCreateResp_Process)
	return bytes
}

func alpQueryService(req *model.AlpComonRequest) []byte {
	bytes, _ := json.Marshal(alpQueryResp_Process)
	return bytes
}
