package feature

import (
    "github.com/name5566/leaf/log"
    "server/msg"
    "time"
)

func NetworkTest(p []interface{}) {
    funcParams := p[0].(FuncParams)

    a := funcParams.Agent
    m := funcParams.Request.Params.(map[string]interface{})
    log.Debug("FuncParams Params words: %v", m["words"])

    // response to client
    r1 := msg.Response{Payload: "from server", Data: 1, Token: funcParams.Request.Token, Time: time.UTC.String(), Ok: 1}

    a.WriteMsg(&r1)
}
