package util

import (
	"errors"
	"reflect"

	"github.com/CardInfoLink/mahonia"
)

// GBKTranscoder GBK 编码转换器
var GBKTranscoder = &Transcoder{Encoding: "gbk"}

// Transcoder 编码转换器
type Transcoder struct {
	Encoding string
	encoder  mahonia.Encoder
	decoder  mahonia.Decoder
}

// Encode 编码，把 UTF8 编码 GBK 等编码的字符串转成
func (t *Transcoder) Encode(utf8 string) (other string, ok bool) {
	if t.encoder == nil {
		t.encoder = mahonia.NewEncoder(t.Encoding)
	}
	return t.encoder.ConvertStringOK(utf8)
}

// Decode 解码，把 GBK 等编码的字符串转成 UTF8 编码
func (t *Transcoder) Decode(other string) (utf8 string, ok bool) {
	if t.decoder == nil {
		t.decoder = mahonia.NewDecoder(t.Encoding)
	}
	return t.decoder.ConvertStringOK(other)
}

func (t *Transcoder) EncodeStructString(v interface{}) error {
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return errors.New("interface not ptr")
	}
	struct_val := reflect.ValueOf(v)
	struct_type := struct_val.Type()
	for i := 0; i < struct_type.Elem().NumField(); i++ {

		tag := struct_type.Elem().Field(i).Tag
		urlTag := tag.Get("url")
		if urlTag == "-" || urlTag == "" {
			continue
		}
		field := struct_val.Elem().Field(i)
		if field.Type().Kind() == reflect.String {
			if str, ok := t.Encode(field.String()); ok {
				field.SetString(str)
			}

		}
	}
	return nil
}

func (t *Transcoder) DecodeStructString(v interface{}) error {
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return errors.New("interface not ptr")
	}
	struct_val := reflect.ValueOf(v)
	struct_type := struct_val.Type()
	for i := 0; i < struct_type.Elem().NumField(); i++ {

		tag := struct_type.Elem().Field(i).Tag
		urlTag := tag.Get("url")
		if urlTag == "-" || urlTag == "" {
			continue
		}
		field := struct_val.Elem().Field(i)
		if field.Type().Kind() == reflect.String {
			if str, ok := t.Decode(field.String()); ok {
				field.SetString(str)
			}
		}
	}
	return nil
}
