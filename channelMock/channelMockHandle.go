package channelMock

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/CardInfoLink/bubble-gum/channelMock/model"
	"github.com/CardInfoLink/bubble-gum/channelMock/util"
	"github.com/CardInfoLink/log"
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
	req := &model.MybankReq{
		Sign: &model.MybankSignature{},
	}
	err = xml.Unmarshal(reqData, req)
	log.Errorf("[mybank] unmarshal error:%v", err)
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
		w.Write([]byte("mybank mock error"))
		return
	}
	defer r.Body.Close()

	v, err := url.ParseQuery(string(reqData))

	reqvalue := v.Get("req")
	log.Debugf("[rcv req]%s", reqvalue)

	reqvalue, err = url.QueryUnescape(reqvalue)
	if err != nil {
		log.Errorf("[fyp] urlqueryUnescape error:%v", err)
		w.Write([]byte("fuiou unescape mock error"))
		return
	}
	req := &model.FypPayReq{}
	ret, ok := util.GBKTranscoder.Decode(reqvalue)
	if !ok {
		log.Errorf("[fyp]error in request gbk2utf8")
		w.Write([]byte("fuiou gbk-utf-8 mock error"))
		return
	}

	reqReader := bytes.NewReader([]byte(ret))

	//原报文指定了gbk所以需要指定decoder
	decoder := xml.NewDecoder(reqReader)
	decoder.CharsetReader = makeCharsetReader
	err = decoder.Decode(req)
	log.Errorf("[fuiou] unmarshal error:%v", err)
	if err != nil {
		w.Write([]byte("fuiou Unmarshal mock error"))
		return
	}
	respBytes := fypServive(req)
	if MbpSleep > 0 {
		time.Sleep(time.Duration(MbpSleep) * time.Millisecond)
	}
	log.Debugf("[send resp]%+s", respBytes)

	//整体gbk
	msg, ok := util.GBKTranscoder.Encode(string(respBytes))
	if !ok {
		w.Write([]byte("fuiou encode mock error"))
		return
	}

	msg = url.QueryEscape(msg)

	w.Write([]byte(msg))
	return
}

func FypQueryHandle(w http.ResponseWriter, r *http.Request) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("mybank mock error"))
		return
	}
	defer r.Body.Close()

	v, err := url.ParseQuery(string(reqData))

	reqvalue := v.Get("req")
	log.Debugf("[rcv req]%s", reqvalue)

	reqvalue, err = url.QueryUnescape(reqvalue)
	if err != nil {
		log.Errorf("[fyp] urlqueryUnescape error:%v", err)
		w.Write([]byte("fuiou unescape mock error"))
		return
	}
	req := &model.FyQueryReq{}

	ret, ok := util.GBKTranscoder.Decode(reqvalue)
	if !ok {
		log.Errorf("[fyp]error in request gbk2utf8")
		w.Write([]byte("fuiou gbk-utf-8 mock error"))
		return
	}

	reqReader := bytes.NewReader([]byte(ret))

	//原报文指定了gbk所以需要指定decoder
	decoder := xml.NewDecoder(reqReader)
	decoder.CharsetReader = makeCharsetReader

	err = decoder.Decode(req)
	log.Errorf("[fuiou] unmarshal error:%v", err)
	if err != nil {
		w.Write([]byte("fuiou Unmarshal mock error"))
		return
	}
	respBytes := fypQueryServive(req)
	if MbpSleep > 0 {
		time.Sleep(time.Duration(MbpSleep) * time.Millisecond)
	}
	log.Debugf("[send resp]%+s", respBytes)
	//整体gbk
	msg, ok := util.GBKTranscoder.Encode(string(respBytes))
	if !ok {
		w.Write([]byte("fuiou encode mock error"))
		return
	}

	msg = url.QueryEscape(msg)

	w.Write([]byte(msg))
	return
}

func FypPrePayHandle(w http.ResponseWriter, r *http.Request) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("mybank mock error"))
		return
	}
	defer r.Body.Close()

	v, err := url.ParseQuery(string(reqData))

	reqvalue := v.Get("req")
	log.Debugf("[rcv req]%s", reqvalue)

	reqvalue, err = url.QueryUnescape(reqvalue)
	if err != nil {
		log.Errorf("[fyp] urlqueryUnescape error:%v", err)
		w.Write([]byte("fuiou unescape mock error"))
		return
	}
	req := &model.FyPreCreateReq{}
	ret, ok := util.GBKTranscoder.Decode(reqvalue)
	if !ok {
		log.Errorf("[fyp]error in request gbk2utf8")
		w.Write([]byte("fuiou gbk-utf-8 mock error"))
		return
	}

	reqReader := bytes.NewReader([]byte(ret))

	//原报文指定了gbk所以需要指定decoder
	decoder := xml.NewDecoder(reqReader)
	decoder.CharsetReader = makeCharsetReader

	err = decoder.Decode(req)
	log.Errorf("[fuiou] unmarshal error:%v", err)
	if err != nil {
		w.Write([]byte("fuiou Unmarshal mock error"))
		return
	}
	respBytes := fypPrePayServive(req)
	if MbpSleep > 0 {
		time.Sleep(time.Duration(MbpSleep) * time.Millisecond)
	}
	log.Debugf("[send resp]%+s", respBytes)
	msg, ok := util.GBKTranscoder.Encode(string(respBytes))
	if !ok {
		w.Write([]byte("fuiou encode mock error"))
		return
	}

	msg = url.QueryEscape(msg)

	w.Write([]byte(msg))

	noty := &model.FuiouNotifyReq{
		ResultCode:     "000000",
		ResultMsg:      "SUCCESS",
		InsCD:          req.InsCD,
		MchntCD:        req.MchntCD,
		OrderAmt:       req.OrderAmt,
		TransactionId:  fmt.Sprintf("%d", time.Now().Unix()),
		RandomStr:      "123456",
		Sign:           "nocheck",
		SettleOrderAmt: req.OrderAmt,
		MchntOrderNo:   req.MchntOrderNo,
		TxnFinTs:       time.Now().Format("20060102150405"),
	}

	bs, err := xml.Marshal(noty)
	if err != nil {
		log.Errorf("send notify error:%v", err)
		return
	}

	nty, ok := util.GBKTranscoder.Encode(string(bs))
	if !ok {
		w.Write([]byte("fuiou encode mock error"))
		return
	}

	nty = url.QueryEscape(nty)

	http.Post(req.NotifyURL, "application/xml", bytes.NewReader([]byte(nty)))
	return
}

func makeCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	if charset == "GBK" {
		return input, nil

	}
	return nil, fmt.Errorf("Unknown charset: %s", charset)
}
