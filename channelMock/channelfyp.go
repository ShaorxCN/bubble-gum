package channelMock

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/CardInfoLink/bubble-gum/channelMock/model"
	"github.com/CardInfoLink/log"
)

func fypServive(req *model.FypPayReq) []byte {
	log.Debugf("[rcv req]%+v", req)
	return fypPayService(req)
}

func fypPayService(req *model.FypPayReq) []byte {
	fypResp := &model.FypPayResp{
		CommonBody: model.CommonBody{
			ResultCode: "000000",
			ResultMsg:  "SUCCESS",
			InsCD:      req.CommonParams.InsCD,
			MchntCD:    req.CommonParams.MchntCD,
			TermID:     req.CommonParams.TermID,
			RandomStr:  req.CommonParams.RandomStr,
		},

		OrderType:            req.OrderType,
		TotalAmount:          req.OrderAmt,
		AddnInf:              req.AddnInf,
		ReservedCouponFee:    "12",
		ReservedMchntOrderNo: req.MchntOrderNo,
	}
	bytes, _ := xml.Marshal(fypResp)
	return bytes
}

func fypQueryServive(req *model.FyQueryReq) []byte {
	log.Debugf("[rcv req]%+v", req)
	return fypQueryService(req)
}

func fypQueryService(req *model.FyQueryReq) []byte {
	fyQueryResp := &model.FyQueryResp{
		CommonBody: model.CommonBody{
			ResultCode: "000000",
			ResultMsg:  "SUCCESS",
			InsCD:      req.CommonParams.InsCD,
			MchntCD:    req.CommonParams.MchntCD,
			TermID:     req.CommonParams.TermID,
			RandomStr:  req.CommonParams.RandomStr,
		},

		OrderType:         req.OrderType,
		OrderAmt:          "100",
		TransactionID:     fmt.Sprintf("%d", time.Now().Unix()),
		MchntOrderNo:      req.MchntOrderNo,
		ReservedCouponFee: "1",
		TransStat:         "SUCCESS",
	}
	bytes, _ := xml.Marshal(fyQueryResp)
	return bytes
}

func fypPrePayServive(req *model.FyPreCreateReq) []byte {
	log.Debugf("[rcv req]%+v", req)
	return fypPrePayService(req)
}

func fypPrePayService(req *model.FyPreCreateReq) []byte {
	fypResp := &model.FyPreCreateResp{
		CommonBody: model.CommonBody{
			ResultCode: "000000",
			ResultMsg:  "SUCCESS",
			InsCD:      req.CommonParams.InsCD,
			MchntCD:    req.CommonParams.MchntCD,
			TermID:     req.CommonParams.TermID,
			RandomStr:  req.CommonParams.RandomStr,
		},

		OrderType:              req.OrderType,
		QrCode:                 "for test",
		ReservedChannelOrderID: fmt.Sprintf("%d", time.Now().Unix()),
	}
	bytes, _ := xml.Marshal(fypResp)
	return bytes
}
