package feature

import (
    "github.com/name5566/leaf/util"
    "github.com/name5566/leaf/log"
    "server/msg"
    "server/conf"
    "strings"
)

func Manifest(p []interface{}) {
    funcParams := p[0].(FuncParams)

    a := funcParams.Agent
    m := funcParams.Request.Params.(map[string]interface{})
    log.Debug("FuncParams Params os: %v", m["os"])
    log.Debug("FuncParams Params client_version: %v", m["client_version"])

    bundle_version := ""
    if strings.Contains(m["os"].(string), "ios") {
        bundle_version = conf.IOSBundleVersion
    } else {
        bundle_version = conf.AndroidBundleVersion
    }

    app_version := struct {
        ConfigVersion string `json:"config_version"`
        BundleVersion string `json:"bundle_version"`
    }{conf.ConfigVersion, bundle_version}

    payload := struct {
        AppVersion interface{} `json:"app_version"`
    }{app_version}

    // response to client
    r1 := msg.Response{Payload: payload, Token: funcParams.Request.Token, Time: util.NowUTCTime(), Ok: 1}

    a.WriteMsg(&r1)
}
