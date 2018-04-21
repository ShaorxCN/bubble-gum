package channelMock

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"encoding/xml"
	"github.com/CardInfoLink/bubble-gum/channelMock/model"
	"time"
	"github.com/CardInfoLink/log"
	"fmt"
	"bytes"
)

var MbpSleep = 0

func AlpHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vals := r.Form
	m := make(map[string]string, len(vals))
	for k, v := range vals {
		m[k] = v[0]
	}
	log.Debugf("[rcv req]%+v", m)

	bytes, _ := json.Marshal(m)
	req := model.AlpComonRequest{}
	//log.Debugf("[rcv req]%s", bytes)
	json.Unmarshal(bytes, &req)
	respBytes := alpServive(&req)
	time.Sleep(300 * time.Millisecond)
	log.Debugf("[send resp]%+s", respBytes)
	w.Write(respBytes)
	return
}

func WxpHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("to do"))
	return
}

func MbpHandle(w http.ResponseWriter, r *http.Request) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("mybank mock error"))
		return
	}
	defer r.Body.Close()
	log.Debugf("[rcv req]%s", string(reqData))
	req := &model.MybankReq {
		Sign: &model.MybankSignature{},
	}
	err = xml.Unmarshal(reqData, req)
	if err != nil {
		w.Write([]byte("mybank Unmarshal mock error"))
		return
	}
	respBytes := mbpServive(req)
	if MbpSleep > 0 {
		time.Sleep(time.Duration(MbpSleep) * time.Millisecond)
	}
	log.Debugf("[send resp]%+s", respBytes)
	w.Write(respBytes)
	return
}


func FypHandle(w http.ResponseWriter, r *http.Request) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("fuiou mock error"))
		return
	}
	defer r.Body.Close()
	log.Debugf("[rcv req]%s", string(reqData))
	req := &model.FypPayReq{
	}
	err = xml.Unmarshal(reqData, req)
	if err != nil {
		w.Write([]byte("fuiou Unmarshal mock error"))
		return
	}
	respBytes := fypServive(req)
	if MbpSleep > 0 {
		time.Sleep(time.Duration(MbpSleep) * time.Millisecond)
	}
	log.Debugf("[send resp]%+s", respBytes)
	w.Write(respBytes)
	return
}



func FypQueryHandle(w http.ResponseWriter, r *http.Request) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("fuiou mock error"))
		return
	}
	defer r.Body.Close()
	log.Debugf("[rcv req]%s", string(reqData))
	req := &model.FyQueryReq{
	}
	err = xml.Unmarshal(reqData, req)
	if err != nil {
		w.Write([]byte("fuiou Unmarshal mock error"))
		return
	}
	respBytes := fypQueryServive(req)
	if MbpSleep > 0 {
		time.Sleep(time.Duration(MbpSleep) * time.Millisecond)
	}
	log.Debugf("[send resp]%+s", respBytes)
	w.Write(respBytes)
	return
}


func FypPrePayHandle(w http.ResponseWriter, r *http.Request) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("fuiou mock error"))
		return
	}
	defer r.Body.Close()
	log.Debugf("[rcv req]%s", string(reqData))
	req := &model.FyPreCreateReq{
	}
	err = xml.Unmarshal(reqData, req)
	if err != nil {
		w.Write([]byte("fuiou Unmarshal mock error"))
		return
	}
	respBytes := fypPrePayServive(req)
	if MbpSleep > 0 {
		time.Sleep(time.Duration(MbpSleep) * time.Millisecond)
	}
	log.Debugf("[send resp]%+s", respBytes)
	w.Write(respBytes)

	noty := &model.FuiouNotifyReq{
		ResultCode :"000000",
		ResultMsg:"SUCCESS",
		InsCD:req.InsCD,
		MchntCD:req.MchntCD,
		OrderAmt:req.OrderAmt,
		TransactionId:fmt.Sprintf("%d",time.Now().Unix()),
		RandomStr:"123456",
		Sign:"nocheck",
		SettleOrderAmt:req.OrderAmt,
		MchntOrderNo:req.MchntOrderNo,
		TxnFinTs:time.Now().Format("20060102150405"),
	}


	bs ,err := xml.Marshal(noty)
	if err != nil{
		log.Errorf("send notify error:%v", err)
		return
	}

	http.Post(req.NotifyURL, "application/xml",bytes.NewReader(bs))
	return
}