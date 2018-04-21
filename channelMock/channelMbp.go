package channelMock

import (
	"fmt"
	"github.com/CardInfoLink/bubble-gum/channelMock/model"
	"encoding/xml"
	"github.com/CardInfoLink/log"
	"time"
)

func mbpServive(req *model.MybankReq) []byte {
	log.Debugf("[rcv req]%+v", req)

	switch req.Head.Function {
	case "ant.mybank.bkmerchanttrade.pay":
		return mbpPayService(req)
	case "ant.mybank.bkmerchanttrade.payQuery":
		return mbpPayQueryService(req)

	case "ant.mybank.bkmerchanttrade.prePay":
		return mbpPayService(req)
	}
	return nil
}

func mbpPayService(req *model.MybankReq) []byte {
	mbpResp := &model.MybankResp {
		PayResp: model.MybankPayResp {
			Id: "response",
			Head: model.MybankRespHead {
				Version: req.PayReq.Head.Version,
				Appid: req.PayReq.Head.Appid,
				Function: req.PayReq.Head.Function,
				ReqMsgId: req.PayReq.Head.ReqMsgId,
				InputCharset: req.PayReq.Head.InputCharset,
				Reserve: req.PayReq.Head.Reserve,
				SignType: req.PayReq.Head.SignType,
				RespTime: time.Now().Format("20060102150405"),
				RespTimeZone: "UTC+8",
			},
			Body: model.MybankPayRespBody {
				RespInfo: model.MybankRespInfo {
					ResultStatus: "S",
					ResultCode: "0000",
				},
				OutTradeNo: req.PayReq.Body.OutTradeNo,
				TotalAmount: fmt.Sprintf("%d", req.PayReq.Body.TotalAmount),
				MerchantID: req.PayReq.Body.MerchantID,
				Currency: req.PayReq.Body.Currency,
				OrderNo: req.PayReq.Body.OutTradeNo,
			},
		},
		Sign: &model.MybankSignature{},
	}
	bytes, _ := xml.Marshal(mbpResp)
	return bytes
}


func mbpPayQueryService(req *model.MybankReq) []byte {
	mbpResp := &model.MyBankQueryResp {
		PayResp: model.MyBankQueryRespBody {
			Id: "response",
			Head: model.MybankRespHead {
				Version: req.PayReq.Head.Version,
				Appid: req.PayReq.Head.Appid,
				Function: req.PayReq.Head.Function,
				ReqMsgId: req.PayReq.Head.ReqMsgId,
				InputCharset: req.PayReq.Head.InputCharset,
				Reserve: req.PayReq.Head.Reserve,
				SignType: req.PayReq.Head.SignType,
				RespTime: time.Now().Format("20060102150405"),
				RespTimeZone: "UTC+8",
			},
			Body: model.MybankPayRespBody {
				RespInfo: model.MybankRespInfo {
					ResultStatus: "S",
					ResultCode: "0000",
				},
				OutTradeNo: req.PayReq.Body.OutTradeNo,
				TotalAmount: fmt.Sprintf("%d", req.PayReq.Body.TotalAmount),
				MerchantID: req.PayReq.Body.MerchantID,
				Currency: req.PayReq.Body.Currency,
				OrderNo: req.PayReq.Body.OutTradeNo,
			},
		},
		Sign: &model.MybankSignature{},
	}
	bytes, _ := xml.Marshal(mbpResp)
	return bytes
}

