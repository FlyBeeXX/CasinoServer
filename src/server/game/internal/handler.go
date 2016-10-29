package internal

import (
    "github.com/name5566/leaf/log"
    "github.com/name5566/leaf/gate"
    "reflect"
    "server/msg"
    "server/funcmap"
    "server/game/feature"
)

func init() {
    handler(&msg.Request{}, handleRequest)
}

func handler(m interface{}, h interface{}) {
    skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func dispatcher(m string, p interface{}) {
    funcmap.Call(m, p)
}

func handleRequest(args []interface{}) {
    m := args[0].(*msg.Request)
    a := args[1].(gate.Agent)

    log.Debug("request method: %v", m.Method)
    log.Debug("request params: %v", m.Params)

    //a.WriteMsg(&msg.Request{
    //    Method: "Server Response",
    //})
    funcParams := feature.FuncParams{
        Agent: a,
        Request: m,
    }

    dispatcher(m.Method, funcParams)
}
