package feature

import (
    "github.com/name5566/leaf/gate"
    "server/msg"
)

type FuncParams struct {
    Agent gate.Agent
    Request *msg.Request
}