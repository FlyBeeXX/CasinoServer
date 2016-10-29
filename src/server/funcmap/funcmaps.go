package funcmap

import (
    "github.com/name5566/leaf/log"
    "server/game/feature"
)

var (
    funcMaps = map[string]interface{} {
        "network:Test": feature.NetworkTest,
    }

    funcs = NewFuncs(1024)
)

func init() {
    bindFuncMap()
}

func bindFuncMap() {
    for k, v := range funcMaps {
        err := funcs.Bind(k, v)
        if err != nil {
            log.Error("Bind %s: %s", k, err)
        }
    }
}

func Call(name string, params ... interface{}) {
    funcs.Call(name, params)
}
