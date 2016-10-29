package feature

import (
    "github.com/name5566/leaf/log"
    "server/msg"
)

func NetworkTest(p []interface{}) {
    funcParams := p[0].(FuncParams)

    a := funcParams.Agent
    m := funcParams.Request.Params.(map[string]interface{})
    log.Debug("FuncParams Params words: %v", m["words"])

    // response to client
    r1 := msg.Response{Payload: "from server", Data: 1}
    r2 := msg.Response{Payload: "from server", Data: 2}
    r3 := msg.Response{Payload: "from server", Data: 3}

    a.WriteMsg(&r1)
    a.WriteMsg(&r2)
    a.WriteMsg(&r3)
}
