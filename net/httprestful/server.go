package httprestful

import (
	. "BOT/common/config"
	"BOT/core/ledger"
	"BOT/events"
	"BOT/net/httprestful/common"
	. "BOT/net/httprestful/restful"
	. "BOT/net/protocol"
	"strconv"
)

func StartServer(n Noder) {
	common.SetNode(n)
	ledger.DefaultLedger.Blockchain.BCEvents.Subscribe(events.EventBlockPersistCompleted, SendBlock2NoticeServer)
	func() {
		rest := InitRestServer(common.CheckAccessToken)
		go rest.Start()
	}()
}

func SendBlock2NoticeServer(v interface{}) {

	if len(Parameters.NoticeServerUrl) == 0 || !common.CheckPushBlock() {
		return
	}
	go func() {
		req := make(map[string]interface{})
		req["Height"] = strconv.FormatInt(int64(ledger.DefaultLedger.Blockchain.BlockHeight), 10)
		req = common.GetBlockByHeight(req)

		repMsg, _ := common.PostRequest(req, Parameters.NoticeServerUrl)
		if repMsg[""] == nil {
			//TODO
		}
	}()
}
