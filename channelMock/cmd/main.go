package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/CardInfoLink/bubble-gum/channelMock"
)

func main() {
	startMock()
}

func startMock() {
	flag.IntVar(&channelMock.MbpSleep, "mbpSleep", 0, "Sleep [mbpSleep] ms before mybank responds")
	flag.Parse()

	http.HandleFunc("/mock/alp", channelMock.AlpHandle)
	http.HandleFunc("/mock/wxp", channelMock.WxpHandle)
	http.HandleFunc("/mock/mbp", channelMock.MbpHandle)
	http.HandleFunc("/mock/fyp/micropay", channelMock.FypHandle)
	http.HandleFunc("/mock/fyp/commonQuery", channelMock.FypQueryHandle)
	http.HandleFunc("/mock/fyp/preCreate", channelMock.FypPrePayHandle)
	if err := http.ListenAndServe(":9900", nil); err != nil {
		fmt.Printf("%s\n", err)
	}
}
