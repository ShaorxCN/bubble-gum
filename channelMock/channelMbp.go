package channelMock

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"github.com/CardInfoLink/bubble-gum/channelMock/model"
	"github.com/CardInfoLink/log"
)

func mbpServive(req *model.MybankReq) []byte {
	log.Debugf("[rcv req]%+v", req)

	switch req.Req.Head.Function {
	case "ant.mybank.bkmerchanttrade.pay":
		return mbpPayService(req)
	case "ant.mybank.bkmerchanttrade.payQuery":
		return mbpPayQueryService(req)

	case "ant.mybank.bkmerchanttrade.prePay":
		return mbpPayPrePayService(req)
	}
	return nil
}

func mbpPayService(req *model.MybankReq) []byte {
	mbpResp := &model.MybankResp{
		Resp: model.MybankPayResp{
			Id: "response",
			Head: model.MybankRespHead{
				Version:      req.Req.Head.Version,
				Appid:        req.Req.Head.Appid,
				Function:     req.Req.Head.Function,
				ReqMsgId:     req.Req.Head.ReqMsgId,
				InputCharset: req.Req.Head.InputCharset,
				Reserve:      req.Req.Head.Reserve,
				SignType:     req.Req.Head.SignType,
				RespTime:     time.Now().Format("20060102150405"),
				RespTimeZone: "UTC+8",
			},
			Body: model.MybankPayRespBody{
				RespInfo: model.MybankRespInfo{
					ResultStatus: "S",
					ResultCode:   "0000",
				},
				OutTradeNo:     req.Req.Body.OutTradeNo,
				TotalAmount:    fmt.Sprintf("%d", req.Req.Body.TotalAmount),
				MerchantID:     req.Req.Body.MerchantID,
				Currency:       req.Req.Body.Currency,
				OrderNo:        req.Req.Body.OutTradeNo,
				ReceiptAmount:  fmt.Sprintf("%0.f", float64(req.Req.Body.TotalAmount)*float64(0.9)),
				BuyerPayAmount: fmt.Sprintf("%0.f", float64(req.Req.Body.TotalAmount)*float64(0.8)),
				CouponFee:      "100",
			},
		},
		Sign: &model.MybankSignature{},
	}
	bytes, _ := xml.Marshal(mbpResp)
	return bytes
}

func mbpPayQueryService(req *model.MybankReq) []byte {
	mbpResp := &model.MybankResp{
		Resp: model.MyBankQueryResp{
			Id: "response",
			Head: model.MybankRespHead{
				Version:      req.Req.Head.Version,
				Appid:        req.Req.Head.Appid,
				Function:     req.Req.Head.Function,
				ReqMsgId:     req.Req.Head.ReqMsgId,
				InputCharset: req.Req.Head.InputCharset,
				Reserve:      req.Req.Head.Reserve,
				SignType:     req.Req.Head.SignType,
				RespTime:     time.Now().Format("20060102150405"),
				RespTimeZone: "UTC+8",
			},
			Body: model.MybankQueryRespBody{
				RespInfo: model.MybankRespInfo{
					ResultStatus: "S",
					ResultCode:   "0000",
				},
				OutTradeNo:      req.Req.Body.OutTradeNo,
				TotalAmount:     "100",
				MerchantID:      req.Req.Body.MerchantID,
				Currency:        req.Req.Body.Currency,
				OrderNo:         req.Req.Body.OutTradeNo,
				ReceiptAmount:   "90",
				BuyerPayAmount:  "80",
				CouponFee:       "100",
				BuyerLogonID:    "123456",
				BuyerUserID:     "23456",
				MerchantOrderNo: fmt.Sprintf("%d", time.Now().Unix()),
				TradeStatus:     "succ",
				GmtPayment:      time.Now().Format("20060102150405"),
			},
		},
		Sign: &model.MybankSignature{},
	}
	bytes, _ := xml.Marshal(mbpResp)
	return bytes
}

func mbpPayPrePayService(req *model.MybankReq) []byte {
	mbpResp := &model.MybankResp{
		Resp: model.MybankPrePayResp{
			ID: "response",
			Head: model.MybankRespHead{
				Version:      req.Req.Head.Version,
				Appid:        req.Req.Head.Appid,
				Function:     req.Req.Head.Function,
				ReqMsgId:     req.Req.Head.ReqMsgId,
				InputCharset: req.Req.Head.InputCharset,
				Reserve:      req.Req.Head.Reserve,
				SignType:     req.Req.Head.SignType,
				RespTime:     time.Now().Format("20060102150405"),
				RespTimeZone: "UTC+8",
			},
			Body: model.MybankPrePayRespBody{
				RespInfo: model.MybankRespInfo{
					ResultStatus: "S",
					ResultCode:   "0000",
				},
				OutTradeNo: req.Req.Body.OutTradeNo,
				OrderNo:    fmt.Sprintf("%d", time.Now().Unix()),
				QrCodeURL:  "just for test",
			},
		},
		Sign: &model.MybankSignature{},
	}
	bes, _ := xml.Marshal(mbpResp)

	nty := &model.MybankNotifyReqDocument{
		Base: model.MybankNotifyReq{
			ID: "rquest",
			Head: model.MybankReqHead{
				Version:      req.Req.Head.Version,
				Appid:        "2017120400000109",
				Function:     "ant.mybank.bkmerchanttrade.prePayNotice",
				ReqMsgId:     "2018042117475060026861",
				InputCharset: "UTF-8",
				SignType:     "RSA",
				ReqTime:      time.Now().Format("20060102150405"),
				ReqTimeZone:  "UTC+8",
			},
			Body: model.MybankNotifyReqBody{},
		},

		Sign: &model.MybankSignature{},
	}

	bs, err := xml.Marshal(nty)
	if err != nil {
		log.Errorf("[mybank] notify error:%v", err)
	}

	http.Post("http://test.quick.ipay.so/scanpay/upNotify/mybank", "application/xml", bytes.NewReader(bs))
	http.Post("http://sandbox.showmoney.cn/scanpay/upNotify/mybank", "application/xml", bytes.NewReader(bs))
	return bes
}
