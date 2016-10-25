package gate

import (
    "server/msg"
    "server/game"
)

func init() {
    msg.Processor.SetRouter(&msg.ReqParams{}, game.ChanRPC)
}
