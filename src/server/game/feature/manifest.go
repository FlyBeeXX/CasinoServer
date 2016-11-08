package feature

import (
    "github.com/name5566/leaf/util"
    "github.com/name5566/leaf/log"
    "server/msg"
    "server/conf"
)

func Manifest(p []interface{}) {
    funcParams := p[0].(FuncParams)

    a := funcParams.Agent
    m := funcParams.Request.Params.(map[string]interface{})
    log.Debug("FuncParams Params os: %v", m["os"])
    log.Debug("FuncParams Params client_version: %v", m["client_version"])

    payload := struct {
        ConfigVersion string `json:"config_version"`
        BundleVersion string `json:"bundle_version"`
    }{conf.ConfigVersion, conf.BundleVersion}

    // response to client
    r1 := msg.Response{Payload: payload, Token: funcParams.Request.Token, Time: util.NowUTCTime(), Ok: 1}

    a.WriteMsg(&r1)
}
