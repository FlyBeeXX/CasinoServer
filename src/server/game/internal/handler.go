package internal

import (
    "github.com/name5566/leaf/log"
    "github.com/name5566/leaf/gate"
    "reflect"
    "server/msg"
)

func init() {
    handler(&msg.ReqParams{}, handlerRequest)
}

func handler(m interface{}, h interface{}) {
    skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlerRequest(args []interface{}) {
    m := args[0].(*msg.ReqParams)
    a := args[1].(gate.Agent)

    log.Debug("request method: %v", m.Method)

    a.WriteMsg(&msg.ReqParams{
        Method: "Server Response",
    })
}